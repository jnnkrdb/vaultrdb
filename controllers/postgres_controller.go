package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	"github.com/jnnkrdb/vaultrdb/controllers/vaultrequest"
	"github.com/jnnkrdb/vaultrdb/svc/postgres"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type PostgresReconciler struct {
	client.Client
	Postgres *sql.DB
}

func (r *PostgresReconciler) Reconcile(ctx context.Context, log logr.Logger) error {

	rows, err := r.Postgres.Query("SELECT psqlid FROM public.vault;")
	if err != nil {
		log.V(0).Error(err, "error receiving list of psqlids from database")
		return err
	}

	// parsing over all received psqlids
	for rows.Next() {
		var psqlid string
		if err = rows.Scan(&psqlid); err != nil {
			log.V(2).Info("error binding psqlid to string", "err.message", err)
			return err
		}

		var configMapList = &v1.ConfigMapList{}
		if err = r.List(ctx, configMapList, &client.ListOptions{
			LabelSelector: labels.SelectorFromSet(labels.Set{
				fmt.Sprintf("v1.vaultrequest.jnnkrdb.de/%s", psqlid): "validkey",
			}),
		}); err != nil {
			log.V(2).Info("error receiving list of affected configmaps", "err.message", err)
			return err
		}

		var secretList = &v1.SecretList{}
		if err = r.List(ctx, secretList, &client.ListOptions{
			LabelSelector: labels.SelectorFromSet(labels.Set{
				fmt.Sprintf("v1.vaultrequest.jnnkrdb.de/%s", psqlid): "validkey",
			}),
		}); err != nil {
			log.V(2).Info("error receiving list of affected secrets", "err.message", err)
			return err
		}

		recon := func(namespace, name string) error {
			var rl = log.WithValues("namespace", namespace, "name", name)

			rl.V(2).Info("requesting vaultrequest")

			var vr = &jnnkrdbdev1.VaultRequest{}
			if err = r.Get(ctx, types.NamespacedName{
				Namespace: namespace,
				Name:      name,
			}, vr, &client.GetOptions{}); err != nil {
				rl.V(2).Info("error receiving vaultrequest")
				return err
			}

			if _, err = vaultrequest.Reconcile(rl, ctx, r.Client, vr); err != nil {
				rl.V(2).Info("error reconciling vaultrequest", "err.message", err)
				return err
			}

			return nil
		}

		if len(configMapList.Items) > 0 {
			if err = recon(
				configMapList.Items[0].Annotations["v1.vaultrequest.jnnkrdb.de/source.namespace"],
				configMapList.Items[0].Annotations["v1.vaultrequest.jnnkrdb.de/source.name"],
			); err != nil {
				return err
			}
		}

		if len(secretList.Items) > 0 {
			if err = recon(
				secretList.Items[0].Annotations["v1.vaultrequest.jnnkrdb.de/source.namespace"],
				secretList.Items[0].Annotations["v1.vaultrequest.jnnkrdb.de/source.name"],
			); err != nil {
				return err
			}
		}

	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PostgresReconciler) SetupWithManager(mgr ctrl.Manager) error {

	var log = log.Log.WithName("postgres-reconciler")

	if postgres.USEPOSTGRES {
		var err error
		if err = r.Postgres.Ping(); err != nil {
			return err
		}
		// starting pseudo reconcilation loop for postgresdb
		go func() {
			for {
				if err = r.Reconcile(context.Background(), log); err != nil {
					log.V(0).Error(err, "error reconciling postgres datasets")
				}
				time.Sleep(time.Minute * 2)
			}
		}()
	}

	return nil
}

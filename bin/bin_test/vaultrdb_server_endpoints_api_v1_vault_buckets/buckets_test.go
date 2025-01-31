package vaultrdbserverendpointsapiv1vaultbuckets_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	api_v1_buckets "github.com/jnnkrdb/vaultrdb/bin/vaultrdb/server/endpoints/api/v1/vault/buckets"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
)

func prep(t *testing.T) *http.Server {

	// preparing the server
	if err := vaultrdbstore.DB.OpenDB(t.TempDir()+"/tmp.vault.db", nil); err != nil {
		t.Fatalf("creating new vault failed: %s", err.Error())
	}

	// create server
	var router *mux.Router = mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/test").Path("/list/{bucketpath}").Methods(http.MethodGet).HandlerFunc(api_v1_buckets.List)
	router.PathPrefix("/test").Path("/create/{bucketpath}").Methods(http.MethodPost).HandlerFunc(api_v1_buckets.Create)
	router.PathPrefix("/test").Path("/delete/{bucketpath}/bucket/{bucket}").Methods(http.MethodDelete).HandlerFunc(api_v1_buckets.Delete)

	var srv = (&http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	})

	go func() {
		err := srv.ListenAndServe()

		t.Logf("hosting http server: %s", srv.Addr)

		t.Cleanup(func() {
			srv.Close()
			t.Errorf("error with httpserv: %s", err.Error())
		})
	}()

	return srv
}

func Test_CreateListDelete(t *testing.T) {

	srv := prep(t)
	// preparing the tests

	var list = []struct {
		path            string
		bucketpath      string
		estimatedbucket string
	}{
		{path: "single", bucketpath: "@", estimatedbucket: "single"},
		{path: "multi.level", bucketpath: "multi", estimatedbucket: "level"},
		{path: "multi.level1", bucketpath: "multi", estimatedbucket: "level1"},
		{path: "multi2.level1", bucketpath: "multi2", estimatedbucket: "level1"},
		{path: "multi2.level2", bucketpath: "multi2", estimatedbucket: "level2"},
		{path: "multi2.level2", bucketpath: "multi2", estimatedbucket: "level2"},
		{path: "multi3.level.asdf.ghnb", bucketpath: "multi3.level.asdf", estimatedbucket: "ghnb"},
	}

	time.Sleep(time.Second)

	// running the tests
	for _, ll := range list {

		t.Run(fmt.Sprintf("testing-create_%s", ll.path), func(t *testing.T) {

			var url string = fmt.Sprintf("http://%s/test/create/%s", srv.Addr, ll.path)

			t.Logf("running post on (%s)", url)
			resp, err := http.DefaultClient.Post(url, "", nil)
			if err != nil {
				t.Fatalf("error making post request to (%s): %s", url, err.Error())
			}
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("code(%d): error creating buckets[%s]: %s", resp.StatusCode, ll.path, err.Error())
			}

			// checking the path

			url = fmt.Sprintf("http://%s/test/list/%s", srv.Addr, ll.bucketpath)
			t.Logf("running get on (%s)", url)
			resp, err = http.DefaultClient.Get(url)
			if err != nil {
				t.Fatalf("error making get request to (%s): %s", url, err.Error())
			}
			defer resp.Body.Close()

			var res helpers.StringValueJson
			if res.Receive(nil, resp.Body); err != nil {
				t.Fatalf("error parsing body into stringvaluejson object: %s", err.Error())
			}
			t.Logf("res:(%v)", res)

			var contains bool = false
			for i := range res.Values {
				if res.Values[i] == ll.estimatedbucket {
					contains = true
					break
				}
			}

			if !contains {
				t.Fatalf("[%v]string does not contain the estimated bucket [%s]", res.Values, ll.estimatedbucket)
			}

		})
	}
}

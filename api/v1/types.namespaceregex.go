package v1

import (
	"context"
	"regexp"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// struct which contains the information about the namespace regex
type NamespacesRegex struct {

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Name string `json:"name,omitempty"`

	// +kubebuilder:default=[]
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	MustAvoidRegex []string `json:"mustavoidregex"`

	// +kubebuilder:default=[]
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	MustMatchRegex []string `json:"mustmatchregex"`

	// +kubebuilder:validation:Enum={"Secret","ConfigMap"}
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default="Secret"
	Kind string `json:"kind,omitempty"`

	// +kubebuilder:validation:Enum={"Opaque","kubernetes.io/service-account-token","kubernetes.io/dockercfg","kubernetes.io/dockerconfigjson","kubernetes.io/basic-auth","kubernetes.io/ssh-auth","kubernetes.io/tls","bootstrap.kubernetes.io/token"}
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default="Opaque"
	Type string `json:"type,omitempty"`
}

// get two lists of namespaces
//
// the 1. list contains all namespaces, which should be used in the creation process
//
// the 2. list, contains all namespaces, which shuold be avoided in the creation process
// and the already existing entities should be removed from
func (nsr NamespacesRegex) CalculateNamespaces(l logr.Logger, ctx context.Context, c client.Client) (mustMatch, mustAvoid []v1.Namespace, err error) {

	var namespaceList = &v1.NamespaceList{}

	// create the compare function
	// check whether a string exists in a list of regexpressions or not
	stringMatchesRegExpList := func(comp string, regexpList []string) (matched bool, err error) {

		for i := range regexpList {

			// if the matchstring function fails, the response will be false, error
			// else, the error will be return, or the "true" value
			if matched, err = regexp.MatchString(regexpList[i], comp); err != nil || matched {
				return
			}
		}

		// since no string in the regexpList matches the string comp
		// false and error nil will be returned
		return false, nil
	}

	l.Info("calculating namespaces for the following lists", "NamespacesRegex", nsr)

	if err = c.List(ctx, namespaceList, &client.ListOptions{}); err == nil {

		// parse through all registered namespaces
		for i := range namespaceList.Items {

			// check, if the namespace has to be avoided during deployment
			var inList bool = false
			if inList, err = stringMatchesRegExpList(namespaceList.Items[i].Name, nsr.MustAvoidRegex); err != nil {

				l.Error(err, "error calculating avoids", "current namespace", namespaceList.Items[i].Name, "MustAvoidRegex", nsr.MustAvoidRegex)

				return

			} else {
				if inList {
					// if the namespace is in the list [MustAvoidRegex], append the namespace to the list [mustAvoid]
					// and calculate the next namespace
					mustAvoid = append(mustAvoid, namespaceList.Items[i])
					continue
				}
			}

			// if the namespace is not in the list [MustAvoidRegex], check if the namespace is in the list [MustMatchRegex]
			if inList, err = stringMatchesRegExpList(namespaceList.Items[i].Name, nsr.MustMatchRegex); err != nil {

				l.Error(err, "error calculating matches", "current namespace", namespaceList.Items[i].Name, "MustMatchRegex", nsr.MustMatchRegex)

				return

			} else {

				if inList {
					// if the namespace is in the list [MustMatchRegex], then append the namespace to
					// the namespaces [mustMatch]
					mustMatch = append(mustMatch, namespaceList.Items[i])

				} else {
					// if the namespace also is not in the list [MustMatchRegex], then append the namespace to
					// the namespaces [mustAvoid]
					mustAvoid = append(mustAvoid, namespaceList.Items[i])
				}
			}
		}
	}
	return
}

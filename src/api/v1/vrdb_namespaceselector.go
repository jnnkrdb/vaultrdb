/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"context"
	"fmt"
	"regexp"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// VRDBNamespaceSelector defines the Namespaces, the operator should be
// looking for, while distributing the child objects to the cluster
type VRDBNamespaceSelector struct {

	// +kubebuilder:default={}
	// +operator-sdk:csv:customresourcedefinitions:type=namespaceSelector
	Avoid []string `json:"rx.avoid"`

	// +kubebuilder:default={}
	// +operator-sdk:csv:customresourcedefinitions:type=namespaceSelector
	Match []string `json:"rx.match"`
}

// validate the namespace selector for valid regex'
func (nselect VRDBNamespaceSelector) Validate() error {

	for _, str := range append(nselect.Avoid, nselect.Match...) {

		if len(str) == 0 {
			return fmt.Errorf("regex cannot be empty")
		}

		if _, err := regexp.Compile(str); err != nil {
			return fmt.Errorf("regex compilation error: %v", err)
		}
	}

	return nil
}

// calculates two collections of namespaces, which should weither be avoided
// or provided with the requested object
func (namespaceSelector VRDBNamespaceSelector) CalculateCollections(ctx context.Context, c client.Client) (avoids []string, matches []string, err error) {

	var (
		_log          = log.FromContext(ctx).WithValues("avoid", namespaceSelector.Avoid, "match", namespaceSelector.Match)
		allNamespaces = &v1.NamespaceList{}
	)

	_log.Info("calculating namespaces with privided namespace regex lists")

	// request a list of all accessable namespaces, to calculate the matching regex from
	if err = c.List(ctx, allNamespaces, &client.ListOptions{}); err != nil {
		_log.Error(err, "error requesting list of namespaces")
		return
	}

	// parse through every namespace and compare the namespace
	// with the given lists of required matches and avoids
	//
	// avoids are higher prioritized than matches, which means, if a namespace
	// matches with regexes in the avoids-list and the matches-list, then
	// the namespace will still be avoided
	//
	// if a namespace matches with no regex, weither from avoids nor matches,
	// the namespace will still be added to the avoids
	for _, ns := range allNamespaces.Items {

		if regexListContainsItem(ctx, ns.Name, namespaceSelector.Avoid, true) {
			avoids = append(avoids, ns.Name)
			continue
		}

		if regexListContainsItem(ctx, ns.Name, namespaceSelector.Match, false) {
			matches = append(matches, ns.Name)
			continue
		}

		avoids = append(avoids, ns.Name)
	}

	return
}

// the actual filter function
//
// parses through the given list and returns true, if the
// list contains a regex, which matches with the given namespace
//
// if an error occurs, it logs the error and returns the specifie
// default return value
func regexListContainsItem(ctx context.Context, namespace string, regexpList []string, defaultIfError bool) bool {
	var _log = log.FromContext(ctx).WithValues("checkedNamespace", namespace, "defaultResponseIfError", defaultIfError)

	// parse through the list of regexes and compile them, then
	// check for matchings with the given namespace
	for _, regExpression := range regexpList {

		// if an error occurs, thats most likely, to a not compilable regexp
		// even though they were checked, before creation
		match, err := regexp.MatchString(regExpression, namespace)
		if err != nil {
			_log.Error(err, "error checking the namespace with the given regex, returning default response", "regex", regExpression)
			return defaultIfError
		}

		// if the regexpression matches, the function will response with true
		if match {
			return true
		}
	}

	// no regex from the list, does match the given namespace
	return false
}

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
	"regexp"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// the function uses an interface, which requires the Specs-Fields, to have two functions, which
// return the string-arrays of the regexes, one for the matches and one for the avoids
func CalculateNamespaces(
	ctx context.Context,
	c client.Client,
	nsList interface {
		GetAvoidingRegexList() []string
		GetMatchingRegexList() []string
	}) ([]v1.Namespace, []v1.Namespace, error) {

	var (
		_log          = log.FromContext(ctx).WithValues("avoid", nsList.GetAvoidingRegexList(), "match", nsList.GetMatchingRegexList())
		allNamespaces = &v1.NamespaceList{}
		avoids        []v1.Namespace
		matches       []v1.Namespace
	)

	_log.Info("calculating namespaces with privided namespace regex lists")

	// request a list of all accessable namespaces, to calculate the matching regex from
	if err := c.List(ctx, allNamespaces, &client.ListOptions{}); err != nil {
		_log.Error(err, "error requesting list of namespaces")
		return nil, nil, err
	}

	// parse through every namespace and compare the namespace
	// with the given lists of required matches and avoids
	//
	// avoids are higher prioritized than matches, which means, if a namespace
	// matches with regexes in the avoids-list and the matches-list, then
	// the namespace will still be avoided
	for _, ns := range allNamespaces.Items {

		// first compare the avoiding regexes, so the comparision of matching
		// regexes can be skipped, if the namespace is in the avoid-regexes
		var inList, e = compare(ns.Name, nsList.GetAvoidingRegexList())
		if e != nil {
			_log.Error(e, "error calculating avoids", "current_namespace", ns.Name)
			return nil, nil, e
		} else if inList {
			avoids = append(avoids, ns)
			continue
		}

		// compare the matching regexes with the namespace, if the namespace is not
		// in the matching-regexes-list, the namespace will be added to the avoids
		inList, e = compare(ns.Name, nsList.GetMatchingRegexList())
		if e != nil {
			_log.Error(e, "error calculating matches", "current_namespace", ns.Name)
			return nil, nil, e
		} else if inList {
			matches = append(matches, ns)
		} else {
			avoids = append(avoids, ns)
		}
	}

	return avoids, matches, nil
}

// parse through every configured regex to calculate the namespaces which match
// with the given regex
func compare(comp string, regexpList []string) (bool, error) {
	for i := range regexpList {
		if matched, err := regexp.MatchString(regexpList[i], comp); err != nil || matched {
			return true, err
		}
	}

	return false, nil
}

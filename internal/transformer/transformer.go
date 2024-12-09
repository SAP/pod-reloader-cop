/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and pod-reloader-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package transformer

import (
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"

	operatorv1alpha1 "github.com/sap/pod-reloader-cop/api/v1alpha1"
)

type transformer struct{}

func NewParameterTransformer() *transformer {
	return &transformer{}
}

func (t *transformer) TransformParameters(namespace string, name string, parameters componentoperatorruntimetypes.Unstructurable) (componentoperatorruntimetypes.Unstructurable, error) {
	s := parameters.(*operatorv1alpha1.PodReloaderSpec)
	v := parameters.ToUnstructured()

	v["fullnameOverride"] = name

	if s.Image.PullSecret != "" {
		v["imagePullSecrets"] = []any{map[string]any{"name": s.Image.PullSecret}}
		delete(v["image"].(map[string]any), "pullSecret")
	}

	if s.ObjectSelector != nil || s.NamespaceSelector != nil {
		v["webhook"] = make(map[string]any)
		if s.ObjectSelector != nil {
			v["webhook"].(map[string]any)["objectSelector"] = v["objectSelector"]
		}
		if s.NamespaceSelector != nil {
			v["webhook"].(map[string]any)["namespaceSelector"] = v["namespaceSelector"]
		}
	}
	delete(v, "objectSelector")
	delete(v, "namespaceSelector")

	delete(v, "namespace")
	delete(v, "name")

	return componentoperatorruntimetypes.UnstructurableMap(v), nil
}

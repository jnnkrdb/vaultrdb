# VaultRDB

## Install to Cluster

To install the operator, use the following manifests:

- Namespace
```yaml
apiVersion: v1
kind: Namespace
metadata:
labels:
    control-plane: controller-manager
name: vaultrdb
```

- Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: vaultrdb
  namespace: vaultrdb
  labels:
    control-plane: controller-manager
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
  replicas: 1
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      securityContext:
        runAsNonRoot: true
      containers:
        image: ghcr.io/jnnkrdb/vaultrdb:latest
        name: vaultrdb
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
              - "ALL"
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8081
          initialDelaySeconds: 15
          periodSeconds: 20
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8081
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          limits:
            cpu: 500m
            memory: 128Mi
          requests:
            cpu: 10m
            memory: 64Mi
      serviceAccountName: vaultrdb-sa
      terminationGracePeriodSeconds: 10
```

- ServiceAccount
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    control-plane: controller-manager
  name: vaultrdb-sa
  namespace: vaultrdb
```

- ClusterRoleBinding
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    control-plane: controller-manager
  name: vaultrdb-crb
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: vaultrdb-cr
subjects:
- kind: ServiceAccount
  name: vaultrdb-sa
  namespace: vaultrdb
```

- ClusterRole
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: vaultrdb-cr
  labels:
    control-plane: controller-manager
rules:
- apiGroups: [ "" ]
  resources: [ namespaces ]
  verbs: [ get, list, watch ]

- apiGroups: [ jnnkrdb.de ]
  resources: [ vrdbconfigs, vrdbrequests, vrdbsecrets ]
  verbs: [ create, delete, get, list, patch, update, watch ]

- apiGroups: [ jnnkrdb.de ]
  resources: [ vrdbconfigs/finalizers, vrdbrequests/finalizers, vrdbsecrets/finalizers ]
  verbs: [ update ]

- apiGroups: [ jnnkrdb.de ]
  resources: [ vrdbconfigs/status, vrdbrequests/status, vrdbsecrets/status ]
  verbs: [ get, patch, update ]

- apiGroups: [ "" ]
  resources: [ configmaps, secrets ]
  verbs: [ create, delete, get, list, patch, update, watch ]

- apiGroups: [ coordination.k8s.io ]
  resources: [ leases ]
  verbs: [ create, delete, get, list, patch, update, watch ]
  
- apiGroups: [ "" ]
  resources: [ events ]
  verbs: [ create, patch ]
```

- Service
```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    control-plane: controller-manager
  name: vaultrdb
  namespace: vaultrdb
spec:
  ports:
    - port: 443
      protocol: TCP
      targetPort: 9443
  selector:
    control-plane: controller-manager
```

- ValidatingWebhookConfiguration
```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  creationTimestamp: null
  name: validating-webhook-configuration-vaultrdb
webhooks:
- admissionReviewVersions: [ v1 ]
  clientConfig:
    service:
      name: vaultrdb
      namespace: vaultrdb
      path: /validate-jnnkrdb-de-v1-vrdbconfig
  failurePolicy: Fail
  name: vvrdbconfig.kb.io
  rules:
  - apiGroups: [ jnnkrdb.de ]
    apiVersions: [ v1 ]
    operations: [ CREATE, UPDATE ]
    resources: [ vrdbconfigs ]
  sideEffects: None
- admissionReviewVersions: [ v1 ]
  clientConfig:
    service:
      name: vaultrdb
      namespace: vaultrdb
      path: /validate-jnnkrdb-de-v1-vrdbrequest
  failurePolicy: Fail
  name: vvrdbrequest.kb.io
  rules:
  - apiGroups: [ jnnkrdb.de ]
    apiVersions: [ v1 ]
    operations: [ CREATE, UPDATE ]
    resources: [ vrdbrequests ]
  sideEffects: None
- admissionReviewVersions: [ v1 ]
  clientConfig:
    service:
      name: vaultrdb
      namespace: vaultrdb
      path: /validate-jnnkrdb-de-v1-vrdbsecret
  failurePolicy: Fail
  name: vvrdbsecret.kb.io
  rules:
  - apiGroups: [ jnnkrdb.de ]
    apiVersions: [ v1 ]
    operations: [ CREATE, UPDATE ]
    resources: [ vrdbsecrets ]
  sideEffects: None
```

- MutatingWebhookConfiguration
```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  creationTimestamp: null
  name: mutating-webhook-configuration-vaultrdb
webhooks:
- admissionReviewVersions: [ v1 ]
  clientConfig:
    service:
      name: vaultrdb
      namespace: vaultrdb
      path: /mutate-jnnkrdb-de-v1-vrdbconfig
  failurePolicy: Fail
  name: mvrdbconfig.kb.io
  rules:
  - apiGroups: [ jnnkrdb.de ]
    apiVersions: [ v1 ]
    operations: [ CREATE, UPDATE ]
    resources: [ vrdbconfigs ]
  sideEffects: None
- admissionReviewVersions: [ v1 ]
  clientConfig:
    service:
      name: vaultrdb
      namespace: vaultrdb
      path: /mutate-jnnkrdb-de-v1-vrdbrequest
  failurePolicy: Fail
  name: mvrdbrequest.kb.io
  rules:
  - apiGroups: [ jnnkrdb.de ]
    apiVersions: [ v1 ]
    operations: [ CREATE, UPDATE ]
    resources: [ vrdbrequests ]
  sideEffects: None
- admissionReviewVersions: [ v1 ]
  clientConfig:
    service:
      name: vaultrdb
      namespace: vaultrdb
      path: /mutate-jnnkrdb-de-v1-vrdbsecret
  failurePolicy: Fail
  name: mvrdbsecret.kb.io
  rules:
  - apiGroups: [ jnnkrdb.de ]
    apiVersions: [ v1 ]
    operations: [ CREATE, UPDATE ]
    resources: [ vrdbsecrets ]
  sideEffects: None
```

- CustomResourceDefinition - VRDBConfig
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.1
  labels:
    control-plane: controller-manager
  name: vrdbconfigs.jnnkrdb.de
spec:
  group: jnnkrdb.de
  names:
    kind: VRDBConfig
    listKind: VRDBConfigList
    plural: vrdbconfigs
    singular: vrdbconfig
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: VRDBConfig is the Schema for the vrdbconfigs API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          data:
            additionalProperties:
              type: string
            type: object
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          namespaceSelector:
            description: VRDBNamespaceSelector defines the Namespaces, the operator
              should be looking for, while distributing the child objects to the cluster
            properties:
              rx.avoid:
                items:
                  type: string
                type: array
              rx.match:
                items:
                  type: string
                type: array
            required:
            - rx.avoid
            - rx.match
            type: object
          status:
            properties:
              namespaces:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
```

- CustomResourceDefinition - VRDBRequest
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.1
  labels:
    control-plane: controller-manager
  name: vrdbrequests.jnnkrdb.de
spec:
  group: jnnkrdb.de
  names:
    kind: VRDBRequest
    listKind: VRDBRequestList
    plural: vrdbrequests
    singular: vrdbrequest
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: VRDBRequest is the Schema for the vrdbrequests API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          data:
            additionalProperties:
              type: string
            type: object
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          namespaceSelector:
            description: VRDBNamespaceSelector defines the Namespaces, the operator
              should be looking for, while distributing the child objects to the cluster
            properties:
              rx.avoid:
                items:
                  type: string
                type: array
              rx.match:
                items:
                  type: string
                type: array
            required:
            - rx.avoid
            - rx.match
            type: object
          status:
            properties:
              namespaces:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
```

- CustomResourceDefinition - VRDBSecret
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.1
  labels:
    control-plane: controller-manager
  name: vrdbsecrets.jnnkrdb.de
spec:
  group: jnnkrdb.de
  names:
    kind: VRDBSecret
    listKind: VRDBSecretList
    plural: vrdbsecrets
    singular: vrdbsecret
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: VRDBSecret is the Schema for the vrdbsecrets API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          data:
            additionalProperties:
              type: string
            type: object
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          namespaceSelector:
            description: VRDBNamespaceSelector defines the Namespaces, the operator
              should be looking for, while distributing the child objects to the cluster
            properties:
              rx.avoid:
                items:
                  type: string
                type: array
              rx.match:
                items:
                  type: string
                type: array
            required:
            - rx.avoid
            - rx.match
            type: object
          status:
            properties:
              namespaces:
                items:
                  type: string
                type: array
            type: object
          stringData:
            additionalProperties:
              type: string
            type: object
          type:
            type: string
        type: object
    served: true
    storage: true
    subresources:
      status: {}
```



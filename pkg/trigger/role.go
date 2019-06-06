package trigger

const servingRole = `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: {{}}
  name: serving-role
rules:
- apiGroups: ["serving.knative.dev"]
  resources: ["*"]
  verbs: ["get", "list", "create", "watch", "patch", "delete"]
`

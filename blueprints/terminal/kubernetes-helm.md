---
ide: terminal
appName: Terminal
description: Kubernetes Helm Chart App
---

helm create generates a complete Helm chart with Deployment, Service, Ingress, and HPA.
Deploy: helm install {name} ./{name} --namespace default --create-namespace

```json
[
  { "action": "shell", "cmd": "helm create {name}", "cwd": "{outputDir}", "timeout": 10000 }
]
```

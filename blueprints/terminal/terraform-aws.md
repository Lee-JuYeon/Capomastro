---
ide: terminal
appName: Terminal
description: Terraform Infrastructure as Code (AWS provider)
---

terraform init bootstraps a new Terraform project with AWS provider.
Configure credentials: aws configure or AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY env vars.
Plan: terraform plan. Apply: terraform apply.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/modules/{name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "terraform init", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

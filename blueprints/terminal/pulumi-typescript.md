---
ide: terminal
appName: Terminal
description: Pulumi IaC (TypeScript + AWS provider)
---

pulumi new aws-typescript scaffolds a Pulumi stack with the AWS TypeScript provider.
Configure: pulumi config set aws:region us-east-1. Deploy: pulumi up.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "pulumi new aws-typescript --name {name} --stack dev --yes", "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```

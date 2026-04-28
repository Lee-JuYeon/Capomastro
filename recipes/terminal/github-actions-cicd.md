---
ide: terminal
appName: Terminal
description: GitHub Actions CI/CD Pipeline
---

Creates .github/workflows/ with a CI workflow: build, test, lint on push/PR.
Extend with deploy.yml for staging/production deployment.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/.github/workflows", "cwd": "{outputDir}" }
]
```

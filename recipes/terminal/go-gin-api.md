---
ide: terminal
appName: Terminal
description: Go REST API (Gin + Go Modules)
---

go mod init creates the module; go get fetches Gin. Generates a minimal main.go with health endpoint.
Run with: go run . (listens on :8080)

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "go mod init {bundle}", "cwd": "{outputDir}/{name}", "timeout": 15000 },
  { "action": "shell", "cmd": "go get github.com/gin-gonic/gin@latest", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

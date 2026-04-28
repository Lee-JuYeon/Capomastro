---
ide: terminal
appName: Terminal
description: Go CLI Tool (Cobra)
---

cobra-cli init generates a Cobra-based CLI with subcommand support and help text.
Install cobra-cli first: go install github.com/spf13/cobra-cli@latest

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "go mod init {bundle}", "cwd": "{outputDir}/{name}", "timeout": 15000 },
  { "action": "shell", "cmd": "go get github.com/spf13/cobra@latest && cobra-cli init --pkg-name {bundle}", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

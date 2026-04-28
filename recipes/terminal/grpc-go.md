---
ide: terminal
appName: Terminal
description: gRPC Service (Go + Protocol Buffers)
---

Creates a Go gRPC service skeleton with protoc-gen-go installed.
Requires protoc compiler in PATH. Define .proto files in proto/ directory.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/proto {name}/server", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "go mod init {bundle}", "cwd": "{outputDir}/{name}", "timeout": 15000 },
  { "action": "shell", "cmd": "go get google.golang.org/grpc@latest google.golang.org/protobuf@latest", "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```

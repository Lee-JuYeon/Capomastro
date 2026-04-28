---
ide: terminal
appName: Terminal
description: MongoDB Database (Docker Compose + Mongo Express)
---

Docker Compose boots MongoDB 7 + Mongo Express web UI.
Connect: mongosh mongodb://root:root@localhost:27017
Mongo Express UI: http://localhost:8081

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" }
]
```

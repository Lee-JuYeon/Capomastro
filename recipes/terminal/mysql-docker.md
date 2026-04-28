---
ide: terminal
appName: Terminal
description: MySQL Database (Docker Compose)
---

Docker Compose boots MySQL 8.3 + Adminer web UI.
Connect: mysql -h 127.0.0.1 -P 3306 -u root -proot {name}
Adminer UI: http://localhost:8080

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" }
]
```

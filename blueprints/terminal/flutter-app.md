---
ide: terminal
appName: Terminal
description: Flutter Cross-Platform App (Dart + Flutter)
---

flutter create generates iOS/Android/Web/Desktop targets in one command.
Requires Flutter SDK in PATH. For a specific platform only, add --platforms=ios,android.

```json
[
  { "action": "shell", "cmd": "flutter create {name} --org {orgId} --platforms ios,android", "cwd": "{outputDir}", "timeout": 90000 }
]
```

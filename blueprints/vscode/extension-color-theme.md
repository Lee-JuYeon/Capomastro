---
ide: vscode
appName: Visual Studio Code
description: VSCode Color Theme Extension (yo code)
---

Requires: Node.js + `npm install -g yo generator-code` installed.
Variables: {name}=theme name, {orgId}=publisher id, {outputDir}=parent dir.

```json
[
  { "action": "shell", "cmd": "mkdir -p \"{outputDir}\"", "timeout": 3000 },
  { "action": "shell", "cmd": "yo code --extensionType=ext-colortheme --extensionName=\"{name}\" --extensionDisplayName=\"{name}\" --extensionDescription=\"\" --extensionVersion=0.0.1 --extensionPublisher=\"{orgId}\" --gitInit=false --openWithCode=false", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "code \"{outputDir}/{name}\"", "timeout": 5000 }
]
```

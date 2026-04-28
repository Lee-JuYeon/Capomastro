---
ide: vscode
appName: Visual Studio Code
description: VSCode Snippet Pack Extension (yo code)
---

Requires: Node.js + `npm install -g yo generator-code` installed.
Variables: {name}=snippet pack name, {orgId}=publisher id, {outputDir}=parent dir.

```json
[
  { "action": "shell", "cmd": "mkdir -p \"{outputDir}\"", "timeout": 3000 },
  { "action": "shell", "cmd": "yo code --extensionType=ext-snippets --extensionName=\"{name}\" --extensionDisplayName=\"{name}\" --extensionDescription=\"\" --extensionVersion=0.0.1 --extensionPublisher=\"{orgId}\" --gitInit=false --openWithCode=false", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "code \"{outputDir}/{name}\"", "timeout": 5000 }
]
```

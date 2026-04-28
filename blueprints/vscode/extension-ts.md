---
ide: vscode
appName: Visual Studio Code
description: VSCode TypeScript Extension (yo code)
---

Requires: Node.js + `npm install -g yo generator-code` installed.
Uses yo code non-interactive flags — no GUI wizard needed.
Variables: {name}=extension name, {orgId}=publisher id, {outputDir}=parent dir.

```json
[
  { "action": "shell", "cmd": "mkdir -p \"{outputDir}\"", "timeout": 3000 },
  { "action": "shell", "cmd": "yo code --extensionType=ext-command-ts --extensionName=\"{name}\" --extensionDisplayName=\"{name}\" --extensionDescription=\"\" --extensionVersion=0.0.1 --extensionPublisher=\"{orgId}\" --gitInit=false --pkgManager=npm --openWithCode=false", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "code \"{outputDir}/{name}\"", "timeout": 5000 }
]
```

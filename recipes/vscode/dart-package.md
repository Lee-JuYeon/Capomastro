---
ide: vscode
appName: Visual Studio Code
description: Dart Package via VSCode Flutter Extension
---

Cmd+Shift+P → "Flutter: New Project" → Package → enter name → pick folder.
For a plugin (with native code), select "Plugin" instead of "Package".

```json
[
  { "action": "shell", "cmd": "osascript -e 'tell application \"Visual Studio Code\" to activate'", "timeout": 3000 },
  { "action": "wait",  "ms": 800 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to keystroke \"p\" using {command down, shift down}'", "timeout": 2000 },
  { "action": "wait",  "ms": 500 },
  { "action": "filter", "text": "Flutter: New Project" },
  { "action": "wait",  "ms": 400 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to key code 36'", "timeout": 2000 },
  { "action": "wait",  "ms": 600 },
  { "action": "filter", "text": "Package" },
  { "action": "wait",  "ms": 300 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to key code 36'", "timeout": 2000 },
  { "action": "wait",  "ms": 500 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to keystroke \"a\" using {command down}'", "timeout": 2000 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to keystroke \"{name}\"'", "timeout": 2000 },
  { "action": "wait",  "ms": 300 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to key code 36'", "timeout": 2000 },
  { "action": "wait",  "ms": 800 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to keystroke \"g\" using {command down, shift down}'", "timeout": 2000 },
  { "action": "wait",  "ms": 400 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to keystroke \"{outputDir}\"'", "timeout": 2000 },
  { "action": "wait",  "ms": 300 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to key code 36'", "timeout": 2000 },
  { "action": "wait",  "ms": 500 },
  { "action": "shell", "cmd": "osascript -e 'tell application \"System Events\" to key code 36'", "timeout": 2000 }
]
```

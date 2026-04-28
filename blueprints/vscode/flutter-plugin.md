---
ide: vscode
appName: Visual Studio Code
description: Flutter Plugin (native platform code — iOS + Android) via VSCode
---

Cmd+Shift+P → "Flutter: New Project" → Plugin → enter name → pick folder.
Generates Kotlin (Android) + Swift (iOS) native method channel scaffold alongside Dart API.

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
  { "action": "filter", "text": "Plugin" },
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

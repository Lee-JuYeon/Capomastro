---
ide: vscode
appName: Visual Studio Code
description: Flutter App via VSCode Flutter Extension
---

Requires: VSCode + Flutter extension (Dart-Code.flutter) installed.
Flow: Cmd+Shift+P → "Flutter: New Project" → Application → enter name → pick save folder.
shell actions send key presses via osascript since VSCode quick-pick is not AX-clickable.

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
  { "action": "filter", "text": "Application" },
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

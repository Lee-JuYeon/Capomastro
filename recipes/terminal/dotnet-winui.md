---
ide: terminal
appName: Terminal
description: Windows WinUI 3 Desktop App (.NET + C#)
---

dotnet new winui generates a WinUI 3 XAML app. Requires Windows + .NET 8 SDK.
On macOS, generates the project structure only — build and run on a Windows machine.

```json
[
  { "action": "shell", "cmd": "dotnet new winui -n {name} -o {outputDir}/{name}", "cwd": "{outputDir}", "timeout": 30000 }
]
```

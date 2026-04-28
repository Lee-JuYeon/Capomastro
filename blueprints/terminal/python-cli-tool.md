---
ide: terminal
appName: Terminal
description: Python CLI Tool (Typer + Rich)
---

Creates a Typer CLI with Rich output formatting. Typer uses type annotations for argument definition.
Install and run: pip install -e . && {name} --help

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet typer rich", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

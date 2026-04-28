---
ide: terminal
appName: Terminal
description: Python FastAPI REST API
---

Creates a minimal FastAPI project: virtual environment + FastAPI + Uvicorn + pyproject.toml.
Activate with: source {name}/.venv/bin/activate && uvicorn main:app --reload

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 30000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet fastapi 'uvicorn[standard]' && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

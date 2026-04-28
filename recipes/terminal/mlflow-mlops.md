---
ide: terminal
appName: Terminal
description: MLflow MLOps Project (experiment tracking + model registry)
---

Creates an MLflow project with local tracking server.
Start UI: mlflow ui (port 5000). Logs experiments, parameters, metrics, and model artifacts.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/runs {name}/models", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet mlflow torch scikit-learn && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```

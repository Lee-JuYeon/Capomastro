---
ide: terminal
appName: Terminal
description: ML Inference Server (FastAPI + Hugging Face Transformers)
---

FastAPI + Transformers for model serving. Load any HuggingFace model via model_name env var.
Run with: uvicorn main:app --reload (listens on :8000).

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet fastapi 'uvicorn[standard]' transformers torch && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 180000 }
]
```

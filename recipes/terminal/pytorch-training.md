---
ide: terminal
appName: Terminal
description: PyTorch ML Training Project (Python + PyTorch + Lightning)
---

Creates a PyTorch Lightning training project with virtual environment.
GPU support: install pytorch with CUDA extras on Linux/Windows. MPS (Apple Silicon) works out of box.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/data {name}/models {name}/notebooks", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet torch torchvision lightning datasets transformers && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 180000 }
]
```

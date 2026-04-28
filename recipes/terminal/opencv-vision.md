---
ide: terminal
appName: Terminal
description: Computer Vision App (Python + OpenCV + YOLO)
---

Creates a computer vision project with OpenCV + Ultralytics YOLO for object detection.
Supports image, video, and webcam input. GPU optional (falls back to CPU).

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/input {name}/output {name}/models", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet opencv-python ultralytics Pillow numpy && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```

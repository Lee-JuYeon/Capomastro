---
ide: terminal
appName: Terminal
description: Speech / Audio App (Python + OpenAI Whisper + PyDub)
---

Whisper + PyDub for transcription, noise reduction, and audio processing.
Transcribe: whisper audio.mp3 --model base (auto-downloads model on first run).

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/audio {name}/transcripts", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet openai-whisper pydub soundfile numpy && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```

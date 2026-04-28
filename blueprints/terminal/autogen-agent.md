---
ide: terminal
appName: Terminal
description: AI Agent Framework (Python + AutoGen + MCP)
---

AutoGen multi-agent setup with a GroupChat manager and specialist agents.
Provider-agnostic: configure OpenAI / Anthropic / Ollama via config_list in .env.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/agents {name}/tools", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet 'pyautogen[all]' anthropic openai && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```

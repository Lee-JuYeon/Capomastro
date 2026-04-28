---
ide: terminal
appName: Terminal
description: LangChain RAG App (Python + LangChain + FAISS)
---

RAG scaffold: LangChain + FAISS vector store + OpenAI embeddings (swap provider via .env).
Structure: loader → splitter → embedder → retriever → chain.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/docs {name}/chains", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet langchain langchain-community langchain-openai faiss-cpu && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```

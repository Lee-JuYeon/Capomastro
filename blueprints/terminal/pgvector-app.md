---
ide: terminal
appName: Terminal
description: Vector Database App (PostgreSQL + pgvector + LangChain)
---

pgvector extension on PostgreSQL for semantic search. LangChain PGVector integration included.
Requires: Supabase local or standalone PostgreSQL with pgvector extension enabled.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "python3 -m venv {outputDir}/{name}/.venv", "cwd": "{outputDir}", "timeout": 20000 },
  { "action": "shell", "cmd": "{outputDir}/{name}/.venv/bin/pip install --quiet pgvector psycopg2-binary langchain-community openai && {outputDir}/{name}/.venv/bin/pip freeze > {outputDir}/{name}/requirements.txt", "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```

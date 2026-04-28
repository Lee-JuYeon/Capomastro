---
ide: terminal
appName: Terminal
description: Supabase Local Dev (PostgreSQL + Auth + Storage + Edge Functions)
---

supabase init + supabase start boots a full Supabase stack locally via Docker.
Studio UI: http://localhost:54323. DB: postgresql://postgres:postgres@localhost:54322/postgres

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "supabase init", "cwd": "{outputDir}/{name}", "timeout": 15000 },
  { "action": "shell", "cmd": "supabase start", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```

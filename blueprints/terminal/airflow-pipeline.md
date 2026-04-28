---
ide: terminal
appName: Terminal
description: Apache Airflow Data Pipeline (Python + Astronomer Cosmos)
---

astro dev init (Astronomer CLI) scaffolds a local Airflow project with Docker Compose.
Start: astro dev start. DAGs go in /dags, plugins in /plugins.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "cd {outputDir}/{name} && astro dev init", "cwd": "{outputDir}/{name}", "timeout": 30000 }
]
```

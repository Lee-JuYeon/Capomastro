---
ide: terminal
appName: Terminal
description: Log Collection Stack (Elasticsearch + Logstash + Kibana)
---

Docker Compose boots ELK stack. Kibana UI: http://localhost:5601.
Ship logs via Logstash TCP input or Filebeat agent.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/logstash/pipeline {name}/logstash/config", "cwd": "{outputDir}" }
]
```

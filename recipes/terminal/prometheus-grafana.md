---
ide: terminal
appName: Terminal
description: Monitoring Stack (Prometheus + Grafana + Node Exporter)
---

Docker Compose boots Prometheus (metrics collection) + Grafana (dashboards) + Node Exporter.
Grafana UI: http://localhost:3000 (admin/admin). Prometheus UI: http://localhost:9090

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}/prometheus {name}/grafana/dashboards {name}/grafana/datasources", "cwd": "{outputDir}" }
]
```

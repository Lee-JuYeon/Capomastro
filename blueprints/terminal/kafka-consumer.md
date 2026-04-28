---
ide: terminal
appName: Terminal
description: Kafka Consumer / Producer (Node.js + KafkaJS)
---

Creates a KafkaJS consumer + producer project with TypeScript.
Broker URL defaults to localhost:9092 — override via KAFKA_BROKERS env var.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "npm init -y && npm install kafkajs && npm install -D typescript ts-node-dev @types/node", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

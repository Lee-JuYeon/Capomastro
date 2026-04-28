---
ide: terminal
appName: Terminal
description: Java Spring Boot REST API (Maven + Java 21)
---

Spring Initializr CLI (curl) downloads a ready-to-run Maven project with Spring Web + Actuator.
Run with: ./mvnw spring-boot:run (listens on :8080). Requires Java 21+ in PATH.

```json
[
  { "action": "shell", "cmd": "curl -s 'https://start.spring.io/starter.tgz?type=maven-project&language=java&bootVersion=3.2.0&baseDir={name}&groupId={orgId}&artifactId={name}&name={name}&packageName={bundle}&dependencies=web,actuator' -o /tmp/{name}.tgz && tar -xzf /tmp/{name}.tgz -C {outputDir}", "cwd": "{outputDir}", "timeout": 30000 }
]
```

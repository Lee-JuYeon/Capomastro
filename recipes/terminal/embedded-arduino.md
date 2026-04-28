---
ide: terminal
appName: Terminal
description: Embedded / IoT Project (Arduino / ESP32 + PlatformIO)
---

PlatformIO CLI scaffold for Arduino-compatible boards (ESP32, STM32, Arduino Uno, etc.).
pio run builds. pio run --target upload flashes to connected device.

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "pio project init --board esp32dev", "cwd": "{outputDir}/{name}", "timeout": 60000 }
]
```

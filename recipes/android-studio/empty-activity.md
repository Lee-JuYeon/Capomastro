---
ide: android-studio
appName: Android Studio
description: Android Empty Activity (Kotlin + Views)
---

Android Studio new project: File → New → New Project → Phone and Tablet tab → Empty Views Activity → Next → fill → Finish.
Language: Kotlin (default). Minimum SDK: API 24 (default). Build system: Gradle (Kotlin DSL).
NOTE: Android Studio uses IntelliJ-based Swing UI on macOS — AX titles may differ; fallback to LLM on failure.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Project…"] },
  { "action": "click_tab",  "title": "Phone and Tablet" },
  { "action": "click_cell", "title": "Empty Views Activity" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Package name",           "value": "{bundle}" },
  { "action": "fill_field", "title": "Save location",          "value": "{outputDir}/{name}" },
  { "action": "click",      "title": "Finish" }
]
```

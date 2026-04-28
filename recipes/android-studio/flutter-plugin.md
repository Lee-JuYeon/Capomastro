---
ide: android-studio
appName: Android Studio
description: Flutter Plugin (native platform channels) via Android Studio
---

New Flutter Project → Plugin template → fill name, org, location → Create.
Generates method channel scaffolding for Android (Kotlin) and iOS (Swift) alongside Dart API layer.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Flutter Project…"] },
  { "action": "click_cell", "title": "Flutter" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Project name",      "value": "{name}" },
  { "action": "fill_field", "title": "Project location",  "value": "{outputDir}/{name}" },
  { "action": "click_cell", "title": "Plugin" },
  { "action": "click",      "title": "Create" }
]
```

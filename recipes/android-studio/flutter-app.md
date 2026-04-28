---
ide: android-studio
appName: Android Studio
description: Flutter App via Android Studio Flutter Plugin
---

Requires: Android Studio + Flutter plugin + Dart plugin installed.
File → New → New Flutter Project → Flutter → fill SDK path, name, location → Finish.
The Flutter plugin adds a "Flutter" entry in the new project wizard alongside native Android templates.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Flutter Project…"] },
  { "action": "click_cell", "title": "Flutter" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Project name",      "value": "{name}" },
  { "action": "fill_field", "title": "Project location",  "value": "{outputDir}/{name}" },
  { "action": "click",      "title": "Create" }
]
```

---
ide: android-studio
appName: Android Studio
description: Android ARCore App (Augmented Reality)
---

Phone and Tablet tab → filter "AR" → select "AR Core Empty Activity" → Next → fill → Finish.
Includes ARCore SDK (com.google.ar:core) and required camera permissions in AndroidManifest.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Project…"] },
  { "action": "click_tab",  "title": "Phone and Tablet" },
  { "action": "filter",     "text": "AR" },
  { "action": "click_cell", "title": "AR Core Empty Activity" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Package name",           "value": "{bundle}" },
  { "action": "fill_field", "title": "Save location",          "value": "{outputDir}/{name}" },
  { "action": "click",      "title": "Finish" }
]
```

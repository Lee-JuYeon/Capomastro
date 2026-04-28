---
ide: android-studio
appName: Android Studio
description: Android Jetpack Compose Activity
---

Phone and Tablet tab → Empty Activity (Compose) → Next → fill → Finish.
"Empty Activity" in newer Android Studio versions defaults to Compose. If the template name is
"Empty Views Activity", scroll down to find the Compose variant or use filter.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Project…"] },
  { "action": "click_tab",  "title": "Phone and Tablet" },
  { "action": "click_cell", "title": "Empty Activity" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Package name",           "value": "{bundle}" },
  { "action": "fill_field", "title": "Save location",          "value": "{outputDir}/{name}" },
  { "action": "click",      "title": "Finish" }
]
```

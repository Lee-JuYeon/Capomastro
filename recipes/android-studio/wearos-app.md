---
ide: android-studio
appName: Android Studio
description: Wear OS App (Kotlin + Compose for Wear)
---

Wear OS tab → Blank Tile template → Next → fill → Finish.
For a watch face, select "Watch Face" cell instead. Compose for Wear OS (androidx.wear.compose) included.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Project…"] },
  { "action": "click_tab",  "title": "Wear OS" },
  { "action": "click_cell", "title": "Blank Tile" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Package name",           "value": "{bundle}" },
  { "action": "fill_field", "title": "Save location",          "value": "{outputDir}/{name}" },
  { "action": "click",      "title": "Finish" }
]
```

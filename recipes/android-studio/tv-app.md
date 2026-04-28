---
ide: android-studio
appName: Android Studio
description: Android TV App (Kotlin + Leanback)
---

TV tab → Android TV Activity → Next → fill → Finish.
Includes Leanback library + BrowseFragment/PlaybackFragment scaffold.

```json
[
  { "action": "menu",       "path": ["File", "New", "New Project…"] },
  { "action": "click_tab",  "title": "TV" },
  { "action": "click_cell", "title": "Android TV Activity" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Package name",           "value": "{bundle}" },
  { "action": "fill_field", "title": "Save location",          "value": "{outputDir}/{name}" },
  { "action": "click",      "title": "Finish" }
]
```

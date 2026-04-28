---
ide: unreal
appName: Unreal Editor
description: Unreal Engine Blank Game (C++)
---

Unreal Project Browser → Games category → Blank template → set project name + location → Create.
Default: C++ project, Desktop target, Maximum quality, Starter Content excluded.
NOTE: Unreal Project Browser AX tree is best-guess; validate on first run.

```json
[
  { "action": "click_cell", "title": "Games" },
  { "action": "click_cell", "title": "Blank" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Folder",                 "value": "{outputDir}" },
  { "action": "click",      "title": "Create" }
]
```

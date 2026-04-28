---
ide: unreal
appName: Unreal Editor
description: Unreal Engine First Person Shooter (C++)
---

Games → First Person template → set name + location → Create.
Includes FirstPersonCharacter BP, FirstPersonProjectile, FirstPersonHUD.

```json
[
  { "action": "click_cell", "title": "Games" },
  { "action": "click_cell", "title": "First Person" },
  { "action": "fill_field", "title": "Name",                   "value": "{name}" },
  { "action": "fill_field", "title": "Folder",                 "value": "{outputDir}" },
  { "action": "click",      "title": "Create" }
]
```

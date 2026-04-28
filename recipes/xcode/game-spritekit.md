---
ide: xcode
appName: Xcode
description: iOS SpriteKit Game
---

iOS tab → filter "Game" → Game template → Next → fill options (tech defaults to SpriteKit) → Next → save → Create.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "filter",     "text": "Game" },
  { "action": "click_cell", "title": "Game" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

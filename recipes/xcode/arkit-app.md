---
ide: xcode
appName: Xcode
description: iOS ARKit App (Augmented Reality)
---

iOS tab → filter "Augmented Reality" → Augmented Reality App → Next → fill options → Next → save → Create.
Default: SceneKit. Switch to RealityKit in project settings post-creation.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "filter",     "text": "Augmented Reality" },
  { "action": "click_cell", "title": "Augmented Reality App" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

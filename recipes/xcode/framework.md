---
ide: xcode
appName: Xcode
description: iOS Framework
---

iOS tab → filter "Framework" → select Framework cell → Next → fill options → Next → save → Create.
Produces a .framework dynamic library. For static library, filter "Library" instead.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "filter",     "text": "Framework" },
  { "action": "click_cell", "title": "Framework" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

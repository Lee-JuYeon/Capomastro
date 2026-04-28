---
ide: xcode
appName: Xcode
description: watchOS App
---

watchOS tab → Watch App template → Next → fill options → Next → save → Create.
For a companion iPhone app, use "Watch App with Companion iPhone App" template.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "watchOS" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

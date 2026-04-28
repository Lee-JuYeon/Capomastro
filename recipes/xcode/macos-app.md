---
ide: xcode
appName: Xcode
description: macOS SwiftUI App
---

macOS tab → App template → Next → fill options → Next → save → Create.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "macOS" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

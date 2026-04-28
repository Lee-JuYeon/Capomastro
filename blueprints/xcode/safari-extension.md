---
ide: xcode
appName: Xcode
description: Safari Web Extension (macOS + iOS)
---

macOS tab → filter "Safari" → Safari Extension App → Next → fill options → Next → save → Create.
Generates Swift host app + JS content scripts. Works on both macOS and iOS Safari.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "macOS" },
  { "action": "filter",     "text": "Safari" },
  { "action": "click_cell", "title": "Safari Extension App" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

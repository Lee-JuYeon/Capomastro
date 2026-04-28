---
ide: xcode
appName: Xcode
description: iOS Document-Based App
---

iOS tab → filter "Document" → Document App → Next → fill options → Next → save → Create.
Register UTI + file extension in Info.plist post-creation.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "filter",     "text": "Document" },
  { "action": "click_cell", "title": "Document App" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

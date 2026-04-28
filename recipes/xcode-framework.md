---
ide: xcode
appName: Xcode
description: iOS Framework
---

```json
[
  { "action": "menu", "path": ["File", "New", "Project…"] },
  { "action": "click_tab", "title": "iOS" },
  { "action": "filter", "text": "Framework" },
  { "action": "click_cell", "title": "Framework" },
  { "action": "click", "title": "Next" },
  { "action": "fill_field", "title": "Product Name", "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "set_location" },
  { "action": "click", "title": "Create" }
]
```

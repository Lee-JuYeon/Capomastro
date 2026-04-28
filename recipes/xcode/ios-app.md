---
ide: xcode
appName: Xcode
description: iOS SwiftUI App
---

Two-page Xcode wizard: (1) template chooser → Next → (2) options form (Name, OrgId) → Next → (3) save location (Where + Create).
iOS tab is default — clicking it is a safe no-op if already active.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

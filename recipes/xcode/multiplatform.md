---
ide: xcode
appName: Xcode
description: Multiplatform SwiftUI App
---

Multiplatform tab → App template (only one; auto-selected) → Next → fill options → Next → save → Create.
SUPPORTED_PLATFORMS: iphoneos iphonesimulator macosx xros xrsimulator.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "Multiplatform" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

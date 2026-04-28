---
ide: xcode
appName: Xcode
description: Multiplatform SwiftUI App
---

Xcode new project: Multiplatform tab → App (only one template; auto-selected) → Next → fill form → save → Create.
Multiplatform builds for iPhone + iPad + Mac (Designed for iPad) in one target.
SUPPORTED_PLATFORMS includes iphoneos iphonesimulator macosx xros xrsimulator.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "Multiplatform" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",              "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier",   "value": "{orgId}" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

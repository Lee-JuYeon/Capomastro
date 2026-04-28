---
ide: xcode
appName: Xcode
description: iOS SwiftUI App
---

Xcode new project flow: File → New → Project → iOS tab → App → Next → fill form → save → Create.
iOS tab is selected by default; clicking it again is a safe no-op if already active.
Language: Swift / Interface: SwiftUI are Xcode defaults — no need to change them.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",              "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier",   "value": "{orgId}" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

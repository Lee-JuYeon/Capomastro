---
ide: xcode
appName: Xcode
description: Swift Package (SPM)
---

File → New → Package opens a NSSavePanel (not a project sheet). The filename field has no explicit
AXTitle — fill_field matches the first AXTextField found. Xcode creates Package.swift + Sources/
automatically.

```json
[
  { "action": "menu",       "path": ["File", "New", "Package…"] },
  { "action": "fill_field", "title": "Save As",                 "value": "{name}" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

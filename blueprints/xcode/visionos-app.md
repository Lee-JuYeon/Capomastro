---
ide: xcode
appName: Xcode
description: visionOS App (Apple Vision Pro)
---

visionOS tab → App template (Window + ImmersiveSpace) → Next → fill options → Next → save → Create.
RealityKit + SwiftUI included by default.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "visionOS" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

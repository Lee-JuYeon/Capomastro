---
ide: xcode
appName: Xcode
description: tvOS SwiftUI App
---

tvOS tab → App template → Next → fill options → Next → save → Create.
Default: SwiftUI. TVUIKit focus engine available post-creation.

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "tvOS" },
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

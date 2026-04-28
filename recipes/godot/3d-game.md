---
ide: godot
appName: Godot
description: Godot 3D Game (GDScript)
---

New → fill name + path → set renderer to "Forward+" for 3D → Create & Edit.
For lower-spec hardware, use "Compatibility" renderer instead.

```json
[
  { "action": "click",      "title": "New" },
  { "action": "fill_field", "title": "Project Name",           "value": "{name}" },
  { "action": "fill_field", "title": "Project Path",           "value": "{outputDir}/{name}" },
  { "action": "click_cell", "title": "Forward+" },
  { "action": "click",      "title": "Create & Edit" }
]
```

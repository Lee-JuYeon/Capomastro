---
ide: godot
appName: Godot
description: Godot 2D Game (GDScript)
---

Godot Project Manager → New → fill name + path → set renderer to "Mobile" for 2D → Create & Edit.
NOTE: Godot Project Manager AX tree is best-guess; validate on first run.

```json
[
  { "action": "click",      "title": "New" },
  { "action": "fill_field", "title": "Project Name",           "value": "{name}" },
  { "action": "fill_field", "title": "Project Path",           "value": "{outputDir}/{name}" },
  { "action": "click_cell", "title": "Mobile" },
  { "action": "click",      "title": "Create & Edit" }
]
```

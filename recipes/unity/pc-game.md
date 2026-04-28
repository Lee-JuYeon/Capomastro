---
ide: unity-hub
appName: Unity Hub
description: Unity PC Game (C# + High Definition Render Pipeline)
---

New project → "3D (HDRP)" or "3D Core" template → set name + location → Create project.
For a simple 3D project without HDRP overhead, use "3D Core" template.

```json
[
  { "action": "click",      "title": "New project" },
  { "action": "click_cell", "title": "3D (URP)" },
  { "action": "fill_field", "title": "Project name",           "value": "{name}" },
  { "action": "fill_field", "title": "Location",               "value": "{outputDir}" },
  { "action": "click",      "title": "Create project" }
]
```

---
ide: unity-hub
appName: Unity Hub
description: Unity Mobile Game (C# + Universal Render Pipeline)
---

Unity Hub new project: New project button → select "Mobile 2D" or "Mobile 3D" core template →
set project name + location → Create project.
After creation, switch Build Target to iOS or Android in Build Settings.
NOTE: Unity Hub AX tree uses custom Java bridge on macOS — AX titles are best-guess, validate on first run.

```json
[
  { "action": "click",      "title": "New project" },
  { "action": "click_cell", "title": "Mobile 2D" },
  { "action": "fill_field", "title": "Project name",           "value": "{name}" },
  { "action": "fill_field", "title": "Location",               "value": "{outputDir}" },
  { "action": "click",      "title": "Create project" }
]
```

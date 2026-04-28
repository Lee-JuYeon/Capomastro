---
ide: unity-hub
appName: Unity Hub
description: Unity AR Game (C# + AR Foundation — ARKit + ARCore)
---

New project → "AR Core" template → set name + location → Create project.
AR Foundation 6.x included. Supports ARKit (iOS) and ARCore (Android) from one codebase.

```json
[
  { "action": "click",      "title": "New project" },
  { "action": "click_cell", "title": "AR Core" },
  { "action": "fill_field", "title": "Project name",           "value": "{name}" },
  { "action": "fill_field", "title": "Location",               "value": "{outputDir}" },
  { "action": "click",      "title": "Create project" }
]
```

---
ide: unity-hub
appName: Unity Hub
description: Unity VR Game (C# + XR Interaction Toolkit — Meta Quest / SteamVR)
---

New project → "VR Core" template → set name + location → Create project.
XR Interaction Toolkit + XR Management included. Switch target to Quest via OpenXR settings post-creation.

```json
[
  { "action": "click",      "title": "New project" },
  { "action": "click_cell", "title": "VR Core" },
  { "action": "fill_field", "title": "Project name",           "value": "{name}" },
  { "action": "fill_field", "title": "Location",               "value": "{outputDir}" },
  { "action": "click",      "title": "Create project" }
]
```

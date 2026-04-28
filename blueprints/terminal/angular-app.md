---
ide: terminal
appName: Terminal
description: Angular App (TypeScript + Angular CLI)
---

@angular/cli new scaffolds a full Angular project with standalone components and SCSS.
--skip-git avoids git init conflicts. --package-manager=npm forces npm for consistency.

```json
[
  { "action": "shell", "cmd": "npx @angular/cli@latest new {name} --style=scss --routing --ssr=false --skip-git --package-manager=npm", "cwd": "{outputDir}", "timeout": 120000 }
]
```

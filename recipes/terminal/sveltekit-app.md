---
ide: terminal
appName: Terminal
description: SvelteKit App (TypeScript)
---

sv create (the official SvelteKit scaffolder) creates a new project with TypeScript and ESLint.
Selects the "SvelteKit minimal" template in non-interactive mode via --template=minimal.

```json
[
  { "action": "shell", "cmd": "npx sv@latest create {name} --template minimal --types ts --no-add-ons", "cwd": "{outputDir}", "timeout": 60000 },
  { "action": "shell", "cmd": "npm install",                                                              "cwd": "{outputDir}/{name}", "timeout": 90000 }
]
```

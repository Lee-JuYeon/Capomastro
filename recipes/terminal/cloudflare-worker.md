---
ide: terminal
appName: Terminal
description: Cloudflare Worker (TypeScript + Wrangler)
---

create cloudflare scaffolds a Cloudflare Worker with TypeScript and Wrangler CLI.
Deploy with: npx wrangler deploy. Runs on Cloudflare's V8-based edge runtime globally.

```json
[
  { "action": "shell", "cmd": "npm create cloudflare@latest {name} -- --type hello-world --lang ts --no-deploy", "cwd": "{outputDir}", "timeout": 60000 }
]
```

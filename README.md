<p align="center">
  <img src="banner.png" alt="Capomastro" width="100%">
</p>

# Capomastro

**Vision-first IDE automation agent for BigBoss OS — the "eyes" that build your projects.**

Capomastro watches the IDE on screen and drives it the way a developer would — looking, clicking, typing, navigating — to scaffold and build **real projects inside real IDEs**. Think of it as a Manus-style agent specialized for IDE GUI automation.

> Part of the [BigBoss OS](https://github.com/Lee-JuYeon/bigbossos) ecosystem.

---

## Concept: the eyes of the agent (Vision-First)

Capomastro is **vision-first**. A vision model looks at the current IDE screen, decides the next action, and acts:

```
  ┌─ Vision model SEES the current IDE screen
  │     (windows, menus, fields, buttons — as pixels, not a fixed script)
  ▼
  Decide next action  →  act (menu / click / fill / set location / shell)
  ▲                                         │
  └─────────────  loop until the project is created  ◄┘
```

- **Reads the UI, doesn't replay a script** — so it adapts across IDEs, versions, and layouts instead of breaking when a button moves.
- **Real IDE artifacts** — the same files/structure a developer would get by clicking through the IDE. Never code-faked output.
- **The "eyes" of BigBoss OS** — when the C-level agents finalize a spec ("the blueprint"), Capomastro is the hand-and-eye that builds the actual project ("the house").

### Why vision-first

Capomastro went through two earlier, more brittle approaches (both preserved below as legacy):

1. **Go generator** — wrote project files directly (pbxproj, gradle…). Problem: the output didn't match what the IDE itself creates.
2. **Deterministic MD blueprints** — fixed click sequences per IDE. Problem: they break the moment an IDE's UI shifts, a version changes, or a new IDE appears, and every target needs its own hand-authored script.

Vision-first generalizes: because Capomastro *sees* the screen, one agent handles many IDEs and adapts to change — no per-screen hardcoding to maintain.

---

## Target IDEs & platforms

Capomastro targets the full scaffolding surface (driven by vision; the legacy blueprints below map the same coverage):

- **Xcode** — iOS / macOS / multiplatform / watchOS / tvOS / visionOS / framework / Swift package / ARKit / SpriteKit / Safari extension / document app
- **Android Studio** — empty & Compose activities / Wear OS / Android TV / ARCore / Flutter app & plugin
- **VSCode** — Flutter / Dart package / extensions (TS, color-theme, snippet)
- **Unity / Unreal / Godot** — mobile/PC/VR/AR games, FPS, 2D/3D
- **Terminal / CLI (52+)** — web (Next/Nuxt/Angular/SvelteKit/Astro…), backend (FastAPI/Gin/Axum/Spring/gRPC…), mobile (Flutter/React Native), AI/ML, blockchain, DevOps/IaC, databases

---

## Legacy / reference (superseded by vision-first)

> The implementations below are kept for reference and during the transition. **Vision-first is the current direction**; these are being phased out.

<details>
<summary><b>Deterministic MD Blueprints</b> — fixed click sequences (legacy)</summary>

A blueprint is an MD file with a JSON list of GUI steps; the BigBoss OS server substitutes variables and drives the IDE via the AX API, falling back to the vision loop on step failure.

```markdown
---
ide: xcode
appName: Xcode
description: iOS SwiftUI App
---

```json
[
  { "action": "menu",       "path": ["File", "New", "Project…"] },
  { "action": "click_tab",  "title": "iOS" },
  { "action": "fill_field", "title": "Product Name", "value": "{name}" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

**Variables:** `{name}` · `{orgId}` · `{bundle}` · `{outputDir}` · `{team}`

**Action types:** `menu` · `click_tab` · `click` · `click_cell` · `filter` · `fill_field` · `set_location` · `wait` · `shell`

Blueprint files live under `blueprints/{xcode,android-studio,vscode,unity,unreal,godot,terminal}/`.
</details>

<details>
<summary><b>Go Generator</b> — direct file generation (oldest, legacy)</summary>

The original Capomastro was a Go binary that generated project files directly (pbxproj, gradle, etc.) without IDE interaction.

```bash
./Capomastro --platform ios --pkg com.example.myapp --name MyApp
```

Superseded first by blueprints, now by vision-first.
</details>

---

## License

MIT

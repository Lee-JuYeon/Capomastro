<p align="center">
  <img src="banner.png" alt="Capomastro" width="100%">
</p>

# Capomastro

**IDE Blueprint agent for BigBoss OS.**
Deterministic project scaffolding via MD blueprint files — no LLM required for the happy path.

> Part of the [BigBoss OS](https://github.com/Lee-JuYeon/bigbossos) ecosystem.

---

## Concept: ViewHolder Agent

Capomastro is a **ViewHolder agent** — a slot-based, sequential executor that reads blueprint MD files and drives IDEs deterministically.

```
Capomastro GitHub repo (this repo)
  └─ blueprints/xcode/*.md
  └─ blueprints/android-studio/*.md
  └─ blueprints/terminal/*.md
  └─ blueprints/vscode/*.md
  └─ blueprints/unity/*.md
  └─ blueprints/unreal/*.md
  └─ blueprints/godot/*.md
        │
        ▼
  BigBoss OS server (ide-blueprint.ts)
  ├─ Fetches blueprints from this repo
  ├─ Substitutes variables: {name} {orgId} {bundle} {outputDir}
  ├─ Drives IDE via IDEDriver (AX API)
  └─ Falls back to LLM vision loop only on step failure
```

**Why blueprints instead of code generation?**
- Xcode, Android Studio, Unity — GUI IDEs require the same click sequence every time
- MD blueprint = hardcoded, reproducible, zero-LLM
- Same blueprint works across projects; only variable values change

---

## Blueprint Format

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
  { "action": "click",      "title": "Next" },
  { "action": "fill_field", "title": "Product Name",            "value": "{name}" },
  { "action": "fill_field", "title": "Organization Identifier", "value": "{orgId}" },
  { "action": "click",      "title": "Next" },
  { "action": "set_location" },
  { "action": "click",      "title": "Create" }
]
```

### Variables

| Variable | Description |
|----------|-------------|
| `{name}` | Project name |
| `{orgId}` | Organization identifier (e.g. `com.company`) |
| `{bundle}` | Full bundle ID (e.g. `com.company.app`) |
| `{outputDir}` | Output directory (set at CLI startup via `PROJECT_DIR`) |
| `{team}` | Apple Team ID (optional) |

### Action Types

| Action | Description |
|--------|-------------|
| `menu` | Click menu bar + navigate submenus |
| `click_tab` | Click tab / radio button |
| `click` | Click button |
| `click_cell` | Select cell in grid or list |
| `filter` | Type in search field (auto-reveals toggle if hidden) |
| `fill_field` | Triple-click + type in text field |
| `set_location` | Handle "Where:" popup → set to `{outputDir}` |
| `wait` | Pause for `ms` milliseconds |
| `shell` | Run CLI command (for Terminal-based platforms) |

---

## Supported Platforms

### Xcode (12 blueprints)
| Blueprint | Description |
|--------|-------------|
| `xcode/ios-app.md` | iOS SwiftUI App |
| `xcode/macos-app.md` | macOS SwiftUI App |
| `xcode/multiplatform.md` | Multiplatform SwiftUI App |
| `xcode/watchos-app.md` | watchOS App |
| `xcode/tvos-app.md` | tvOS App |
| `xcode/visionos-app.md` | visionOS App |
| `xcode/framework.md` | iOS Framework |
| `xcode/swift-package.md` | Swift Package |
| `xcode/arkit-app.md` | ARKit App |
| `xcode/game-spritekit.md` | SpriteKit Game |
| `xcode/safari-extension.md` | Safari Web Extension |
| `xcode/document-app.md` | Document-Based App |

### Android Studio (7 blueprints)
| Blueprint | Description |
|--------|-------------|
| `android-studio/empty-activity.md` | Empty Activity (Views) |
| `android-studio/compose-activity.md` | Jetpack Compose Activity |
| `android-studio/wearos-app.md` | Wear OS App |
| `android-studio/tv-app.md` | Android TV App |
| `android-studio/arcore-app.md` | ARCore App |
| `android-studio/flutter-app.md` | Flutter App (via plugin) |
| `android-studio/flutter-plugin.md` | Flutter Plugin |

### VSCode (6 blueprints)
| Blueprint | Description |
|--------|-------------|
| `vscode/flutter-app.md` | Flutter App (via extension) |
| `vscode/dart-package.md` | Dart Package |
| `vscode/flutter-plugin.md` | Flutter Plugin |
| `vscode/extension-ts.md` | VSCode TypeScript Extension (yo code) |
| `vscode/extension-color-theme.md` | VSCode Color Theme Extension (yo code) |
| `vscode/extension-snippet.md` | VSCode Snippet Pack Extension (yo code) |

> VSCode blueprints = VSCode-specific wizards only. Web/backend scaffolding → Terminal blueprints.

### Unity (4 blueprints)
`unity/mobile-game.md` · `unity/pc-game.md` · `unity/vr-game.md` · `unity/ar-game.md`

### Unreal (2 blueprints)
`unreal/blank-game.md` · `unreal/fps-game.md`

### Godot (2 blueprints)
`godot/2d-game.md` · `godot/3d-game.md`

### Terminal / CLI (52+ blueprints)

**Mobile**: `flutter-app` · `react-native-app`

**Web**: `nextjs-app` · `nuxt-app` · `angular-app` · `sveltekit-app` · `astro-site` · `chrome-extension` · `firefox-extension`

**Desktop**: `electron-app` · `tauri-app` · `dotnet-winui`

**Backend**: `node-fastify-api` · `python-fastapi` · `go-gin-api` · `rust-axum-api` · `java-springboot-api` · `graphql-apollo` · `grpc-go` · `cloudflare-worker` · `node-websocket` · `kafka-consumer`

**System / CLI tools**: `rust-cli-tool` · `go-cli-tool` · `python-cli-tool` · `embedded-arduino` · `ros2-robot`

**AI / ML**: `pytorch-training` · `fastapi-inference` · `langchain-rag` · `mlflow-mlops` · `airflow-pipeline` · `opencv-vision` · `whisper-audio` · `autogen-agent`

**Blockchain**: `hardhat-evm` · `anchor-solana`

**AR / VR / Games**: `webxr-threejs` · `phaser-webgame`

**Data**: `playwright-scraper` · `prometheus-grafana`

**Database**: `supabase-local` · `mysql-docker` · `mongodb-docker` · `redis-docker` · `pgvector-app`

**DevOps / IaC**: `terraform-aws` · `pulumi-typescript` · `kubernetes-helm` · `github-actions-cicd` · `elk-logging`

---

## Legacy: Go Generator

The original Capomastro was a Go binary that generated project files directly (pbxproj, gradle, etc.) without IDE interaction. It is preserved in this repo for reference.

```bash
# Legacy CLI (Go binary)
./Capomastro --platform ios --pkg com.example.myapp --name MyApp
./Capomastro --platform android --pkg com.example.myapp --name MyApp
```

The new blueprint-based approach supersedes this for most use cases.

---

## License

MIT

# ProjectBuilder

CLI tool that generates buildable Xcode/Android Studio projects. Detects your IDE version and creates compatible projects with zero external dependencies.

## Install

```bash
go install github.com/cavss/ProjectBuilder@latest
```

Or download the binary from [Releases](https://github.com/Lee-JuYeon/ProjectBuilder/releases).

## Usage

```bash
ProjectBuilder --platform <PLATFORM> --pkg <PACKAGE> --name <NAME> [options]
```

### Required flags

| Flag | Description | Example |
|------|-------------|---------|
| `--platform` | Target platform | `ios`, `macos`, `android`, `flutter` |
| `--pkg` | Bundle/package identifier | `com.example.myapp` |
| `--name` | Project name | `MyApp` |

### Optional flags

| Flag | Default | Description |
|------|---------|-------------|
| `--out` | `.` | Output directory |
| `--framework` | `swiftui` (iOS) / `compose` (Android) | UI framework: `swiftui`, `uikit`, `compose`, `xml` |
| `--min-version` | Auto-detected | Minimum OS version (e.g. `15.0`, `24`) |
| `--coredata` | `false` | Include Core Data model (iOS/macOS only) |

## Examples

### iOS (SwiftUI)

```bash
ProjectBuilder --platform ios --pkg com.example.myapp --name MyApp
```

### iOS (UIKit + CoreData)

```bash
ProjectBuilder --platform ios --pkg com.example.myapp --name MyApp --framework uikit --coredata
```

### iOS (minimum version override)

```bash
ProjectBuilder --platform ios --pkg com.example.myapp --name MyApp --min-version 15.0
```

### macOS

```bash
ProjectBuilder --platform macos --pkg com.example.myapp --name MyApp
```

### Android (Jetpack Compose)

```bash
ProjectBuilder --platform android --pkg com.example.myapp --name MyApp
```

### Android (XML layout)

```bash
ProjectBuilder --platform android --pkg com.example.myapp --name MyApp --framework xml
```

### Flutter

```bash
ProjectBuilder --platform flutter --pkg com.example --name MyApp
```

## How it works

1. Detects the host machine's IDE version
   - **Xcode**: `xcodebuild -version`, `swift --version`
   - **Android Studio**: reads `build.txt` from installation
2. Maps IDE version to compatible build tool versions (AGP, Gradle, Kotlin, Swift)
3. Generates a complete, buildable project with proper structure

### Generated iOS project

```
MyApp/
├── MyApp.xcodeproj/project.pbxproj   ← generated directly (no XcodeGen)
├── MyApp/
│   ├── MyAppApp.swift
│   ├── ContentView.swift
│   └── Assets.xcassets/
├── MyAppTests/
│   └── MyAppTests.swift
└── MyAppUITests/
    ├── MyAppUITests.swift
    └── MyAppUITestsLaunchTests.swift
```

### Generated Android project

```
MyApp/
├── build.gradle.kts
├── settings.gradle.kts
├── gradle.properties
├── gradle/wrapper/gradle-wrapper.properties
└── app/
    ├── build.gradle.kts
    └── src/main/
        ├── AndroidManifest.xml
        ├── java/com/example/myapp/MainActivity.kt
        └── res/values/strings.xml
```

## Key features

- **Zero external dependencies** — no XcodeGen, no Tuist, no Cocoapods
- **IDE version detection** — generates projects compatible with your installed IDE
- **Single binary** — download and run, no runtime needed
- **CLI-first** — non-interactive, scriptable, CI/CD friendly

## Supported Android Studio versions

| Version | Codename | AGP | Gradle | Kotlin |
|---------|----------|-----|--------|--------|
| 2024.3 | Meerkat | 8.9.0 | 8.11.1 | 2.1.0 |
| 2024.2 | Ladybug | 8.8.0 | 8.10.2 | 2.0.21 |
| 2024.1 | Koala | 8.5.0 | 8.7 | 2.0.0 |
| 2023.3 | Jellyfish | 8.4.0 | 8.6 | 1.9.23 |
| 2023.2 | Iguana | 8.3.0 | 8.4 | 1.9.22 |
| 2023.1 | Hedgehog | 8.2.0 | 8.2 | 1.9.20 |
| 2022.3 | Giraffe | 8.1.0 | 8.0 | 1.9.0 |
| 2022.2 | Flamingo | 8.0.0 | 8.0 | 1.8.10 |
| 2022.1 | Electric Eel | 7.4.0 | 7.5 | 1.8.0 |
| 2021.3 | Dolphin | 7.3.0 | 7.4 | 1.7.20 |
| 2021.2 | Chipmunk | 7.2.0 | 7.3.3 | 1.7.10 |
| 2021.1 | Bumblebee | 7.1.0 | 7.2 | 1.6.21 |
| 2020.3 | Arctic Fox | 7.0.0 | 7.0.2 | 1.5.31 |

## License

MIT

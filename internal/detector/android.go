package detector

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type AndroidStudioInfo struct {
	VersionName   string // e.g. "Giraffe", "Bumblebee"
	VersionCode   string // e.g. "2023.3.1"
	AGPVersion    string // Android Gradle Plugin version
	GradleVersion string // Gradle wrapper version
	KotlinVersion string // Kotlin version
}

// studioVersionMap maps Android Studio codename prefix to compatible versions.
// Key: first part of version code (year.major)
var studioVersionMap = map[string]struct {
	name    string
	agp     string
	gradle  string
	kotlin  string
}{
	"2024.3": {name: "Meerkat", agp: "8.10.1", gradle: "8.11.1", kotlin: "2.0.21"},
	"2024.2": {name: "Ladybug", agp: "8.8.2", gradle: "8.10.2", kotlin: "2.0.21"},
	"2024.1": {name: "Koala", agp: "8.5.2", gradle: "8.7", kotlin: "2.0.0"},
	"2023.3": {name: "Jellyfish", agp: "8.4.0", gradle: "8.6", kotlin: "1.9.23"},
	"2023.2": {name: "Iguana", agp: "8.3.0", gradle: "8.4", kotlin: "1.9.22"},
	"2023.1": {name: "Hedgehog", agp: "8.2.0", gradle: "8.2", kotlin: "1.9.20"},
	"2022.3": {name: "Giraffe", agp: "8.1.0", gradle: "8.0", kotlin: "1.9.0"},
	"2022.2": {name: "Flamingo", agp: "8.0.0", gradle: "8.0", kotlin: "1.8.10"},
	"2022.1": {name: "Electric Eel", agp: "7.4.0", gradle: "7.5", kotlin: "1.8.0"},
	"2021.3": {name: "Dolphin", agp: "7.3.0", gradle: "7.4", kotlin: "1.7.20"},
	"2021.2": {name: "Chipmunk", agp: "7.2.0", gradle: "7.3.3", kotlin: "1.7.10"},
	"2021.1": {name: "Bumblebee", agp: "7.1.0", gradle: "7.2", kotlin: "1.6.21"},
	"2020.3": {name: "Arctic Fox", agp: "7.0.0", gradle: "7.0.2", kotlin: "1.5.31"},
}

func DetectAndroidStudio() (*AndroidStudioInfo, error) {
	buildFile := findBuildFile()
	if buildFile == "" {
		return nil, fmt.Errorf("Android Studio not found — is it installed?")
	}

	data, err := os.ReadFile(buildFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read Android Studio build file: %w", err)
	}

	// build.txt contains something like "AI-2022.1.1.21" or "AI-2024.2.1.11"
	raw := strings.TrimSpace(string(data))
	versionCode := strings.TrimPrefix(raw, "AI-")

	// build.txt format: "AI-243.26053.27.2432.13536105"
	// 243 → 20(24).(3) → 2024.3
	// Or older format: "AI-2022.1.1.21" → 2022.1
	parts := strings.Split(versionCode, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("unexpected Android Studio version format: %s", raw)
	}

	key := ""
	if len(parts[0]) == 3 {
		// New format: 243 → 2024.3, 232 → 2023.2, 211 → 2021.1
		code := parts[0]
		year := "20" + code[:2]
		major := string(code[2])
		key = year + "." + major
	} else {
		// Old format: 2022.1.1.21 → 2022.1
		key = parts[0] + "." + parts[1]
	}

	info := &AndroidStudioInfo{
		VersionCode: versionCode,
	}

	if mapping, ok := studioVersionMap[key]; ok {
		info.VersionName = mapping.name
		info.AGPVersion = mapping.agp
		info.GradleVersion = mapping.gradle
		info.KotlinVersion = mapping.kotlin
	} else {
		// Unknown version — use latest safe defaults
		info.VersionName = "Unknown"
		info.AGPVersion = "8.5.0"
		info.GradleVersion = "8.7"
		info.KotlinVersion = "2.0.0"
	}

	return info, nil
}

func findBuildFile() string {
	var candidates []string

	switch runtime.GOOS {
	case "darwin":
		home, _ := os.UserHomeDir()
		// Standard installation
		candidates = append(candidates,
			"/Applications/Android Studio.app/Contents/Resources/build.txt",
			filepath.Join(home, "Applications/Android Studio.app/Contents/Resources/build.txt"),
		)
		// Toolbox installations
		toolboxBase := filepath.Join(home, "Library/Application Support/JetBrains/Toolbox/apps")
		if entries, err := os.ReadDir(toolboxBase); err == nil {
			for _, e := range entries {
				if strings.Contains(strings.ToLower(e.Name()), "android") {
					subPath := filepath.Join(toolboxBase, e.Name())
					if channels, err := os.ReadDir(subPath); err == nil {
						for _, ch := range channels {
							candidates = append(candidates,
								filepath.Join(subPath, ch.Name(), "Android Studio.app/Contents/Resources/build.txt"),
							)
						}
					}
				}
			}
		}
	case "linux":
		home, _ := os.UserHomeDir()
		candidates = append(candidates,
			filepath.Join(home, ".local/share/JetBrains/Toolbox/apps/AndroidStudio/ch-0/build.txt"),
			"/opt/android-studio/build.txt",
			"/usr/local/android-studio/build.txt",
		)
	case "windows":
		programFiles := os.Getenv("ProgramFiles")
		candidates = append(candidates,
			filepath.Join(programFiles, "Android/Android Studio/build.txt"),
		)
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

package android

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cavss/ProjectBuilder/internal/detector"
)

type Config struct {
	Name        string
	PackageName string
	OutputDir   string
	Framework   string // "compose" or "xml"
	MinVersion  string // minSdk override
	TargetSdk   string // targetSdk/compileSdk override
	JavaVersion string // "11", "17", "21"
	StudioInfo  *detector.AndroidStudioInfo
}

func Generate(cfg Config) error {
	projectRoot := filepath.Join(cfg.OutputDir, cfg.Name)
	appDir := filepath.Join(projectRoot, "app")
	srcMain := filepath.Join(appDir, "src", "main")
	javaDir := filepath.Join(srcMain, "java", pkgToPath(cfg.PackageName))
	resValues := filepath.Join(srcMain, "res", "values")
	gradleWrapper := filepath.Join(projectRoot, "gradle", "wrapper")

	// Determine SDK versions
	minSdk := cfg.MinVersion
	if minSdk == "" {
		minSdk = "24"
	}
	targetSdk := cfg.TargetSdk
	if targetSdk == "" {
		targetSdk = "35"
	}
	compileSdk := targetSdk

	javaVer := cfg.JavaVersion
	if javaVer == "" {
		javaVer = "17"
	}

	agp := cfg.StudioInfo.AGPVersion
	gradle := cfg.StudioInfo.GradleVersion
	kotlin := cfg.StudioInfo.KotlinVersion

	// Create directories
	dirs := []string{javaDir, resValues, gradleWrapper}
	if cfg.Framework == "xml" {
		dirs = append(dirs, filepath.Join(srcMain, "res", "layout"))
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", d, err)
		}
	}

	// Root files
	files := map[string]string{
		filepath.Join(projectRoot, "build.gradle.kts"):            rootBuildGradle(agp, kotlin),
		filepath.Join(projectRoot, "settings.gradle.kts"):         settingsGradle(cfg.Name),
		filepath.Join(projectRoot, "gradle.properties"):           gradleProperties(),
		filepath.Join(gradleWrapper, "gradle-wrapper.properties"): gradleWrapperProperties(gradle),
		filepath.Join(srcMain, "AndroidManifest.xml"):             androidManifest(cfg.PackageName, cfg.Framework),
		filepath.Join(resValues, "strings.xml"):                   stringsXML(cfg.Name),
		filepath.Join(appDir, "proguard-rules.pro"):               proguardRules(),
	}

	// App build.gradle.kts
	if cfg.Framework == "compose" {
		files[filepath.Join(appDir, "build.gradle.kts")] = appBuildGradleCompose(cfg.PackageName, agp, kotlin, minSdk, targetSdk, compileSdk, javaVer)
		files[filepath.Join(javaDir, "MainActivity.kt")] = mainActivityCompose(cfg.PackageName)
	} else {
		files[filepath.Join(appDir, "build.gradle.kts")] = appBuildGradleXML(cfg.PackageName, agp, kotlin, minSdk, targetSdk, compileSdk, javaVer)
		files[filepath.Join(javaDir, "MainActivity.kt")] = mainActivityXML(cfg.PackageName)
		files[filepath.Join(srcMain, "res", "layout", "activity_main.xml")] = activityMainXML()
	}

	// Write all files
	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", path, err)
		}
	}

	fmt.Printf("Android project created: %s\n", projectRoot)
	fmt.Printf("  minSdk: %s, targetSdk: %s, Java: %s\n", minSdk, targetSdk, javaVer)
	fmt.Printf("  AGP: %s, Gradle: %s, Kotlin: %s\n", agp, gradle, kotlin)
	fmt.Printf("  Framework: %s\n", cfg.Framework)
	return nil
}

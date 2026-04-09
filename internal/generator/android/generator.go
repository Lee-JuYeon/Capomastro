package android

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cavss/Capomastro/internal/detector"
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

	// Test directories
	testDir := filepath.Join(appDir, "src", "test", "java", pkgToPath(cfg.PackageName))
	androidTestDir := filepath.Join(appDir, "src", "androidTest", "java", pkgToPath(cfg.PackageName))

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

	// Resource directories
	resDrawable := filepath.Join(srcMain, "res", "drawable")
	resDrawableV24 := filepath.Join(srcMain, "res", "drawable-v24")
	resMipmapAnydpi := filepath.Join(srcMain, "res", "mipmap-anydpi-v26")
	resXml := filepath.Join(srcMain, "res", "xml")

	// ui.theme directory (Compose only)
	themeDir := filepath.Join(srcMain, "java", pkgToPath(cfg.PackageName), "ui", "theme")

	// Create directories
	dirs := []string{javaDir, resValues, resDrawable, resDrawableV24, resMipmapAnydpi, resXml, gradleWrapper, testDir, androidTestDir}
	if cfg.Framework == "compose" {
		dirs = append(dirs, themeDir)
	}
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
		filepath.Join(projectRoot, ".gitignore"):                   gitignore(),
		filepath.Join(projectRoot, "local.properties"):            localProperties(),
		filepath.Join(projectRoot, "gradle", "libs.versions.toml"): libsVersionsToml(agp, kotlin),
		filepath.Join(gradleWrapper, "gradle-wrapper.properties"): gradleWrapperProperties(gradle),
		filepath.Join(srcMain, "AndroidManifest.xml"):             androidManifest(cfg.PackageName, cfg.Framework),
		filepath.Join(resValues, "strings.xml"):                   stringsXML(cfg.Name),
		filepath.Join(appDir, "proguard-rules.pro"):               proguardRules(),
	}

	// gradlew scripts
	files[filepath.Join(projectRoot, "gradlew")] = gradlewScript()
	files[filepath.Join(projectRoot, "gradlew.bat")] = gradlewBat()

	// Test files
	files[filepath.Join(testDir, cfg.Name+"Test.kt")] = unitTestFile(cfg.PackageName, cfg.Name)
	files[filepath.Join(androidTestDir, cfg.Name+"InstrumentedTest.kt")] = instrumentedTestFile(cfg.PackageName, cfg.Name)

	// Resource files (shared)
	files[filepath.Join(resDrawable, "ic_launcher_background.xml")] = icLauncherBackground()
	files[filepath.Join(resDrawableV24, "ic_launcher_foreground.xml")] = icLauncherForeground()
	files[filepath.Join(resMipmapAnydpi, "ic_launcher.xml")] = icLauncherXML()
	files[filepath.Join(resMipmapAnydpi, "ic_launcher_round.xml")] = icLauncherXML()
	files[filepath.Join(resValues, "colors.xml")] = colorsXML()
	files[filepath.Join(resValues, "themes.xml")] = themesXML(cfg.Name)
	files[filepath.Join(resXml, "backup_rules.xml")] = backupRulesXML()
	files[filepath.Join(resXml, "data_extraction_rules.xml")] = dataExtractionRulesXML()

	// App build.gradle.kts
	if cfg.Framework == "compose" {
		files[filepath.Join(appDir, "build.gradle.kts")] = appBuildGradleCompose(cfg.PackageName, agp, kotlin, minSdk, targetSdk, compileSdk, javaVer)
		files[filepath.Join(javaDir, "MainActivity.kt")] = mainActivityCompose(cfg.PackageName)
		// ui.theme files
		files[filepath.Join(themeDir, "Color.kt")] = colorKt(cfg.PackageName)
		files[filepath.Join(themeDir, "Theme.kt")] = themeKt(cfg.PackageName, cfg.Name)
		files[filepath.Join(themeDir, "Type.kt")] = typeKt(cfg.PackageName)
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

	// Make gradlew executable
	os.Chmod(filepath.Join(projectRoot, "gradlew"), 0755)

	// Detect sdk.dir and append to local.properties
	sdkDir := findAndroidSDK()
	if sdkDir != "" {
		lp := filepath.Join(projectRoot, "local.properties")
		content, _ := os.ReadFile(lp)
		os.WriteFile(lp, []byte(string(content)+fmt.Sprintf("sdk.dir=%s\n", sdkDir)), 0644)
	}

	// Copy gradle-wrapper.jar from system if available
	copyGradleWrapperJar(filepath.Join(gradleWrapper, "gradle-wrapper.jar"))

	fmt.Printf("Android project created: %s\n", projectRoot)
	fmt.Printf("  minSdk: %s, targetSdk: %s, Java: %s\n", minSdk, targetSdk, javaVer)
	fmt.Printf("  AGP: %s, Gradle: %s, Kotlin: %s\n", agp, gradle, kotlin)
	fmt.Printf("  Framework: %s\n", cfg.Framework)
	return nil
}

func findAndroidSDK() string {
	// Check ANDROID_HOME / ANDROID_SDK_ROOT
	for _, env := range []string{"ANDROID_HOME", "ANDROID_SDK_ROOT"} {
		if v := os.Getenv(env); v != "" {
			return v
		}
	}
	// Default macOS location
	home, _ := os.UserHomeDir()
	defaultPath := filepath.Join(home, "Library", "Android", "sdk")
	if _, err := os.Stat(defaultPath); err == nil {
		return defaultPath
	}
	return ""
}

// G22 근본 해결: 환경 의존성 완전 제거 — 바이너리에 번들링
//
//go:embed assets/gradle-wrapper.jar
var gradleWrapperJarData []byte

func copyGradleWrapperJar(dest string) {
	if err := os.WriteFile(dest, gradleWrapperJarData, 0644); err != nil {
		fmt.Printf("  [warn] gradle-wrapper.jar 쓰기 실패: %v\n", err)
		return
	}
	fmt.Printf("  gradle-wrapper.jar 번들에서 복사 완료\n")
}

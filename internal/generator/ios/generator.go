package ios

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cavss/ProjectBuilder/internal/detector"
)

type Config struct {
	Name       string
	BundleID   string
	OutputDir  string
	Platform   string // "ios" or "macos"
	Framework  string // "swiftui" or "uikit"
	MinVersion string
	CoreData   bool
	Team       string // Apple Development Team ID
	XcodeInfo  *detector.XcodeInfo
}

func Generate(cfg Config) error {
	projectRoot := filepath.Join(cfg.OutputDir, cfg.Name)
	xcodeproj := filepath.Join(projectRoot, cfg.Name+".xcodeproj")
	sourcesDir := filepath.Join(projectRoot, cfg.Name)
	assetsDir := filepath.Join(sourcesDir, "Assets.xcassets")
	appIconDir := filepath.Join(assetsDir, "AppIcon.appiconset")
	accentDir := filepath.Join(assetsDir, "AccentColor.colorset")

	testsDir := filepath.Join(projectRoot, cfg.Name+"Tests")
	uiTestsDir := filepath.Join(projectRoot, cfg.Name+"UITests")

	// Create directories
	dirs := []string{xcodeproj, sourcesDir, appIconDir, accentDir, testsDir, uiTestsDir}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", d, err)
		}
	}

	// Generate project.pbxproj
	pbxproj := generatePbxproj(cfg)
	if err := os.WriteFile(filepath.Join(xcodeproj, "project.pbxproj"), []byte(pbxproj), 0644); err != nil {
		return fmt.Errorf("failed to write project.pbxproj: %w", err)
	}

	// Generate source files
	// 소스 파일은 하위 디렉토리에 같은 이름이 있으면 생성하지 않음 (에이전트 코드 보존)
	writeIfNotExists := func(path string, content string) error {
		if _, err := os.Stat(path); err == nil {
			return nil // 루트에 이미 존재 → 스킵
		}
		// 하위 디렉토리에 같은 이름 파일이 있는지 확인
		base := filepath.Base(path)
		dir := filepath.Dir(path)
		found := false
		filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() { return nil }
			if filepath.Base(p) == base && p != path { found = true }
			return nil
		})
		if found {
			return nil // 하위에 동명 파일 존재 → 스킵
		}
		return os.WriteFile(path, []byte(content), 0644)
	}

	if cfg.Framework == "swiftui" {
		files := map[string]string{
			cfg.Name + "App.swift": swiftuiAppFile(cfg.Name),
			"ContentView.swift":    swiftuiContentView(),
		}
		for name, content := range files {
			if err := writeIfNotExists(filepath.Join(sourcesDir, name), content); err != nil {
				return fmt.Errorf("failed to write %s: %w", name, err)
			}
		}
	} else {
		files := map[string]string{
			"AppDelegate.swift":    uikitAppDelegate(),
			"SceneDelegate.swift":  uikitSceneDelegate(),
			"ViewController.swift": uikitViewController(),
		}
		for name, content := range files {
			if err := writeIfNotExists(filepath.Join(sourcesDir, name), content); err != nil {
				return fmt.Errorf("failed to write %s: %w", name, err)
			}
		}
	}

	// Generate asset catalogs
	if err := os.WriteFile(filepath.Join(assetsDir, "Contents.json"), []byte(assetContentsJSON()), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(appIconDir, "Contents.json"), []byte(appIconContentsJSON()), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(accentDir, "Contents.json"), []byte(accentColorContentsJSON()), 0644); err != nil {
		return err
	}

	// Generate test files
	if err := os.WriteFile(filepath.Join(testsDir, cfg.Name+"Tests.swift"), []byte(unitTestFile(cfg.Name)), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(uiTestsDir, cfg.Name+"UITests.swift"), []byte(uiTestFile(cfg.Name)), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(uiTestsDir, cfg.Name+"UITestsLaunchTests.swift"), []byte(uiTestLaunchFile(cfg.Name)), 0644); err != nil {
		return err
	}

	// Generate CoreData model (optional) — skip if xcdatamodeld already exists in any subdirectory
	if cfg.CoreData {
		modelDir := filepath.Join(sourcesDir, cfg.Name+".xcdatamodeld", cfg.Name+".xcdatamodel")
		existing := false
		filepath.Walk(sourcesDir, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() && filepath.Ext(p) == ".xcdatamodeld" {
				existing = true
			}
			return nil
		})
		if !existing {
			if err := os.MkdirAll(modelDir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(filepath.Join(modelDir, "contents"), []byte(coreDataContents()), 0644); err != nil {
				return err
			}
		}
	}

	fmt.Printf("iOS project created: %s\n", projectRoot)
	fmt.Printf("Open with: open %s\n", filepath.Join(xcodeproj, ".."))
	return nil
}

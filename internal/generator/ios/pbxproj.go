package ios

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func genUUID(seed string) string {
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%X", h[:12])
}

func generatePbxproj(cfg Config) string {
	b := newPbxBuilder(cfg)
	return b.build()
}

type pbxBuilder struct {
	cfg  Config
	name string

	// App target IDs
	projectID                string
	mainGroupID              string
	sourcesGroupID           string
	productsGroupID          string
	appTargetID              string
	appProductID             string
	configListProjectID      string
	configListAppID          string
	debugConfigProjectID     string
	releaseConfigProjectID   string
	debugConfigAppID         string
	releaseConfigAppID       string
	sourcesBuildPhaseAppID   string
	resourcesBuildPhaseAppID string
	frameworksBuildPhaseAppID string

	// Tests target IDs
	testsGroupID              string
	testsTargetID             string
	testsProductID            string
	configListTestsID         string
	debugConfigTestsID        string
	releaseConfigTestsID      string
	sourcesBuildPhaseTestsID  string
	resourcesBuildPhaseTestsID string
	frameworksBuildPhaseTestsID string
	testsDependencyID         string
	testsProxyID              string

	// UITests target IDs
	uiTestsGroupID              string
	uiTestsTargetID             string
	uiTestsProductID            string
	configListUITestsID         string
	debugConfigUITestsID        string
	releaseConfigUITestsID      string
	sourcesBuildPhaseUITestsID  string
	resourcesBuildPhaseUITestsID string
	frameworksBuildPhaseUITestsID string
	uiTestsDependencyID         string
	uiTestsProxyID              string

	// File IDs
	assetsRefID    string
	assetsBuildID  string
	coreDataRefID  string
	coreDataBuildID string

	// Platform settings
	sdkroot            string
	supportedPlatforms string
	deploymentSetting  string
	deviceFamily       string
	swiftVersion       string
}

func newPbxBuilder(cfg Config) *pbxBuilder {
	n := cfg.Name
	b := &pbxBuilder{cfg: cfg, name: n}

	// App
	b.projectID = genUUID(n + ".project")
	b.mainGroupID = genUUID(n + ".mainGroup")
	b.sourcesGroupID = genUUID(n + ".sourcesGroup")
	b.productsGroupID = genUUID(n + ".productsGroup")
	b.appTargetID = genUUID(n + ".appTarget")
	b.appProductID = genUUID(n + ".appProduct")
	b.configListProjectID = genUUID(n + ".configList.project")
	b.configListAppID = genUUID(n + ".configList.target")
	b.debugConfigProjectID = genUUID(n + ".config.project.debug")
	b.releaseConfigProjectID = genUUID(n + ".config.project.release")
	b.debugConfigAppID = genUUID(n + ".config.target.debug")
	b.releaseConfigAppID = genUUID(n + ".config.target.release")
	b.sourcesBuildPhaseAppID = genUUID(n + ".sourcesBuildPhase")
	b.resourcesBuildPhaseAppID = genUUID(n + ".resourcesBuildPhase")
	b.frameworksBuildPhaseAppID = genUUID(n + ".frameworksBuildPhase")

	// Tests
	b.testsGroupID = genUUID(n + ".testsGroup")
	b.testsTargetID = genUUID(n + ".testsTarget")
	b.testsProductID = genUUID(n + ".testsProduct")
	b.configListTestsID = genUUID(n + ".configList.tests")
	b.debugConfigTestsID = genUUID(n + ".config.tests.debug")
	b.releaseConfigTestsID = genUUID(n + ".config.tests.release")
	b.sourcesBuildPhaseTestsID = genUUID(n + ".sourcesBuildPhase.tests")
	b.resourcesBuildPhaseTestsID = genUUID(n + ".resourcesBuildPhase.tests")
	b.frameworksBuildPhaseTestsID = genUUID(n + ".frameworksBuildPhase.tests")
	b.testsDependencyID = genUUID(n + ".testsDependency")
	b.testsProxyID = genUUID(n + ".testsProxy")

	// UITests
	b.uiTestsGroupID = genUUID(n + ".uiTestsGroup")
	b.uiTestsTargetID = genUUID(n + ".uiTestsTarget")
	b.uiTestsProductID = genUUID(n + ".uiTestsProduct")
	b.configListUITestsID = genUUID(n + ".configList.uiTests")
	b.debugConfigUITestsID = genUUID(n + ".config.uiTests.debug")
	b.releaseConfigUITestsID = genUUID(n + ".config.uiTests.release")
	b.sourcesBuildPhaseUITestsID = genUUID(n + ".sourcesBuildPhase.uiTests")
	b.resourcesBuildPhaseUITestsID = genUUID(n + ".resourcesBuildPhase.uiTests")
	b.frameworksBuildPhaseUITestsID = genUUID(n + ".frameworksBuildPhase.uiTests")
	b.uiTestsDependencyID = genUUID(n + ".uiTestsDependency")
	b.uiTestsProxyID = genUUID(n + ".uiTestsProxy")

	// Files
	b.assetsRefID = genUUID(n + ".ref.assets")
	b.assetsBuildID = genUUID(n + ".build.assets")
	b.coreDataRefID = genUUID(n + ".ref.coredata")
	b.coreDataBuildID = genUUID(n + ".build.coredata")

	// Platform
	if cfg.Platform == "ios" {
		b.sdkroot = "iphoneos"
		b.supportedPlatforms = `"iphoneos iphonesimulator"`
		b.deploymentSetting = fmt.Sprintf("IPHONEOS_DEPLOYMENT_TARGET = %s;", cfg.MinVersion)
		b.deviceFamily = `"1,2"`
	} else {
		b.sdkroot = "macosx"
		b.supportedPlatforms = `"macosx"`
		b.deploymentSetting = fmt.Sprintf("MACOSX_DEPLOYMENT_TARGET = %s;", cfg.MinVersion)
		b.deviceFamily = ""
	}

	b.swiftVersion = "5.0"
	if cfg.XcodeInfo != nil && cfg.XcodeInfo.SwiftVersion != "" {
		parts := strings.Split(cfg.XcodeInfo.SwiftVersion, ".")
		if len(parts) >= 2 {
			b.swiftVersion = parts[0] + "." + parts[1]
		} else {
			b.swiftVersion = parts[0] + ".0"
		}
	}

	return b
}

type fileEntry struct {
	refID, buildID, name, path string
}

func (b *pbxBuilder) appSourceFiles() []fileEntry {
	n := b.name
	// 소스 디렉토리가 존재하면 실제 .swift 파일을 스캔
	sourcesDir := filepath.Join(b.cfg.OutputDir, n, n)
	if info, err := os.Stat(sourcesDir); err == nil && info.IsDir() {
		return b.scanSwiftFiles(sourcesDir, "")
	}
	// 디렉토리가 없으면 (신규 생성) 기본 파일만
	if b.cfg.Framework == "swiftui" {
		return []fileEntry{
			{genUUID(n + ".ref.app"), genUUID(n + ".build.app"), n + "App.swift", n + "App.swift"},
			{genUUID(n + ".ref.contentview"), genUUID(n + ".build.contentview"), "ContentView.swift", "ContentView.swift"},
		}
	}
	return []fileEntry{
		{genUUID(n + ".ref.appdelegate"), genUUID(n + ".build.appdelegate"), "AppDelegate.swift", "AppDelegate.swift"},
		{genUUID(n + ".ref.scenedelegate"), genUUID(n + ".build.scenedelegate"), "SceneDelegate.swift", "SceneDelegate.swift"},
		{genUUID(n + ".ref.viewcontroller"), genUUID(n + ".build.viewcontroller"), "ViewController.swift", "ViewController.swift"},
	}
}

// scanSwiftFiles — 디렉토리 재귀 스캔하여 모든 .swift 파일을 fileEntry로 반환
func (b *pbxBuilder) scanSwiftFiles(baseDir string, relDir string) []fileEntry {
	var entries []fileEntry
	dir := baseDir
	if relDir != "" {
		dir = filepath.Join(baseDir, relDir)
	}
	items, err := os.ReadDir(dir)
	if err != nil {
		return entries
	}
	for _, item := range items {
		if item.IsDir() {
			// 하위 디렉토리 재귀 (build, .build, xcassets, xcdatamodeld 제외)
			name := item.Name()
			if name == "build" || name == ".build" || strings.HasSuffix(name, ".xcassets") || strings.HasSuffix(name, ".xcdatamodeld") {
				continue
			}
			sub := name
			if relDir != "" {
				sub = relDir + "/" + name
			}
			entries = append(entries, b.scanSwiftFiles(baseDir, sub)...)
		} else if strings.HasSuffix(item.Name(), ".swift") {
			relPath := item.Name()
			if relDir != "" {
				relPath = relDir + "/" + item.Name()
			}
			// pbxproj에서 특수문자(+, 공백 등) 포함 경로는 따옴표 필요
			quotedPath := relPath
			if strings.ContainsAny(relPath, "+ ") {
				quotedPath = "\"" + relPath + "\""
			}
			seed := b.name + ".ref." + relPath
			entries = append(entries, fileEntry{
				refID:   genUUID(seed),
				buildID: genUUID(b.name + ".build." + relPath),
				name:    item.Name(),  // 주석용 — 따옴표 없이
				path:    quotedPath,   // pbxproj path — 특수문자 시 따옴표
			})
		}
	}
	return entries
}

func (b *pbxBuilder) testsFile() fileEntry {
	return fileEntry{
		genUUID(b.name + ".ref.tests"), genUUID(b.name + ".build.tests"),
		b.name + "Tests.swift", b.name + "Tests.swift",
	}
}

func (b *pbxBuilder) uiTestsFiles() []fileEntry {
	return []fileEntry{
		{genUUID(b.name + ".ref.uitests"), genUUID(b.name + ".build.uitests"),
			b.name + "UITests.swift", b.name + "UITests.swift"},
		{genUUID(b.name + ".ref.uitestslaunch"), genUUID(b.name + ".build.uitestslaunch"),
			b.name + "UITestsLaunchTests.swift", b.name + "UITestsLaunchTests.swift"},
	}
}

func (b *pbxBuilder) build() string {
	var s strings.Builder
	s.WriteString("// !$*UTF8*$!\n{\n\tarchiveVersion = 1;\n\tclasses = {\n\t};\n\tobjectVersion = 56;\n\tobjects = {\n\n")

	b.writeBuildFiles(&s)
	b.writeContainerItemProxy(&s)
	b.writeFileReferences(&s)
	b.writeFrameworksBuildPhases(&s)
	b.writeGroups(&s)
	b.writeNativeTargets(&s)
	b.writeProject(&s)
	b.writeResourcesBuildPhases(&s)
	b.writeSourcesBuildPhases(&s)
	b.writeTargetDependencies(&s)
	b.writeBuildConfigurations(&s)
	b.writeConfigurationLists(&s)

	s.WriteString("\t};\n")
	s.WriteString(fmt.Sprintf("\trootObject = %s /* Project object */;\n", b.projectID))
	s.WriteString("}\n")
	return s.String()
}

func (b *pbxBuilder) writeBuildFiles(s *strings.Builder) {
	s.WriteString("/* Begin PBXBuildFile section */\n")
	for _, f := range b.appSourceFiles() {
		s.WriteString(fmt.Sprintf("\t\t%s /* %s in Sources */ = {isa = PBXBuildFile; fileRef = %s /* %s */; };\n", f.buildID, f.name, f.refID, f.name))
	}
	s.WriteString(fmt.Sprintf("\t\t%s /* Assets.xcassets in Resources */ = {isa = PBXBuildFile; fileRef = %s /* Assets.xcassets */; };\n", b.assetsBuildID, b.assetsRefID))
	if b.cfg.CoreData {
		s.WriteString(fmt.Sprintf("\t\t%s /* %s.xcdatamodeld in Sources */ = {isa = PBXBuildFile; fileRef = %s /* %s.xcdatamodeld */; };\n", b.coreDataBuildID, b.name, b.coreDataRefID, b.name))
	}
	tf := b.testsFile()
	s.WriteString(fmt.Sprintf("\t\t%s /* %s in Sources */ = {isa = PBXBuildFile; fileRef = %s /* %s */; };\n", tf.buildID, tf.name, tf.refID, tf.name))
	for _, f := range b.uiTestsFiles() {
		s.WriteString(fmt.Sprintf("\t\t%s /* %s in Sources */ = {isa = PBXBuildFile; fileRef = %s /* %s */; };\n", f.buildID, f.name, f.refID, f.name))
	}
	s.WriteString("/* End PBXBuildFile section */\n\n")
}

func (b *pbxBuilder) writeContainerItemProxy(s *strings.Builder) {
	s.WriteString("/* Begin PBXContainerItemProxy section */\n")
	// Tests proxy
	s.WriteString(fmt.Sprintf(`		%s /* PBXContainerItemProxy */ = {
			isa = PBXContainerItemProxy;
			containerPortal = %s /* Project object */;
			proxyType = 1;
			remoteGlobalIDString = %s;
			remoteInfo = %s;
		};
`, b.testsProxyID, b.projectID, b.appTargetID, b.name))
	// UITests proxy
	s.WriteString(fmt.Sprintf(`		%s /* PBXContainerItemProxy */ = {
			isa = PBXContainerItemProxy;
			containerPortal = %s /* Project object */;
			proxyType = 1;
			remoteGlobalIDString = %s;
			remoteInfo = %s;
		};
`, b.uiTestsProxyID, b.projectID, b.appTargetID, b.name))
	s.WriteString("/* End PBXContainerItemProxy section */\n\n")
}

func (b *pbxBuilder) writeFileReferences(s *strings.Builder) {
	s.WriteString("/* Begin PBXFileReference section */\n")
	for _, f := range b.appSourceFiles() {
		s.WriteString(fmt.Sprintf("\t\t%s /* %s */ = {isa = PBXFileReference; lastKnownFileType = sourcecode.swift; path = %s; sourceTree = \"<group>\"; };\n", f.refID, f.name, f.path))
	}
	s.WriteString(fmt.Sprintf("\t\t%s /* Assets.xcassets */ = {isa = PBXFileReference; lastKnownFileType = folder.assetcatalog; path = Assets.xcassets; sourceTree = \"<group>\"; };\n", b.assetsRefID))
	if b.cfg.CoreData {
		s.WriteString(fmt.Sprintf("\t\t%s /* %s.xcdatamodeld */ = {isa = PBXFileReference; lastKnownFileType = wrapper.xcdatamodel; name = %s.xcdatamodeld; path = %s.xcdatamodeld; sourceTree = \"<group>\"; };\n", b.coreDataRefID, b.name, b.name, b.name))
	}
	// Products
	s.WriteString(fmt.Sprintf("\t\t%s /* %s.app */ = {isa = PBXFileReference; explicitFileType = wrapper.application; includeInIndex = 0; path = %s.app; sourceTree = BUILT_PRODUCTS_DIR; };\n", b.appProductID, b.name, b.name))
	s.WriteString(fmt.Sprintf("\t\t%s /* %sTests.xctest */ = {isa = PBXFileReference; explicitFileType = wrapper.cfbundle; includeInIndex = 0; path = %sTests.xctest; sourceTree = BUILT_PRODUCTS_DIR; };\n", b.testsProductID, b.name, b.name))
	s.WriteString(fmt.Sprintf("\t\t%s /* %sUITests.xctest */ = {isa = PBXFileReference; explicitFileType = wrapper.cfbundle; includeInIndex = 0; path = %sUITests.xctest; sourceTree = BUILT_PRODUCTS_DIR; };\n", b.uiTestsProductID, b.name, b.name))
	// Test source files
	tf := b.testsFile()
	s.WriteString(fmt.Sprintf("\t\t%s /* %s */ = {isa = PBXFileReference; lastKnownFileType = sourcecode.swift; path = %s; sourceTree = \"<group>\"; };\n", tf.refID, tf.name, tf.path))
	for _, f := range b.uiTestsFiles() {
		s.WriteString(fmt.Sprintf("\t\t%s /* %s */ = {isa = PBXFileReference; lastKnownFileType = sourcecode.swift; path = %s; sourceTree = \"<group>\"; };\n", f.refID, f.name, f.path))
	}
	s.WriteString("/* End PBXFileReference section */\n\n")
}

func (b *pbxBuilder) writeFrameworksBuildPhases(s *strings.Builder) {
	s.WriteString("/* Begin PBXFrameworksBuildPhase section */\n")
	for _, id := range []string{b.frameworksBuildPhaseAppID, b.frameworksBuildPhaseTestsID, b.frameworksBuildPhaseUITestsID} {
		s.WriteString(fmt.Sprintf(`		%s /* Frameworks */ = {
			isa = PBXFrameworksBuildPhase;
			buildActionMask = 2147483647;
			files = (
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
`, id))
	}
	s.WriteString("/* End PBXFrameworksBuildPhase section */\n\n")
}

func (b *pbxBuilder) writeGroups(s *strings.Builder) {
	s.WriteString("/* Begin PBXGroup section */\n")
	// Main group
	s.WriteString(fmt.Sprintf(`		%s = {
			isa = PBXGroup;
			children = (
				%s /* %s */,
				%s /* %sTests */,
				%s /* %sUITests */,
				%s /* Products */,
			);
			sourceTree = "<group>";
		};
`, b.mainGroupID, b.sourcesGroupID, b.name, b.testsGroupID, b.name, b.uiTestsGroupID, b.name, b.productsGroupID))

	// App sources group
	s.WriteString(fmt.Sprintf("\t\t%s /* %s */ = {\n\t\t\tisa = PBXGroup;\n\t\t\tchildren = (\n", b.sourcesGroupID, b.name))
	for _, f := range b.appSourceFiles() {
		s.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s */,\n", f.refID, f.name))
	}
	s.WriteString(fmt.Sprintf("\t\t\t\t%s /* Assets.xcassets */,\n", b.assetsRefID))
	if b.cfg.CoreData {
		s.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s.xcdatamodeld */,\n", b.coreDataRefID, b.name))
	}
	s.WriteString(fmt.Sprintf("\t\t\t);\n\t\t\tpath = %s;\n\t\t\tsourceTree = \"<group>\";\n\t\t};\n", b.name))

	// Products group
	s.WriteString(fmt.Sprintf(`		%s /* Products */ = {
			isa = PBXGroup;
			children = (
				%s /* %s.app */,
				%s /* %sTests.xctest */,
				%s /* %sUITests.xctest */,
			);
			name = Products;
			sourceTree = "<group>";
		};
`, b.productsGroupID, b.appProductID, b.name, b.testsProductID, b.name, b.uiTestsProductID, b.name))

	// Tests group
	tf := b.testsFile()
	s.WriteString(fmt.Sprintf(`		%s /* %sTests */ = {
			isa = PBXGroup;
			children = (
				%s /* %s */,
			);
			path = %sTests;
			sourceTree = "<group>";
		};
`, b.testsGroupID, b.name, tf.refID, tf.name, b.name))

	// UITests group
	s.WriteString(fmt.Sprintf("\t\t%s /* %sUITests */ = {\n\t\t\tisa = PBXGroup;\n\t\t\tchildren = (\n", b.uiTestsGroupID, b.name))
	for _, f := range b.uiTestsFiles() {
		s.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s */,\n", f.refID, f.name))
	}
	s.WriteString(fmt.Sprintf("\t\t\t);\n\t\t\tpath = %sUITests;\n\t\t\tsourceTree = \"<group>\";\n\t\t};\n", b.name))

	s.WriteString("/* End PBXGroup section */\n\n")
}

func (b *pbxBuilder) writeNativeTargets(s *strings.Builder) {
	s.WriteString("/* Begin PBXNativeTarget section */\n")
	// App target
	s.WriteString(fmt.Sprintf(`		%s /* %s */ = {
			isa = PBXNativeTarget;
			buildConfigurationList = %s /* Build configuration list for PBXNativeTarget "%s" */;
			buildPhases = (
				%s /* Sources */,
				%s /* Frameworks */,
				%s /* Resources */,
			);
			buildRules = (
			);
			dependencies = (
			);
			name = %s;
			productName = %s;
			productReference = %s /* %s.app */;
			productType = "com.apple.product-type.application";
		};
`, b.appTargetID, b.name, b.configListAppID, b.name, b.sourcesBuildPhaseAppID, b.frameworksBuildPhaseAppID, b.resourcesBuildPhaseAppID, b.name, b.name, b.appProductID, b.name))

	// Tests target
	s.WriteString(fmt.Sprintf(`		%s /* %sTests */ = {
			isa = PBXNativeTarget;
			buildConfigurationList = %s /* Build configuration list for PBXNativeTarget "%sTests" */;
			buildPhases = (
				%s /* Sources */,
				%s /* Frameworks */,
				%s /* Resources */,
			);
			buildRules = (
			);
			dependencies = (
				%s /* PBXTargetDependency */,
			);
			name = %sTests;
			productName = %sTests;
			productReference = %s /* %sTests.xctest */;
			productType = "com.apple.product-type.bundle.unit-test";
		};
`, b.testsTargetID, b.name, b.configListTestsID, b.name, b.sourcesBuildPhaseTestsID, b.frameworksBuildPhaseTestsID, b.resourcesBuildPhaseTestsID, b.testsDependencyID, b.name, b.name, b.testsProductID, b.name))

	// UITests target
	s.WriteString(fmt.Sprintf(`		%s /* %sUITests */ = {
			isa = PBXNativeTarget;
			buildConfigurationList = %s /* Build configuration list for PBXNativeTarget "%sUITests" */;
			buildPhases = (
				%s /* Sources */,
				%s /* Frameworks */,
				%s /* Resources */,
			);
			buildRules = (
			);
			dependencies = (
				%s /* PBXTargetDependency */,
			);
			name = %sUITests;
			productName = %sUITests;
			productReference = %s /* %sUITests.xctest */;
			productType = "com.apple.product-type.bundle.ui-testing";
		};
`, b.uiTestsTargetID, b.name, b.configListUITestsID, b.name, b.sourcesBuildPhaseUITestsID, b.frameworksBuildPhaseUITestsID, b.resourcesBuildPhaseUITestsID, b.uiTestsDependencyID, b.name, b.name, b.uiTestsProductID, b.name))

	s.WriteString("/* End PBXNativeTarget section */\n\n")
}

func (b *pbxBuilder) writeProject(s *strings.Builder) {
	s.WriteString(fmt.Sprintf(`/* Begin PBXProject section */
		%s /* Project object */ = {
			isa = PBXProject;
			attributes = {
				BuildIndependentTargetsInParallel = 1;
				LastSwiftUpdateCheck = 1600;
				LastUpgradeCheck = 1600;
				TargetAttributes = {
					%s = {
						TestTargetID = %s;
					};
					%s = {
						TestTargetID = %s;
					};
				};
			};
			buildConfigurationList = %s /* Build configuration list for PBXProject "%s" */;
			compatibilityVersion = "Xcode 14.0";
			developmentRegion = en;
			hasScannedForEncodings = 0;
			knownRegions = (
				en,
				Base,
			);
			mainGroup = %s;
			productRefGroup = %s /* Products */;
			projectDirPath = "";
			projectRoot = "";
			targets = (
				%s /* %s */,
				%s /* %sTests */,
				%s /* %sUITests */,
			);
		};
/* End PBXProject section */

`, b.projectID, b.testsTargetID, b.appTargetID, b.uiTestsTargetID, b.appTargetID, b.configListProjectID, b.name, b.mainGroupID, b.productsGroupID, b.appTargetID, b.name, b.testsTargetID, b.name, b.uiTestsTargetID, b.name))
}

func (b *pbxBuilder) writeResourcesBuildPhases(s *strings.Builder) {
	s.WriteString("/* Begin PBXResourcesBuildPhase section */\n")
	s.WriteString(fmt.Sprintf(`		%s /* Resources */ = {
			isa = PBXResourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
				%s /* Assets.xcassets in Resources */,
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
`, b.resourcesBuildPhaseAppID, b.assetsBuildID))
	for _, id := range []string{b.resourcesBuildPhaseTestsID, b.resourcesBuildPhaseUITestsID} {
		s.WriteString(fmt.Sprintf(`		%s /* Resources */ = {
			isa = PBXResourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
`, id))
	}
	s.WriteString("/* End PBXResourcesBuildPhase section */\n\n")
}

func (b *pbxBuilder) writeSourcesBuildPhases(s *strings.Builder) {
	s.WriteString("/* Begin PBXSourcesBuildPhase section */\n")
	// App sources
	s.WriteString(fmt.Sprintf("\t\t%s /* Sources */ = {\n\t\t\tisa = PBXSourcesBuildPhase;\n\t\t\tbuildActionMask = 2147483647;\n\t\t\tfiles = (\n", b.sourcesBuildPhaseAppID))
	for _, f := range b.appSourceFiles() {
		s.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s in Sources */,\n", f.buildID, f.name))
	}
	if b.cfg.CoreData {
		s.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s.xcdatamodeld in Sources */,\n", b.coreDataBuildID, b.name))
	}
	s.WriteString("\t\t\t);\n\t\t\trunOnlyForDeploymentPostprocessing = 0;\n\t\t};\n")

	// Tests sources
	tf := b.testsFile()
	s.WriteString(fmt.Sprintf(`		%s /* Sources */ = {
			isa = PBXSourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
				%s /* %s in Sources */,
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
`, b.sourcesBuildPhaseTestsID, tf.buildID, tf.name))

	// UITests sources
	s.WriteString(fmt.Sprintf("\t\t%s /* Sources */ = {\n\t\t\tisa = PBXSourcesBuildPhase;\n\t\t\tbuildActionMask = 2147483647;\n\t\t\tfiles = (\n", b.sourcesBuildPhaseUITestsID))
	for _, f := range b.uiTestsFiles() {
		s.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s in Sources */,\n", f.buildID, f.name))
	}
	s.WriteString("\t\t\t);\n\t\t\trunOnlyForDeploymentPostprocessing = 0;\n\t\t};\n")

	s.WriteString("/* End PBXSourcesBuildPhase section */\n\n")
}

func (b *pbxBuilder) writeTargetDependencies(s *strings.Builder) {
	s.WriteString("/* Begin PBXTargetDependency section */\n")
	s.WriteString(fmt.Sprintf(`		%s /* PBXTargetDependency */ = {
			isa = PBXTargetDependency;
			target = %s /* %s */;
			targetProxy = %s /* PBXContainerItemProxy */;
		};
`, b.testsDependencyID, b.appTargetID, b.name, b.testsProxyID))
	s.WriteString(fmt.Sprintf(`		%s /* PBXTargetDependency */ = {
			isa = PBXTargetDependency;
			target = %s /* %s */;
			targetProxy = %s /* PBXContainerItemProxy */;
		};
`, b.uiTestsDependencyID, b.appTargetID, b.name, b.uiTestsProxyID))
	s.WriteString("/* End PBXTargetDependency section */\n\n")
}

func (b *pbxBuilder) writeBuildConfigurations(s *strings.Builder) {
	deviceFamilySetting := ""
	if b.deviceFamily != "" {
		deviceFamilySetting = fmt.Sprintf("\n\t\t\t\tTARGETED_DEVICE_FAMILY = %s;", b.deviceFamily)
	}

	s.WriteString("/* Begin XCBuildConfiguration section */\n")

	// Project Debug
	s.WriteString(fmt.Sprintf(`		%s /* Debug */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ALWAYS_SEARCH_USER_PATHS = NO;
				CLANG_ENABLE_MODULES = YES;
				CLANG_ENABLE_OBJC_ARC = YES;
				COPY_PHASE_STRIP = NO;
				DEBUG_INFORMATION_FORMAT = dwarf;
				ENABLE_STRICT_OBJC_MSGSEND = YES;
				ENABLE_TESTABILITY = YES;
				GCC_DYNAMIC_NO_PIC = NO;
				GCC_OPTIMIZATION_LEVEL = 0;
				GCC_PREPROCESSOR_DEFINITIONS = (
					"DEBUG=1",
					"$(inherited)",
				);
				MTL_ENABLE_DEBUG_INFO = INCLUDE_SOURCE;
				ONLY_ACTIVE_ARCH = YES;
				SDKROOT = %s;
				SUPPORTED_PLATFORMS = %s;
				SWIFT_ACTIVE_COMPILATION_CONDITIONS = "$(inherited) DEBUG";
				SWIFT_OPTIMIZATION_LEVEL = "-Onone";
				%s
			};
			name = Debug;
		};
`, b.debugConfigProjectID, b.sdkroot, b.supportedPlatforms, b.deploymentSetting))

	// Project Release
	s.WriteString(fmt.Sprintf(`		%s /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ALWAYS_SEARCH_USER_PATHS = NO;
				CLANG_ENABLE_MODULES = YES;
				CLANG_ENABLE_OBJC_ARC = YES;
				COPY_PHASE_STRIP = NO;
				DEBUG_INFORMATION_FORMAT = "dwarf-with-dsym";
				ENABLE_NS_ASSERTIONS = NO;
				ENABLE_STRICT_OBJC_MSGSEND = YES;
				MTL_ENABLE_DEBUG_INFO = NO;
				SDKROOT = %s;
				SUPPORTED_PLATFORMS = %s;
				SWIFT_COMPILATION_MODE = wholemodule;
				%s
			};
			name = Release;
		};
`, b.releaseConfigProjectID, b.sdkroot, b.supportedPlatforms, b.deploymentSetting))

	// App target Debug/Release
	for _, pair := range []struct{ id, mode string }{{b.debugConfigAppID, "Debug"}, {b.releaseConfigAppID, "Release"}} {
		s.WriteString(fmt.Sprintf(`		%s /* %s */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				ASSETCATALOG_COMPILER_APPICON_NAME = AppIcon;
				ASSETCATALOG_COMPILER_GLOBAL_ACCENT_COLOR_NAME = AccentColor;
				CODE_SIGN_STYLE = Automatic;
				GENERATE_INFOPLIST_FILE = YES;
				INFOPLIST_KEY_UIApplicationSupportsIndirectInputEvents = YES;
				INFOPLIST_KEY_UILaunchScreen_Generation = YES;
				INFOPLIST_KEY_UISupportedInterfaceOrientations_iPad = "UIInterfaceOrientationPortrait UIInterfaceOrientationPortraitUpsideDown UIInterfaceOrientationLandscapeLeft UIInterfaceOrientationLandscapeRight";
				INFOPLIST_KEY_UISupportedInterfaceOrientations_iPhone = "UIInterfaceOrientationPortrait UIInterfaceOrientationLandscapeLeft UIInterfaceOrientationLandscapeRight";
				PRODUCT_BUNDLE_IDENTIFIER = %s;
				PRODUCT_NAME = "$(TARGET_NAME)";
				SWIFT_EMIT_LOC_STRINGS = YES;
				SWIFT_VERSION = %s;%s
				%s
			};
			name = %s;
		};
`, pair.id, pair.mode, b.cfg.BundleID, b.swiftVersion, deviceFamilySetting, b.deploymentSetting, pair.mode))
	}

	// Tests target Debug/Release
	for _, pair := range []struct{ id, mode string }{{b.debugConfigTestsID, "Debug"}, {b.releaseConfigTestsID, "Release"}} {
		s.WriteString(fmt.Sprintf(`		%s /* %s */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				BUNDLE_LOADER = "$(TEST_HOST)";
				CODE_SIGN_STYLE = Automatic;
				GENERATE_INFOPLIST_FILE = YES;
				PRODUCT_BUNDLE_IDENTIFIER = %s.tests;
				PRODUCT_NAME = "$(TARGET_NAME)";
				SWIFT_EMIT_LOC_STRINGS = NO;
				SWIFT_VERSION = %s;%s
				TEST_HOST = "$(BUILT_PRODUCTS_DIR)/%s.app/$(BUNDLE_EXECUTABLE_FOLDER_PATH)/%s";
				%s
			};
			name = %s;
		};
`, pair.id, pair.mode, b.cfg.BundleID, b.swiftVersion, deviceFamilySetting, b.name, b.name, b.deploymentSetting, pair.mode))
	}

	// UITests target Debug/Release
	for _, pair := range []struct{ id, mode string }{{b.debugConfigUITestsID, "Debug"}, {b.releaseConfigUITestsID, "Release"}} {
		s.WriteString(fmt.Sprintf(`		%s /* %s */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
				CODE_SIGN_STYLE = Automatic;
				GENERATE_INFOPLIST_FILE = YES;
				PRODUCT_BUNDLE_IDENTIFIER = %s.uitests;
				PRODUCT_NAME = "$(TARGET_NAME)";
				SWIFT_EMIT_LOC_STRINGS = NO;
				SWIFT_VERSION = %s;%s
				TEST_TARGET_NAME = %s;
				%s
			};
			name = %s;
		};
`, pair.id, pair.mode, b.cfg.BundleID, b.swiftVersion, deviceFamilySetting, b.name, b.deploymentSetting, pair.mode))
	}

	s.WriteString("/* End XCBuildConfiguration section */\n\n")
}

func (b *pbxBuilder) writeConfigurationLists(s *strings.Builder) {
	s.WriteString("/* Begin XCConfigurationList section */\n")

	configs := []struct {
		id, targetType, name, debugID, releaseID string
	}{
		{b.configListProjectID, "PBXProject", b.name, b.debugConfigProjectID, b.releaseConfigProjectID},
		{b.configListAppID, "PBXNativeTarget", b.name, b.debugConfigAppID, b.releaseConfigAppID},
		{b.configListTestsID, "PBXNativeTarget", b.name + "Tests", b.debugConfigTestsID, b.releaseConfigTestsID},
		{b.configListUITestsID, "PBXNativeTarget", b.name + "UITests", b.debugConfigUITestsID, b.releaseConfigUITestsID},
	}

	for _, c := range configs {
		s.WriteString(fmt.Sprintf(`		%s /* Build configuration list for %s "%s" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				%s /* Debug */,
				%s /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
`, c.id, c.targetType, c.name, c.debugID, c.releaseID))
	}

	s.WriteString("/* End XCConfigurationList section */\n")
}

package ios

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// genUUID generates a deterministic 24-char hex UUID from a seed string.
// Xcode uses 24-character uppercase hex strings as object identifiers.
func genUUID(seed string) string {
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%X", h[:12])
}

type pbxGenerator struct {
	cfg     Config
	objects []string
	rootID  string
}

func generatePbxproj(cfg Config) string {
	g := &pbxGenerator{cfg: cfg}
	return g.build()
}

func (g *pbxGenerator) build() string {
	name := g.cfg.Name
	bundleID := g.cfg.BundleID

	// Generate deterministic UUIDs
	projectID := genUUID(name + ".project")
	mainGroupID := genUUID(name + ".mainGroup")
	sourcesGroupID := genUUID(name + ".sourcesGroup")
	productsGroupID := genUUID(name + ".productsGroup")
	appTargetID := genUUID(name + ".appTarget")
	appProductID := genUUID(name + ".appProduct")
	buildConfigListProjectID := genUUID(name + ".configList.project")
	buildConfigListTargetID := genUUID(name + ".configList.target")
	debugConfigProjectID := genUUID(name + ".config.project.debug")
	releaseConfigProjectID := genUUID(name + ".config.project.release")
	debugConfigTargetID := genUUID(name + ".config.target.debug")
	releaseConfigTargetID := genUUID(name + ".config.target.release")
	sourcesBuildPhaseID := genUUID(name + ".sourcesBuildPhase")
	resourcesBuildPhaseID := genUUID(name + ".resourcesBuildPhase")
	frameworksBuildPhaseID := genUUID(name + ".frameworksBuildPhase")

	// Source files
	type fileEntry struct {
		refID   string
		buildID string
		name    string
		path    string
	}

	var sourceFiles []fileEntry
	if g.cfg.Framework == "swiftui" {
		sourceFiles = []fileEntry{
			{genUUID(name + ".ref.app"), genUUID(name + ".build.app"), name + "App.swift", name + "App.swift"},
			{genUUID(name + ".ref.contentview"), genUUID(name + ".build.contentview"), "ContentView.swift", "ContentView.swift"},
		}
	} else {
		sourceFiles = []fileEntry{
			{genUUID(name + ".ref.appdelegate"), genUUID(name + ".build.appdelegate"), "AppDelegate.swift", "AppDelegate.swift"},
			{genUUID(name + ".ref.scenedelegate"), genUUID(name + ".build.scenedelegate"), "SceneDelegate.swift", "SceneDelegate.swift"},
			{genUUID(name + ".ref.viewcontroller"), genUUID(name + ".build.viewcontroller"), "ViewController.swift", "ViewController.swift"},
		}
	}

	// Assets
	assetsRefID := genUUID(name + ".ref.assets")
	assetsBuildID := genUUID(name + ".build.assets")

	// Deployment target key/value
	deploymentTargetSetting := ""
	sdkroot := ""
	targetedDeviceFamily := ""
	supportedPlatforms := ""
	if g.cfg.Platform == "ios" {
		deploymentTargetSetting = fmt.Sprintf("IPHONEOS_DEPLOYMENT_TARGET = %s;", g.cfg.MinVersion)
		sdkroot = "iphoneos"
		targetedDeviceFamily = "\"1,2\""
		supportedPlatforms = "\"iphoneos iphonesimulator\""
	} else {
		deploymentTargetSetting = fmt.Sprintf("MACOSX_DEPLOYMENT_TARGET = %s;", g.cfg.MinVersion)
		sdkroot = "macosx"
		targetedDeviceFamily = ""
		supportedPlatforms = "\"macosx\""
	}

	swiftVersion := "5.0"
	if g.cfg.XcodeInfo != nil && g.cfg.XcodeInfo.SwiftVersion != "" {
		// Use major.minor from detected swift
		parts := strings.Split(g.cfg.XcodeInfo.SwiftVersion, ".")
		if len(parts) >= 2 {
			swiftVersion = parts[0] + "." + parts[1]
		} else {
			swiftVersion = parts[0] + ".0"
		}
	}

	// Build PBXBuildFile section
	var buildFileSection strings.Builder
	for _, f := range sourceFiles {
		buildFileSection.WriteString(fmt.Sprintf("\t\t%s /* %s in Sources */ = {isa = PBXBuildFile; fileRef = %s /* %s */; };\n",
			f.buildID, f.name, f.refID, f.name))
	}
	buildFileSection.WriteString(fmt.Sprintf("\t\t%s /* Assets.xcassets in Resources */ = {isa = PBXBuildFile; fileRef = %s /* Assets.xcassets */; };\n",
		assetsBuildID, assetsRefID))

	// Build PBXFileReference section
	var fileRefSection strings.Builder
	for _, f := range sourceFiles {
		fileRefSection.WriteString(fmt.Sprintf("\t\t%s /* %s */ = {isa = PBXFileReference; lastKnownFileType = sourcecode.swift; path = %s; sourceTree = \"<group>\"; };\n",
			f.refID, f.name, f.path))
	}
	fileRefSection.WriteString(fmt.Sprintf("\t\t%s /* Assets.xcassets */ = {isa = PBXFileReference; lastKnownFileType = folder.assetcatalog; path = Assets.xcassets; sourceTree = \"<group>\"; };\n",
		assetsRefID))
	fileRefSection.WriteString(fmt.Sprintf("\t\t%s /* %s.app */ = {isa = PBXFileReference; explicitFileType = wrapper.application; includeInIndex = 0; path = %s.app; sourceTree = BUILT_PRODUCTS_DIR; };\n",
		appProductID, name, name))

	// Sources group children
	var sourcesChildren strings.Builder
	for _, f := range sourceFiles {
		sourcesChildren.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s */,\n", f.refID, f.name))
	}
	sourcesChildren.WriteString(fmt.Sprintf("\t\t\t\t%s /* Assets.xcassets */,\n", assetsRefID))

	// Build sources list
	var buildFilesList strings.Builder
	for _, f := range sourceFiles {
		buildFilesList.WriteString(fmt.Sprintf("\t\t\t\t%s /* %s in Sources */,\n", f.buildID, f.name))
	}

	// Targeted device family setting
	deviceFamilySetting := ""
	if targetedDeviceFamily != "" {
		deviceFamilySetting = fmt.Sprintf("\n\t\t\t\tTARGETED_DEVICE_FAMILY = %s;", targetedDeviceFamily)
	}

	return fmt.Sprintf(`// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 56;
	objects = {

/* Begin PBXBuildFile section */
%s/* End PBXBuildFile section */

/* Begin PBXFileReference section */
%s/* End PBXFileReference section */

/* Begin PBXFrameworksBuildPhase section */
		%s /* Frameworks */ = {
			isa = PBXFrameworksBuildPhase;
			buildActionMask = 2147483647;
			files = (
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
/* End PBXFrameworksBuildPhase section */

/* Begin PBXGroup section */
		%s = {
			isa = PBXGroup;
			children = (
				%s /* %s */,
				%s /* Products */,
			);
			sourceTree = "<group>";
		};
		%s /* %s */ = {
			isa = PBXGroup;
			children = (
%s			);
			path = %s;
			sourceTree = "<group>";
		};
		%s /* Products */ = {
			isa = PBXGroup;
			children = (
				%s /* %s.app */,
			);
			name = Products;
			sourceTree = "<group>";
		};
/* End PBXGroup section */

/* Begin PBXNativeTarget section */
		%s /* %s */ = {
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
/* End PBXNativeTarget section */

/* Begin PBXProject section */
		%s /* Project object */ = {
			isa = PBXProject;
			attributes = {
				BuildIndependentTargetsInParallel = 1;
				LastSwiftUpdateCheck = 1600;
				LastUpgradeCheck = 1600;
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
			);
		};
/* End PBXProject section */

/* Begin PBXResourcesBuildPhase section */
		%s /* Resources */ = {
			isa = PBXResourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
				%s /* Assets.xcassets in Resources */,
			);
			runOnlyForDeploymentPostprocessing = 0;
		};
/* End PBXResourcesBuildPhase section */

/* Begin PBXSourcesBuildPhase section */
		%s /* Sources */ = {
			isa = PBXSourcesBuildPhase;
			buildActionMask = 2147483647;
			files = (
%s			);
			runOnlyForDeploymentPostprocessing = 0;
		};
/* End PBXSourcesBuildPhase section */

/* Begin XCBuildConfiguration section */
		%s /* Debug */ = {
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
		%s /* Release */ = {
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
		%s /* Debug */ = {
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
			name = Debug;
		};
		%s /* Release */ = {
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
			name = Release;
		};
/* End XCBuildConfiguration section */

/* Begin XCConfigurationList section */
		%s /* Build configuration list for PBXProject "%s" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				%s /* Debug */,
				%s /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
		%s /* Build configuration list for PBXNativeTarget "%s" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				%s /* Debug */,
				%s /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
/* End XCConfigurationList section */
	};
	rootObject = %s /* Project object */;
}
`,
		// PBXBuildFile
		buildFileSection.String(),
		// PBXFileReference
		fileRefSection.String(),
		// PBXFrameworksBuildPhase
		frameworksBuildPhaseID,
		// PBXGroup - main
		mainGroupID,
		sourcesGroupID, name,
		productsGroupID,
		// PBXGroup - sources
		sourcesGroupID, name,
		sourcesChildren.String(),
		name,
		// PBXGroup - products
		productsGroupID,
		appProductID, name,
		// PBXNativeTarget
		appTargetID, name,
		buildConfigListTargetID, name,
		sourcesBuildPhaseID,
		frameworksBuildPhaseID,
		resourcesBuildPhaseID,
		name, name,
		appProductID, name,
		// PBXProject
		projectID,
		buildConfigListProjectID, name,
		mainGroupID,
		productsGroupID,
		appTargetID, name,
		// PBXResourcesBuildPhase
		resourcesBuildPhaseID,
		assetsBuildID,
		// PBXSourcesBuildPhase
		sourcesBuildPhaseID,
		buildFilesList.String(),
		// XCBuildConfiguration - project debug
		debugConfigProjectID,
		sdkroot, supportedPlatforms, deploymentTargetSetting,
		// XCBuildConfiguration - project release
		releaseConfigProjectID,
		sdkroot, supportedPlatforms, deploymentTargetSetting,
		// XCBuildConfiguration - target debug
		debugConfigTargetID,
		bundleID, swiftVersion, deviceFamilySetting, deploymentTargetSetting,
		// XCBuildConfiguration - target release
		releaseConfigTargetID,
		bundleID, swiftVersion, deviceFamilySetting, deploymentTargetSetting,
		// XCConfigurationList - project
		buildConfigListProjectID, name,
		debugConfigProjectID, releaseConfigProjectID,
		// XCConfigurationList - target
		buildConfigListTargetID, name,
		debugConfigTargetID, releaseConfigTargetID,
		// rootObject
		projectID,
	)
}

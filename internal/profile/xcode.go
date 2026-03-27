package profile

// XcodeProfile defines version-specific settings for Xcode project generation.
type XcodeProfile struct {
	ObjectVersion        int
	CompatibilityVersion string
	CheckVersion         string // e.g. "1600" for LastSwiftUpdateCheck/LastUpgradeCheck
	DefaultSwiftVersion  string
	DefaultMinIOS        string
	DefaultMinMacOS      string

	// Feature flags
	GenerateInfoPlist       bool     // GENERATE_INFOPLIST_FILE = YES (no physical Info.plist)
	SingleAppIcon           bool     // Single 1024x1024 vs multi-slot
	PreviewContent          bool     // Preview Content/Preview Assets.xcassets
	UsePreviewMacro         bool     // #Preview {} vs PreviewProvider struct
	PrivacyManifest         bool     // PrivacyInfo.xcprivacy
	UserScriptSandboxing    bool     // ENABLE_USER_SCRIPT_SANDBOXING
	StrictConcurrency       bool     // SWIFT_STRICT_CONCURRENCY = complete
	SwiftTesting            bool     // Swift Testing framework
	ExplicitWarningSettings bool     // 30+ CLANG_WARN_/GCC_WARN_ settings
	NewPbxprojFormat        bool     // objectVersion 77 (Xcode 16 sidecar format)

	// Language standards
	CStandard   string // gnu11 or gnu17
	CXXStandard string // gnu++17 or gnu++20

	// Test class modifier
	TestClassModifier string // "class" or "final class"

	// ContentView style
	ForegroundModifier string // ".foregroundColor(.accentColor)" or ".foregroundStyle(.tint)"
}

var XcodeProfiles = map[string]XcodeProfile{
	"13": {
		ObjectVersion:        55,
		CompatibilityVersion: "Xcode 13.0",
		CheckVersion:         "1300",
		DefaultSwiftVersion:  "5.5",
		DefaultMinIOS:        "15.0",
		DefaultMinMacOS:      "12.0",

		GenerateInfoPlist:       false,
		SingleAppIcon:           false,
		PreviewContent:          false,
		UsePreviewMacro:         false,
		PrivacyManifest:         false,
		UserScriptSandboxing:    false,
		StrictConcurrency:       false,
		SwiftTesting:            false,
		ExplicitWarningSettings: false,
		NewPbxprojFormat:        false,

		CStandard:          "gnu11",
		CXXStandard:        "gnu++17",
		TestClassModifier:  "class",
		ForegroundModifier: ".foregroundColor(.accentColor)",
	},
	"14": {
		ObjectVersion:        56,
		CompatibilityVersion: "Xcode 14.0",
		CheckVersion:         "1400",
		DefaultSwiftVersion:  "5.7",
		DefaultMinIOS:        "16.0",
		DefaultMinMacOS:      "13.0",

		GenerateInfoPlist:       true,
		SingleAppIcon:           true,
		PreviewContent:          true,
		UsePreviewMacro:         false,
		PrivacyManifest:         false,
		UserScriptSandboxing:    true,
		StrictConcurrency:       false,
		SwiftTesting:            false,
		ExplicitWarningSettings: true,
		NewPbxprojFormat:        false,

		CStandard:          "gnu17",
		CXXStandard:        "gnu++20",
		TestClassModifier:  "final class",
		ForegroundModifier: ".foregroundStyle(.tint)",
	},
	"15": {
		ObjectVersion:        56,
		CompatibilityVersion: "Xcode 14.0",
		CheckVersion:         "1500",
		DefaultSwiftVersion:  "5.9",
		DefaultMinIOS:        "17.0",
		DefaultMinMacOS:      "14.0",

		GenerateInfoPlist:       true,
		SingleAppIcon:           true,
		PreviewContent:          true,
		UsePreviewMacro:         true,
		PrivacyManifest:         true,
		UserScriptSandboxing:    true,
		StrictConcurrency:       false,
		SwiftTesting:            false,
		ExplicitWarningSettings: true,
		NewPbxprojFormat:        false,

		CStandard:          "gnu17",
		CXXStandard:        "gnu++20",
		TestClassModifier:  "final class",
		ForegroundModifier: ".foregroundStyle(.tint)",
	},
	"16": {
		ObjectVersion:        56, // Using 56 for compatibility; 77 is new sidecar format
		CompatibilityVersion: "Xcode 14.0",
		CheckVersion:         "1600",
		DefaultSwiftVersion:  "6.0",
		DefaultMinIOS:        "18.0",
		DefaultMinMacOS:      "15.0",

		GenerateInfoPlist:       true,
		SingleAppIcon:           true,
		PreviewContent:          true,
		UsePreviewMacro:         true,
		PrivacyManifest:         true,
		UserScriptSandboxing:    true,
		StrictConcurrency:       true,
		SwiftTesting:            true,
		ExplicitWarningSettings: true,
		NewPbxprojFormat:        false, // Keep false for now; 77 format is too different

		CStandard:          "gnu17",
		CXXStandard:        "gnu++20",
		TestClassModifier:  "final class",
		ForegroundModifier: ".foregroundStyle(.tint)",
	},
}

// GetXcodeProfile returns the profile for the given major version.
// Falls back to the closest known version.
func GetXcodeProfile(majorVersion string) XcodeProfile {
	if p, ok := XcodeProfiles[majorVersion]; ok {
		return p
	}
	// Default to Xcode 15 as safe middle ground
	return XcodeProfiles["15"]
}

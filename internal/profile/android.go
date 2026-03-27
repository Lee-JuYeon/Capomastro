package profile

// AndroidProfile defines version-specific settings for Android project generation.
type AndroidProfile struct {
	Name          string
	BuildPrefix   string // e.g. "243" for Meerkat
	AGP           string
	Gradle        string
	Kotlin        string
	ComposeBOM    string
	CompileSdk    string
	TargetSdk     string
	MinSdk        string
	JavaVersion   string // "11" or "17"

	// Format flags
	UseKTS            bool // .gradle.kts vs .gradle
	UseVersionCatalog bool // libs.versions.toml
	UseComposePlugin  bool // kotlin-compose plugin (vs composeOptions)
	NamespaceInBuild  bool // namespace in build.gradle vs AndroidManifest

	// Compose compiler (empty = use plugin, no composeOptions needed)
	ComposeCompilerExt string
}

var AndroidProfiles = map[string]AndroidProfile{
	// AGP 7.x era (Groovy default, no version catalog)
	"2020.3": {Name: "Arctic Fox", BuildPrefix: "203", AGP: "7.0.4", Gradle: "7.0.2", Kotlin: "1.5.31", ComposeBOM: "", CompileSdk: "31", TargetSdk: "31", MinSdk: "21", JavaVersion: "11", UseKTS: false, UseVersionCatalog: false, UseComposePlugin: false, NamespaceInBuild: false, ComposeCompilerExt: "1.0.5"},
	"2021.1": {Name: "Bumblebee", BuildPrefix: "211", AGP: "7.1.3", Gradle: "7.2", Kotlin: "1.6.21", ComposeBOM: "", CompileSdk: "32", TargetSdk: "32", MinSdk: "21", JavaVersion: "11", UseKTS: false, UseVersionCatalog: false, UseComposePlugin: false, NamespaceInBuild: false, ComposeCompilerExt: "1.2.0"},
	"2021.2": {Name: "Chipmunk", BuildPrefix: "212", AGP: "7.2.2", Gradle: "7.3.3", Kotlin: "1.7.10", ComposeBOM: "", CompileSdk: "32", TargetSdk: "32", MinSdk: "21", JavaVersion: "11", UseKTS: false, UseVersionCatalog: false, UseComposePlugin: false, NamespaceInBuild: false, ComposeCompilerExt: "1.3.0"},
	"2021.3": {Name: "Dolphin", BuildPrefix: "213", AGP: "7.3.1", Gradle: "7.4", Kotlin: "1.7.20", ComposeBOM: "", CompileSdk: "33", TargetSdk: "33", MinSdk: "21", JavaVersion: "11", UseKTS: false, UseVersionCatalog: false, UseComposePlugin: false, NamespaceInBuild: false, ComposeCompilerExt: "1.3.2"},

	// AGP 8.0 transition era (KTS default, namespace in build)
	"2022.1": {Name: "Electric Eel", BuildPrefix: "221", AGP: "7.4.2", Gradle: "7.5", Kotlin: "1.8.0", ComposeBOM: "2023.01.00", CompileSdk: "33", TargetSdk: "33", MinSdk: "24", JavaVersion: "11", UseKTS: true, UseVersionCatalog: false, UseComposePlugin: false, NamespaceInBuild: true, ComposeCompilerExt: "1.4.0"},
	"2022.2": {Name: "Flamingo", BuildPrefix: "222", AGP: "8.0.2", Gradle: "8.0", Kotlin: "1.8.10", ComposeBOM: "2023.03.00", CompileSdk: "33", TargetSdk: "33", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: false, UseComposePlugin: false, NamespaceInBuild: true, ComposeCompilerExt: "1.4.3"},
	"2022.3": {Name: "Giraffe", BuildPrefix: "223", AGP: "8.1.4", Gradle: "8.0", Kotlin: "1.9.0", ComposeBOM: "2023.08.00", CompileSdk: "34", TargetSdk: "34", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: false, NamespaceInBuild: true, ComposeCompilerExt: "1.5.1"},

	// AGP 8.2+ era (version catalog default)
	"2023.1": {Name: "Hedgehog", BuildPrefix: "231", AGP: "8.2.2", Gradle: "8.2", Kotlin: "1.9.20", ComposeBOM: "2023.10.01", CompileSdk: "34", TargetSdk: "34", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: false, NamespaceInBuild: true, ComposeCompilerExt: "1.5.4"},
	"2023.2": {Name: "Iguana", BuildPrefix: "232", AGP: "8.3.2", Gradle: "8.4", Kotlin: "1.9.22", ComposeBOM: "2024.02.00", CompileSdk: "34", TargetSdk: "34", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: false, NamespaceInBuild: true, ComposeCompilerExt: "1.5.10"},
	"2023.3": {Name: "Jellyfish", BuildPrefix: "233", AGP: "8.4.2", Gradle: "8.6", Kotlin: "1.9.23", ComposeBOM: "2024.04.01", CompileSdk: "34", TargetSdk: "34", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: false, NamespaceInBuild: true, ComposeCompilerExt: "1.5.13"},

	// AGP 8.5+ era (kotlin-compose plugin, no composeOptions)
	"2024.1": {Name: "Koala", BuildPrefix: "241", AGP: "8.5.2", Gradle: "8.7", Kotlin: "2.0.0", ComposeBOM: "2024.06.00", CompileSdk: "34", TargetSdk: "34", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: true, NamespaceInBuild: true, ComposeCompilerExt: ""},
	"2024.2": {Name: "Ladybug", BuildPrefix: "242", AGP: "8.7.3", Gradle: "8.9", Kotlin: "2.0.21", ComposeBOM: "2024.09.00", CompileSdk: "35", TargetSdk: "35", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: true, NamespaceInBuild: true, ComposeCompilerExt: ""},
	"2024.3": {Name: "Meerkat", BuildPrefix: "243", AGP: "8.10.1", Gradle: "8.11.1", Kotlin: "2.0.21", ComposeBOM: "2024.09.00", CompileSdk: "35", TargetSdk: "35", MinSdk: "24", JavaVersion: "17", UseKTS: true, UseVersionCatalog: true, UseComposePlugin: true, NamespaceInBuild: true, ComposeCompilerExt: ""},
}

// GetAndroidProfile returns the profile for the given year.major key (e.g. "2024.3").
func GetAndroidProfile(key string) AndroidProfile {
	if p, ok := AndroidProfiles[key]; ok {
		return p
	}
	// Default to Meerkat
	return AndroidProfiles["2024.3"]
}

package detector

import (
	"fmt"
	"os/exec"
	"strings"
)

type XcodeInfo struct {
	Version         string // e.g. "16.2"
	Build           string // e.g. "16C5032a"
	SwiftVersion    string // e.g. "6.0"
	SDKVersion      string // e.g. "18.2"
	DefaultMinIOS   string // safe minimum iOS version for this Xcode
	DefaultMinMacOS string // safe minimum macOS version for this Xcode
}

// xcodeVersionToDefaults maps major Xcode version to recommended defaults.
var xcodeVersionToDefaults = map[string]struct {
	swift  string
	minIOS string
	minMac string
}{
	"16": {swift: "6.0", minIOS: "16.0", minMac: "13.0"},
	"15": {swift: "5.9", minIOS: "15.0", minMac: "13.0"},
	"14": {swift: "5.7", minIOS: "14.0", minMac: "12.0"},
	"13": {swift: "5.5", minIOS: "13.0", minMac: "11.0"},
}

func DetectXcode() (*XcodeInfo, error) {
	info := &XcodeInfo{}

	// Try xcodebuild -version first (requires full Xcode)
	out, err := exec.Command("xcodebuild", "-version").CombinedOutput()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) >= 2 {
			// Parse "Xcode 16.2"
			parts := strings.Fields(lines[0])
			if len(parts) >= 2 {
				info.Version = parts[1]
			}
			// Parse "Build version 16C5032a"
			buildParts := strings.Fields(lines[1])
			if len(buildParts) >= 3 {
				info.Build = buildParts[2]
			}
		}
	}

	// Detect Swift version (works with Command Line Tools too)
	swiftOut, err := exec.Command("swift", "--version").Output()
	if err == nil {
		for _, field := range strings.Fields(string(swiftOut)) {
			if strings.Count(field, ".") >= 1 && field[0] >= '0' && field[0] <= '9' {
				info.SwiftVersion = field
				break
			}
		}
	} else if info.Version == "" {
		return nil, fmt.Errorf("neither Xcode nor Swift found — install Xcode or Command Line Tools")
	}

	// Detect SDK version
	sdkOut, err := exec.Command("xcrun", "--show-sdk-version", "--sdk", "iphoneos").Output()
	if err == nil {
		info.SDKVersion = strings.TrimSpace(string(sdkOut))
	} else {
		// Try macOS SDK as fallback
		sdkOut, err = exec.Command("xcrun", "--show-sdk-version", "--sdk", "macosx").Output()
		if err == nil {
			info.SDKVersion = strings.TrimSpace(string(sdkOut))
		}
	}

	// If no Xcode version, infer from Swift version
	if info.Version == "" && info.SwiftVersion != "" {
		info.Version = swiftToXcodeVersion(info.SwiftVersion)
		fmt.Printf("Note: Full Xcode not found. Using Swift %s to infer settings.\n", info.SwiftVersion)
	}

	// Set defaults based on major version
	major := strings.Split(info.Version, ".")[0]
	if defaults, ok := xcodeVersionToDefaults[major]; ok {
		if info.SwiftVersion == "" {
			info.SwiftVersion = defaults.swift
		}
		info.DefaultMinIOS = defaults.minIOS
		info.DefaultMinMacOS = defaults.minMac
	} else {
		info.DefaultMinIOS = "15.0"
		info.DefaultMinMacOS = "13.0"
		if info.SwiftVersion == "" {
			info.SwiftVersion = "5.9"
		}
	}

	return info, nil
}

func swiftToXcodeVersion(swiftVer string) string {
	major := strings.Split(swiftVer, ".")[0]
	switch major {
	case "6":
		return "16.0"
	case "5":
		return "15.0"
	default:
		return "15.0"
	}
}

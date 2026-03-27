package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cavss/ProjectBuilder/internal/detector"
	"github.com/cavss/ProjectBuilder/internal/generator/android"
	"github.com/cavss/ProjectBuilder/internal/generator/flutter"
	"github.com/cavss/ProjectBuilder/internal/generator/ios"
)

var (
	platform   string
	pkg        string
	name       string
	out        string
	framework  string
	minVersion string
)

func main() {
	root := &cobra.Command{
		Use:   "ProjectBuilder",
		Short: "Generate buildable Xcode/Android Studio projects",
		Long:  "CLI tool that detects your IDE version and generates compatible, buildable projects.",
		RunE:  run,
	}

	root.Flags().StringVar(&platform, "platform", "", "Target platform: ios, macos, android, flutter (required)")
	root.Flags().StringVar(&pkg, "pkg", "", "Package/bundle identifier (e.g. com.example.app) (required)")
	root.Flags().StringVar(&name, "name", "", "Project name (required)")
	root.Flags().StringVar(&out, "out", ".", "Output directory")
	root.Flags().StringVar(&framework, "framework", "", "UI framework: swiftui, uikit (iOS/macOS), compose, xml (Android)")
	root.Flags().StringVar(&minVersion, "min-version", "", "Minimum deployment version override")

	root.MarkFlagRequired("platform")
	root.MarkFlagRequired("pkg")
	root.MarkFlagRequired("name")

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	switch platform {
	case "ios", "macos":
		info, err := detector.DetectXcode()
		if err != nil {
			return fmt.Errorf("Xcode detection failed: %w", err)
		}
		fmt.Printf("Detected Xcode %s (Swift %s, SDK %s)\n", info.Version, info.SwiftVersion, info.SDKVersion)

		fw := framework
		if fw == "" {
			fw = "swiftui"
		}
		minVer := minVersion
		if minVer == "" {
			if platform == "ios" {
				minVer = info.DefaultMinIOS
			} else {
				minVer = info.DefaultMinMacOS
			}
		}

		return ios.Generate(ios.Config{
			Name:       name,
			BundleID:   pkg,
			OutputDir:  out,
			Platform:   platform,
			Framework:  fw,
			MinVersion: minVer,
			XcodeInfo:  info,
		})

	case "android":
		info, err := detector.DetectAndroidStudio()
		if err != nil {
			return fmt.Errorf("Android Studio detection failed: %w", err)
		}
		fmt.Printf("Detected Android Studio %s (AGP %s, Gradle %s)\n", info.VersionName, info.AGPVersion, info.GradleVersion)

		fw := framework
		if fw == "" {
			fw = "compose"
		}

		return android.Generate(android.Config{
			Name:        name,
			PackageName: pkg,
			OutputDir:   out,
			Framework:   fw,
			MinVersion:  minVersion,
			StudioInfo:  info,
		})

	case "flutter":
		return flutter.Generate(flutter.Config{
			Name:      name,
			Org:       pkg,
			OutputDir: out,
		})

	default:
		return fmt.Errorf("unsupported platform: %s (supported: ios, macos, android, flutter)", platform)
	}
}

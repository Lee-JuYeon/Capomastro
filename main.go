package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cavss/Capomastro/internal/detector"
	"github.com/cavss/Capomastro/internal/generator/android"
	"github.com/cavss/Capomastro/internal/generator/api"
	"github.com/cavss/Capomastro/internal/generator/db"
	"github.com/cavss/Capomastro/internal/generator/flutter"
	"github.com/cavss/Capomastro/internal/generator/ios"
	"github.com/cavss/Capomastro/internal/generator/react"
	"github.com/cavss/Capomastro/internal/generator/server"
	"github.com/cavss/Capomastro/internal/generator/web"
)

var (
	platform   string
	pkg        string
	name       string
	out        string
	framework  string
	minVersion string
	coreData   bool

	// iOS/macOS signing
	team string

	// Android-specific overrides
	targetSdk     string
	javaVersion   string
	agpVersion    string
	gradleVersion string
	kotlinVersion string
)

func main() {
	root := &cobra.Command{
		Use:   "ProjectBuilder",
		Short: "Generate buildable projects for any platform",
		Long: `CLI tool that generates complete, buildable project scaffolds.

Platforms:
  ios      — Xcode project (SwiftUI default, --framework uikit)
  macos    — Xcode macOS project (SwiftUI default)
  android  — Android Studio project (Compose default, --framework xml)
  flutter  — Flutter project
  react    — React app (Vite+TS default, --framework nextjs)
  web      — Next.js app (Next.js+Tailwind default)
  api      — REST API server (Express+TS default, --framework hono/fastify)
  server   — Realtime server (WebSocket+Supabase default)
  db       — Database project (Supabase default, --framework prisma)`,
		RunE: run,
	}

	// Common
	root.Flags().StringVar(&platform, "platform", "", "Target platform (required)")
	root.Flags().StringVar(&pkg, "pkg", "", "Package/bundle identifier (e.g. com.example.app)")
	root.Flags().StringVar(&name, "name", "", "Project name (required)")
	root.Flags().StringVar(&out, "out", ".", "Output directory")
	root.Flags().StringVar(&framework, "framework", "", "UI/tech framework override (platform-specific)")
	root.Flags().StringVar(&minVersion, "min-version", "", "Minimum deployment version override")

	// iOS/macOS
	root.Flags().StringVar(&team, "team", "", "Apple Development Team ID")
	root.Flags().BoolVar(&coreData, "coredata", false, "Include Core Data model (iOS/macOS only)")

	// Android
	root.Flags().StringVar(&targetSdk, "target-sdk", "", "Android targetSdk/compileSdk (default: 35)")
	root.Flags().StringVar(&javaVersion, "java-version", "", "Java compatibility version: 11, 17, 21 (default: 17)")
	root.Flags().StringVar(&agpVersion, "agp-version", "", "Android Gradle Plugin version override")
	root.Flags().StringVar(&gradleVersion, "gradle-version", "", "Gradle version override")
	root.Flags().StringVar(&kotlinVersion, "kotlin-version", "", "Kotlin version override")

	root.MarkFlagRequired("platform")
	root.MarkFlagRequired("name")

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	switch platform {

	// ── iOS / macOS ──────────────────────────────────────────
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
			CoreData:   coreData,
			Team:       team,
			XcodeInfo:  info,
		})

	// ── Android ──────────────────────────────────────────────
	case "android":
		info, err := detector.DetectAndroidStudio()
		if err != nil {
			return fmt.Errorf("Android Studio detection failed: %w", err)
		}
		if agpVersion != "" {
			info.AGPVersion = agpVersion
		}
		if gradleVersion != "" {
			info.GradleVersion = gradleVersion
		}
		if kotlinVersion != "" {
			info.KotlinVersion = kotlinVersion
		}
		fmt.Printf("Detected Android Studio %s (AGP %s, Gradle %s, Kotlin %s)\n",
			info.VersionName, info.AGPVersion, info.GradleVersion, info.KotlinVersion)

		fw := framework
		if fw == "" {
			fw = "compose" // default
		}
		return android.Generate(android.Config{
			Name:        name,
			PackageName: pkg,
			OutputDir:   out,
			Framework:   fw,
			MinVersion:  minVersion,
			TargetSdk:   targetSdk,
			JavaVersion: javaVersion,
			StudioInfo:  info,
		})

	// ── Flutter ──────────────────────────────────────────────
	case "flutter":
		return flutter.Generate(flutter.Config{
			Name:      name,
			Org:       pkg,
			OutputDir: out,
		})

	// ── React ─────────────────────────────────────────────────
	// default: Vite + React + TS
	// --framework nextjs → Next.js (web generator 위임)
	case "react":
		if framework == "nextjs" || framework == "next" {
			return web.Generate(web.Config{Name: name, OutputDir: out})
		}
		return react.Generate(react.Config{Name: name, OutputDir: out})

	// ── Web (Next.js) ─────────────────────────────────────────
	// default: Next.js + Tailwind
	// --framework react → Vite React (react generator 위임)
	case "web":
		if framework == "react" || framework == "vite" {
			return react.Generate(react.Config{Name: name, OutputDir: out})
		}
		return web.Generate(web.Config{Name: name, OutputDir: out})

	// ── API (REST) ────────────────────────────────────────────
	// default: Express + TS
	// --framework hono / fastify → 향후 확장
	case "api":
		return api.Generate(api.Config{Name: name, OutputDir: out})

	// ── Server (Realtime) ─────────────────────────────────────
	// default: WebSocket + Supabase Realtime
	case "server":
		return server.Generate(server.Config{Name: name, OutputDir: out})

	// ── DB ────────────────────────────────────────────────────
	// default: Supabase migrations
	// --framework prisma → 향후 확장
	case "db":
		return db.Generate(db.Config{Name: name, OutputDir: out})

	default:
		return fmt.Errorf("unsupported platform: %q\nSupported: ios, macos, android, flutter, react, web, api, server, db", platform)
	}
}

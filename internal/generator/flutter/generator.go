package flutter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Name      string
	Org       string // e.g. "com.example"
	OutputDir string
}

func Generate(cfg Config) error {
	projectName := strings.ToLower(cfg.Name)
	projectRoot := filepath.Join(cfg.OutputDir, cfg.Name)

	// Directories
	lib := filepath.Join(projectRoot, "lib")
	test := filepath.Join(projectRoot, "test")
	androidApp := filepath.Join(projectRoot, "android", "app", "src", "main")
	androidJava := filepath.Join(androidApp, "java", strings.ReplaceAll(cfg.Org+"."+projectName, ".", "/"))
	androidRes := filepath.Join(androidApp, "res")
	androidDrawable := filepath.Join(androidRes, "drawable")
	androidDrawableV24 := filepath.Join(androidRes, "drawable-v24")
	androidMipmap := filepath.Join(androidRes, "mipmap-anydpi-v26")
	androidValues := filepath.Join(androidRes, "values")
	androidGradle := filepath.Join(projectRoot, "android", "gradle", "wrapper")
	iosRunner := filepath.Join(projectRoot, "ios", "Runner")
	iosAssets := filepath.Join(iosRunner, "Assets.xcassets", "AppIcon.appiconset")
	iosXcodeproj := filepath.Join(projectRoot, "ios", "Runner.xcodeproj")

	dirs := []string{
		lib, test,
		androidJava, androidDrawable, androidDrawableV24, androidMipmap, androidValues, androidGradle,
		iosRunner, iosAssets, iosXcodeproj,
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", d, err)
		}
	}

	bundleID := cfg.Org + "." + projectName

	files := map[string]string{
		// Root
		filepath.Join(projectRoot, "pubspec.yaml"):        pubspecYaml(projectName),
		filepath.Join(projectRoot, "analysis_options.yaml"): analysisOptions(),
		filepath.Join(projectRoot, ".gitignore"):           flutterGitignore(),

		// Dart
		filepath.Join(lib, "main.dart"):                   mainDart(),
		filepath.Join(test, "widget_test.dart"):            widgetTestDart(projectName),

		// Android
		filepath.Join(projectRoot, "android", "build.gradle.kts"):            androidRootBuildGradle(),
		filepath.Join(projectRoot, "android", "settings.gradle.kts"):         androidSettingsGradle(projectName),
		filepath.Join(projectRoot, "android", "gradle.properties"):           androidGradleProperties(),
		filepath.Join(androidGradle, "gradle-wrapper.properties"):            androidWrapperProperties(),
		filepath.Join(projectRoot, "android", "app", "build.gradle.kts"):     androidAppBuildGradle(bundleID),
		filepath.Join(androidApp, "AndroidManifest.xml"):                     androidManifest(projectName),
		filepath.Join(androidJava, "MainActivity.kt"):                        androidMainActivity(bundleID),
		filepath.Join(androidDrawable, "launch_background.xml"):              launchBackground(),
		filepath.Join(androidValues, "styles.xml"):                           androidStyles(projectName),
		filepath.Join(androidRes, "values-night", "styles.xml"):              androidStylesNight(projectName),

		// iOS
		filepath.Join(iosRunner, "AppDelegate.swift"):     iosAppDelegate(),
		filepath.Join(iosRunner, "Info.plist"):             iosInfoPlist(cfg.Name),
		filepath.Join(iosAssets, "Contents.json"):          iosAppIconContents(),
	}

	// Write all files
	for path, content := range files {
		dir := filepath.Dir(path)
		os.MkdirAll(dir, 0755)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", path, err)
		}
	}

	fmt.Printf("Flutter project created: %s\n", projectRoot)
	fmt.Printf("  package: %s\n", bundleID)
	fmt.Printf("  Run: cd %s && flutter pub get\n", projectRoot)
	return nil
}

// --- Root files ---

func pubspecYaml(name string) string {
	return fmt.Sprintf(`name: %s
description: A new Flutter project.
publish_to: 'none'
version: 1.0.0+1

environment:
  sdk: ^3.5.0

dependencies:
  flutter:
    sdk: flutter
  cupertino_icons: ^1.0.8

dev_dependencies:
  flutter_test:
    sdk: flutter
  flutter_lints: ^4.0.0

flutter:
  uses-material-design: true
`, name)
}

func analysisOptions() string {
	return `include: package:flutter_lints/flutter.yaml

linter:
  rules:
`
}

func flutterGitignore() string {
	return `.dart_tool/
.packages
build/
.flutter-plugins
.flutter-plugins-dependencies
*.iml
.idea/
.DS_Store
pubspec.lock
`
}

// --- Dart ---

func mainDart() string {
	return `import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorSchemeSeed: Colors.deepPurple,
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});
  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text('You have pushed the button this many times:'),
            Text(
              '$_counter',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}
`
}

func widgetTestDart(name string) string {
	return fmt.Sprintf(`import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:%s/main.dart';

void main() {
  testWidgets('Counter increments smoke test', (WidgetTester tester) async {
    await tester.pumpWidget(const MyApp());
    expect(find.text('0'), findsOneWidget);
    expect(find.text('1'), findsNothing);
    await tester.tap(find.byIcon(Icons.add));
    await tester.pump();
    expect(find.text('0'), findsNothing);
    expect(find.text('1'), findsOneWidget);
  });
}
`, name)
}

// --- Android ---

func androidRootBuildGradle() string {
	return `allprojects {
    repositories {
        google()
        mavenCentral()
    }
}

val newBuildDir: Directory = rootProject.layout.buildDirectory.dir("../../build").get()
rootProject.layout.buildDirectory.value(newBuildDir)

subprojects {
    val newSubprojectBuildDir: Directory = newBuildDir.dir(project.name)
    project.layout.buildDirectory.value(newSubprojectBuildDir)
}
subprojects {
    project.evaluationDependsOn(":app")
}

tasks.register<Delete>("clean") {
    delete(rootProject.layout.buildDirectory)
}
`
}

func androidSettingsGradle(name string) string {
	return fmt.Sprintf(`pluginManagement {
    val flutterSdkPath = run {
        val properties = java.util.Properties()
        file("local.properties").inputStream().use { properties.load(it) }
        val flutterSdkPath = properties.getProperty("flutter.sdk")
        require(flutterSdkPath != null) { "flutter.sdk not set in local.properties" }
        flutterSdkPath
    }

    includeBuild("$flutterSdkPath/packages/flutter_tools/gradle")

    repositories {
        google()
        mavenCentral()
        gradlePluginPortal()
    }
}

plugins {
    id("dev.flutter.flutter-plugin-loader") version "1.0.0"
    id("com.android.application") version "8.7.0" apply false
    id("org.jetbrains.kotlin.android") version "2.0.21" apply false
}

include(":app")
`)
}

func androidGradleProperties() string {
	return `org.gradle.jvmargs=-Xmx4G -XX:+HeapDumpOnOutOfMemoryError
android.useAndroidX=true
android.enableJetifier=true
`
}

func androidWrapperProperties() string {
	return `distributionBase=GRADLE_USER_HOME
distributionPath=wrapper/dists
distributionUrl=https\://services.gradle.org/distributions/gradle-8.10.2-bin.zip
zipStoreBase=GRADLE_USER_HOME
zipStorePath=wrapper/dists
`
}

func androidAppBuildGradle(bundleID string) string {
	return fmt.Sprintf(`plugins {
    id("com.android.application")
    id("kotlin-android")
    id("dev.flutter.flutter-gradle-plugin")
}

android {
    namespace = "%s"
    compileSdk = flutter.compileSdkVersion
    ndkVersion = flutter.ndkVersion

    compileOptions {
        sourceCompatibility = JavaVersion.VERSION_11
        targetCompatibility = JavaVersion.VERSION_11
    }

    kotlinOptions {
        jvmTarget = JavaVersion.VERSION_11.toString()
    }

    defaultConfig {
        applicationId = "%s"
        minSdk = flutter.minSdkVersion
        targetSdk = flutter.targetSdkVersion
        versionCode = flutter.versionCode
        versionName = flutter.versionName
    }

    buildTypes {
        release {
            signingConfig = signingConfigs.getByName("debug")
        }
    }
}

flutter {
    source = "../.."
}
`, bundleID, bundleID)
}

func androidManifest(name string) string {
	return fmt.Sprintf(`<manifest xmlns:android="http://schemas.android.com/apk/res/android">
    <application
        android:label="%s"
        android:name="${applicationName}"
        android:icon="@mipmap/ic_launcher">
        <activity
            android:name=".MainActivity"
            android:exported="true"
            android:launchMode="singleTop"
            android:taskAffinity=""
            android:theme="@style/LaunchTheme"
            android:configChanges="orientation|keyboardHidden|keyboard|screenSize|smallestScreenSize|locale|layoutDirection|fontScale|screenLayout|density|uiMode"
            android:hardwareAccelerated="true"
            android:windowSoftInputMode="adjustResize">
            <meta-data
                android:name="io.flutter.embedding.android.NormalTheme"
                android:resource="@style/NormalTheme" />
            <intent-filter>
                <action android:name="android.intent.action.MAIN"/>
                <category android:name="android.intent.category.LAUNCHER"/>
            </intent-filter>
        </activity>
        <meta-data
            android:name="flutterEmbedding"
            android:value="2" />
    </application>
    <queries>
        <intent>
            <action android:name="android.intent.action.PROCESS_TEXT"/>
            <data android:mimeType="text/plain"/>
        </intent>
    </queries>
</manifest>
`, name)
}

func androidMainActivity(bundleID string) string {
	return fmt.Sprintf(`package %s

import io.flutter.embedding.android.FlutterActivity

class MainActivity: FlutterActivity()
`, bundleID)
}

func launchBackground() string {
	return `<?xml version="1.0" encoding="utf-8"?>
<layer-list xmlns:android="http://schemas.android.com/apk/res/android">
    <item android:drawable="?android:colorBackground" />
</layer-list>
`
}

func androidStyles(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <style name="LaunchTheme" parent="@android:style/Theme.Light.NoTitleBar">
        <item name="android:windowBackground">@drawable/launch_background</item>
    </style>
    <style name="NormalTheme" parent="@android:style/Theme.Light.NoTitleBar">
        <item name="android:windowBackground">?android:colorBackground</item>
    </style>
</resources>
`)
}

func androidStylesNight(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <style name="LaunchTheme" parent="@android:style/Theme.Black.NoTitleBar">
        <item name="android:windowBackground">@drawable/launch_background</item>
    </style>
    <style name="NormalTheme" parent="@android:style/Theme.Black.NoTitleBar">
        <item name="android:windowBackground">?android:colorBackground</item>
    </style>
</resources>
`)
}

// --- iOS ---

func iosAppDelegate() string {
	return `import Flutter
import UIKit

@main
@objc class AppDelegate: FlutterAppDelegate {
    override func application(
        _ application: UIApplication,
        didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?
    ) -> Bool {
        GeneratedPluginRegistrant.register(with: self)
        return super.application(application, didFinishLaunchingWithOptions: launchOptions)
    }
}
`
}

func iosInfoPlist(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>$(DEVELOPMENT_LANGUAGE)</string>
	<key>CFBundleDisplayName</key>
	<string>%s</string>
	<key>CFBundleExecutable</key>
	<string>$(EXECUTABLE_NAME)</string>
	<key>CFBundleIdentifier</key>
	<string>$(PRODUCT_BUNDLE_IDENTIFIER)</string>
	<key>CFBundleName</key>
	<string>%s</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>$(FLUTTER_BUILD_NAME)</string>
	<key>CFBundleVersion</key>
	<string>$(FLUTTER_BUILD_NUMBER)</string>
	<key>LSRequiresIPhoneOS</key>
	<true/>
	<key>UILaunchStoryboardName</key>
	<string>LaunchScreen</string>
	<key>UIMainStoryboardFile</key>
	<string>Main</string>
	<key>UISupportedInterfaceOrientations</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
</dict>
</plist>
`, name, name)
}

func iosAppIconContents() string {
	return `{
  "images" : [
    {
      "idiom" : "universal",
      "platform" : "ios",
      "size" : "1024x1024"
    }
  ],
  "info" : {
    "author" : "xcode",
    "version" : 1
  }
}
`
}

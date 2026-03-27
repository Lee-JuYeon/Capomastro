package android

import (
	"fmt"
	"strings"
)

func rootBuildGradle(agpVersion, kotlinVersion string) string {
	return fmt.Sprintf(`plugins {
    id("com.android.application") version "%s" apply false
    id("org.jetbrains.kotlin.android") version "%s" apply false
}
`, agpVersion, kotlinVersion)
}

func settingsGradle(name string) string {
	return fmt.Sprintf(`pluginManagement {
    repositories {
        google()
        mavenCentral()
        gradlePlugin()
    }
}

dependencyResolutionManagement {
    repositoriesMode.set(RepositoriesMode.FAIL_ON_PROJECT_REPOS)
    repositories {
        google()
        mavenCentral()
    }
}

rootProject.name = "%s"
include(":app")
`, name)
}

func javaVersionConst(ver string) string {
	switch ver {
	case "11":
		return "VERSION_11"
	case "21":
		return "VERSION_21"
	default:
		return "VERSION_17"
	}
}

func appBuildGradleCompose(pkg, agpVersion, kotlinVersion, minSdk, targetSdk, compileSdk, javaVer string) string {
	return fmt.Sprintf(`plugins {
    id("com.android.application")
    id("org.jetbrains.kotlin.android")
}

android {
    namespace = "%s"
    compileSdk = %s

    defaultConfig {
        applicationId = "%s"
        minSdk = %s
        targetSdk = %s
        versionCode = 1
        versionName = "1.0"
    }

    buildTypes {
        release {
            isMinifyEnabled = false
            proguardFiles(getDefaultProguardFile("proguard-android-optimize.txt"), "proguard-rules.pro")
        }
    }

    compileOptions {
        sourceCompatibility = JavaVersion.%s
        targetCompatibility = JavaVersion.%s
    }

    kotlinOptions {
        jvmTarget = "%s"
    }

    buildFeatures {
        compose = true
    }

    composeOptions {
        kotlinCompilerExtensionVersion = "1.5.14"
    }
}

dependencies {
    implementation(platform("androidx.compose:compose-bom:2024.09.00"))
    implementation("androidx.compose.ui:ui")
    implementation("androidx.compose.ui:ui-graphics")
    implementation("androidx.compose.ui:ui-tooling-preview")
    implementation("androidx.compose.material3:material3")
    implementation("androidx.activity:activity-compose:1.9.0")
    implementation("androidx.core:core-ktx:1.13.1")
    debugImplementation("androidx.compose.ui:ui-tooling")
}
`, pkg, compileSdk, pkg, minSdk, targetSdk, javaVersionConst(javaVer), javaVersionConst(javaVer), javaVer)
}

func appBuildGradleXML(pkg, agpVersion, kotlinVersion, minSdk, targetSdk, compileSdk, javaVer string) string {
	return fmt.Sprintf(`plugins {
    id("com.android.application")
    id("org.jetbrains.kotlin.android")
}

android {
    namespace = "%s"
    compileSdk = %s

    defaultConfig {
        applicationId = "%s"
        minSdk = %s
        targetSdk = %s
        versionCode = 1
        versionName = "1.0"
    }

    buildTypes {
        release {
            isMinifyEnabled = false
            proguardFiles(getDefaultProguardFile("proguard-android-optimize.txt"), "proguard-rules.pro")
        }
    }

    compileOptions {
        sourceCompatibility = JavaVersion.%s
        targetCompatibility = JavaVersion.%s
    }

    kotlinOptions {
        jvmTarget = "%s"
    }

    buildFeatures {
        viewBinding = true
    }
}

dependencies {
    implementation("androidx.core:core-ktx:1.13.1")
    implementation("androidx.appcompat:appcompat:1.7.0")
    implementation("com.google.android.material:material:1.12.0")
    implementation("androidx.constraintlayout:constraintlayout:2.1.4")
}
`, pkg, compileSdk, pkg, minSdk, targetSdk, javaVersionConst(javaVer), javaVersionConst(javaVer), javaVer)
}

func gradleProperties() string {
	return `org.gradle.jvmargs=-Xmx2048m -Dfile.encoding=UTF-8
android.useAndroidX=true
kotlin.code.style=official
android.nonTransitiveRClass=true
`
}

func gradleWrapperProperties(gradleVersion string) string {
	return fmt.Sprintf(`distributionBase=GRADLE_USER_HOME
distributionPath=wrapper/dists
distributionUrl=https\://services.gradle.org/distributions/gradle-%s-bin.zip
zipStoreBase=GRADLE_USER_HOME
zipStorePath=wrapper/dists
`, gradleVersion)
}

func androidManifest(pkg string, framework string) string {
	activity := ""
	if framework == "compose" {
		activity = fmt.Sprintf(`        <activity
            android:name=".MainActivity"
            android:exported="true"
            android:theme="@android:style/Theme.Material.Light.NoActionBar">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>`)
	} else {
		activity = fmt.Sprintf(`        <activity
            android:name=".MainActivity"
            android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>`)
	}

	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android">
    <application
        android:allowBackup="true"
        android:icon="@mipmap/ic_launcher"
        android:label="@string/app_name"
        android:supportsRtl="true"
        android:theme="@style/Theme.Material3.DayNight.NoActionBar">
%s
    </application>
</manifest>
`, activity)
}

func mainActivityCompose(pkg string) string {
	return fmt.Sprintf(`package %s

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            MaterialTheme {
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    Greeting()
                }
            }
        }
    }
}

@Composable
fun Greeting() {
    Column(
        modifier = Modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Text(
            text = "Hello, world!",
            style = MaterialTheme.typography.headlineMedium
        )
    }
}
`, pkg)
}

func mainActivityXML(pkg string) string {
	return fmt.Sprintf(`package %s

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }
}
`, pkg)
}

func activityMainXML() string {
	return `<?xml version="1.0" encoding="utf-8"?>
<androidx.constraintlayout.widget.ConstraintLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="Hello, world!"
        android:textSize="24sp"
        app:layout_constraintBottom_toBottomOf="parent"
        app:layout_constraintEnd_toEndOf="parent"
        app:layout_constraintStart_toStartOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

</androidx.constraintlayout.widget.ConstraintLayout>
`
}

func stringsXML(name string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="app_name">%s</string>
</resources>
`, name)
}

func proguardRules() string {
	return "# Add project specific ProGuard rules here.\n"
}

func pkgToPath(pkg string) string {
	return strings.ReplaceAll(pkg, ".", "/")
}

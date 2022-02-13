package androidbinary_test

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/zinego/androidbinary"
	"github.com/zinego/androidbinary/apk"
)

func ExampleNewXMLFile() {
	f, _ := os.Open("testdata/AndroidManifest.xml")
	xmlFile, err := androidbinary.NewXMLFile(f)
	if err != nil {
		panic(err)
	}

	var v apk.Manifest
	dec := xml.NewDecoder(xmlFile.Reader())
	dec.Decode(&v)

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	enc.Encode(v)

	// Output:
	// 	<Manifest package="net.sorablue.shogo.FWMeasure" xmlns:android="http://schemas.android.com/apk/res/android" android:versionCode="1" android:versionName="テスト版">
	// 	<application android:allowTaskReparenting="false" android:allowBackup="false" android:backupAgent="" android:debuggable="false" android:description="" android:enabled="false" android:hasCode="false" android:hardwareAccelerated="false" android:icon="@0x7F020000" android:killAfterRestore="false" android:largeHeap="false" android:label="@0x7F040000" android:logo="" android:manageSpaceActivity="" android:name="" android:permission="" android:persistent="false" android:process="" android:restoreAnyVersion="false" android:requiredAccountType="" android:restrictedAccountType="" android:supportsRtl="false" android:taskAffinity="" android:testOnly="false" android:theme="" android:uiOptions="" android:vmSafeMode="false">
	// 		<activity android:theme="" android:name="FWMeasureActivity" android:label="" android:screenOrientation="0">
	// 			<intent-filter>
	// 				<action android:name="android.intent.action.MAIN"></action>
	// 				<category android:name="android.intent.category.LAUNCHER"></category>
	// 			</intent-filter>
	// 		</activity>
	// 		<activity android:theme="" android:name="MapActivity" android:label="" android:screenOrientation="0"></activity>
	// 		<activity android:theme="" android:name="SettingActivity" android:label="" android:screenOrientation=""></activity>
	// 		<activity android:theme="" android:name="PlaceSettingActivity" android:label="" android:screenOrientation=""></activity>
	// 	</application>
	// 	<instrumentation android:name="" android:targetPackage="" android:handleProfiling="false" android:functionalTest="false"></instrumentation>
	// 	<uses-sdk android:minSdkVersion="0" android:targetSdkVersion="0" android:maxSdkVersion="0"></uses-sdk>
	// 	<uses-permission android:name="android.permission.CAMERA" android:maxSdkVersion="0"></uses-permission>
	// 	<uses-permission android:name="android.permission.WAKE_LOCK" android:maxSdkVersion="0"></uses-permission>
	// 	<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" android:maxSdkVersion="0"></uses-permission>
	// 	<uses-permission android:name="android.permission.INTERNET" android:maxSdkVersion="0"></uses-permission>
	// 	<uses-permission android:name="android.permission.ACCESS_MOCK_LOCATION" android:maxSdkVersion="0"></uses-permission>
	// 	<uses-permission android:name="android.permission.RECORD_AUDIO" android:maxSdkVersion="0"></uses-permission>
	// </Manifest>
}

func ExampleNewTableFile() {
	f, err := os.Open("testdata/resources.arsc")
	if err != nil {
		panic(err)
	}
	tableFile, err := androidbinary.NewTableFile(f)
	if err != nil {
		panic(err)
	}

	val, err := tableFile.GetBestResource(0x7f040000, &androidbinary.ResTableConfig{})
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	// Output:
	// FireworksMeasure
}

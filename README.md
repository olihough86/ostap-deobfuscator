# ostap-deobfuscator
Tool to deobfuscate ostap JSE

## Usage
copy the raw JSE from the document and save it as plain text  
*go run main.go <file-with-jse.txt>*

## Example
PS C:\Users\oliho\go\src\ostap-deobfuscator> go run .\main.go .\mal.txt
```
Tool to deobfuscate ostap JSE - Oliver Hough (@olihough86)
Found possible array identifier: Medaxc
MS Word:This file is in use by another application or user. fromCharCodeWScriptActiveXObjectScriptFullNameCreateObjectWScript.ShelltoLowerCaseindexOftempPopup2090000FiScripting.FileSystemObjectfromCharCodeDrives*.odt *.ods *.odp *.odm *.odc *.odb *.wps *.xlk *.ppt *.pst *.dwg *.dxf *.dxg *.wpd *.doc *.xls *.pdf *.rtf *.txt *.pubbirdy.txt4294967295GetObjectEnumeratorEnvironmentPROCESSItemCOMPUTERNAMEOpenTextFileReadAllClosevar vedoifloorrandom=floorrandom;winmgmts:{impersonationLevel=impersonate}!.rootcimv2ExecQuerySelect * from Win32_ProcessatEnditemName*ExecutablePathfromCharCodefromCharCodemoveNextlengthfaztaxShell.ApplicationExpandEnvironmentStrings%TEMP%lengthcharCodeAtabs.humabs.tumMZMicrosoft.XMLDOMcreateElementbase64ADODB.StreamMSXML2.XMLHTTP.6.0http://185.195.237.14/21RWq/jucE4.php?zs=s20&ed=absGETaKamaZiindexOfOpenTextFileReadAllCloseSleepDeleteFiledataTypebin.base64textsplitjoinOpenTypePositionWritenodeTypedValuereplaceslice.exeSaveToFileCloseShellExecuteopenShellExecutecmd/copenSleepDeleteFileatEndmoveNextitemIsReadyDriveTypeDriveTypesubstringDriveLetterShellExecutecmd/U /Q /C cd /D DriveLetter: && dir /b/s/x >>%TEMP%openSleepSleepGetFileOpenAsTextStreamAtEndOfStreamReadLinesubstringindexOf.ShellExecutecmd/U /Q /C copy /Y  .jse && del /Q/F openCloseDeleteFileSleepDeleteFileDeleteFile&rnx=floorrandomfloorrandomOpenfloorrandomSetRequestHeaderUser-AgentMozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79..3945.88 Safari/537.36SendstatusresponseTextindexOfCreateTextFileWriteCloseCreateTextFileWriteCloseShellExecutewscript/B /E:JScript openShellExecutecscript/B /E:JScript openSleepSleepSleepSleepSleep
```
## Note
The math performed while calculating the array indexes can change which will cause the script to fail,  
but the over all priciple remains the same. A later release will attempt to detect this.

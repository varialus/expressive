// export GOPATH=$HOME/go
// export PATH=$PATH:$GOPATH/bin
// cd ~/go/src/
// go test GolangStandardLibrary/...

package format

import (
	"testing"
	//"io"
	"io/ioutil"
	"fmt"
	"os"
)

// Smoke Test
func TestErrorFormatted(testStateManager *testing.T) {
	if errorValue := ErrorFormatted("%d decimal == %b binary\n", 42, 42); errorValue == nil {
		testStateManager.Errorf("Failed TestErrorFormatted")
	}
}

// Smoke Test
func TestFilePrint(testStateManager *testing.T) {
	if numberOfBytesWritten, errorValue := FilePrint(ioutil.Discard, "42", 42); (numberOfBytesWritten != 4) || (errorValue != nil) {
		testStateManager.Errorf("Failed TestFilePrint")
	}
}

// Smoke Test
func TestFilePrintFormatted(testStateManager *testing.T){
	if numberOfBytesWritten, errorValue := FilePrintFormatted(ioutil.Discard, "%d decimal == %b binary\n", 42, 42); (numberOfBytesWritten != 28) || (errorValue != nil) {
		testStateManager.Errorf("Failed TestFilePrintFormatted")
	}
}

// Smoke Test
func TestFilePrintLine(testStateManager *testing.T){
	if numberOfBytesWritten, errorValue := FilePrintLine(ioutil.Discard, "42", 42); (numberOfBytesWritten != 6) || (errorValue != nil) {
		Print(numberOfBytesWritten)
		testStateManager.Errorf("Failed TestFilePrintLine")
	}
}

// Smoke Test
func TestFileScan(testStateManager *testing.T){
	// Returns Opened File
	tempFile, errorValue := ioutil.TempFile("", "TestFileScan")
	if errorValue != nil {
		testStateManager.Errorf("Failed TestFileScan Setup")
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	tempFile.WriteString("4 2 42")
	tempFile.Seek(0, 0)
	var (
	firstNumberScanned int
	secondNumberScanned int
	thirdNumberScanned int
	)
	if numberOfItemsScanned, errorValue := fmt.Fscan(tempFile, &firstNumberScanned, &secondNumberScanned, &thirdNumberScanned); (numberOfItemsScanned != 3) || (errorValue != nil){
		testStateManager.Errorf("Failed TestFileScan")
	}
}
/*
// Smoke Test
func TestFileScanFormatted(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := FileScanFormatted(reader, unescapedFormatSpecifiers, arguments...); {
		testStateManager.Errorf("Failed TestFileScanFormatted")
	}
}

// Smoke Test
func TestFileScanLine(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := FileScanLine(reader, arguments...); {
		testStateManager.Errorf("Failed TestFileScanLine")
	}
}

// Smoke Test
func TestPrint(testStateManager *testing.T){
	if numberOfBytesWritten, errorValue := Print(arguments...); {
		testStateManager.Errorf("Failed TestPrint")
	}
}

// Smoke Test
func TestPrintFormatted(testStateManager *testing.T) {
	if numberOfBytesPrinted, errorValue := PrintFormatted("%d decimal == %b binary\n", 42, 42); (numberOfBytesPrinted != 28) || (errorValue != nil) {
		testStateManager.Errorf("Failed TestPrintFormatted")
	}
}

// Smoke Test
func TestPrintLine(testStateManager *testing.T) {
	if numberOfBytesWritten, errorValue := PrintLine(arguments...); {
		testStateManager.Errorf("Failed TestPrintLine")
	}
}

// Smoke Test
func TestScan(testStateManager *testing.T) {
	if numberOfBytesRead, errorValue := Scan(arguments...); {
		testStateManager.Errorf("Failed TestScan")
	}
}

// Smoke Test
func TestScanFormatted(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := ScanFormatted(unescapedFormatSpecifiers, arguments...); {
		testStateManager.Errorf("Failed TestScanFormatted")
	}
}

// Smoke Test
func TestScanLine(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := ScanLine(arguments...); {
		testStateManager.Errorf("Failed TestScanLine")
	}
}

// Smoke Test
func TestStringPrint(testStateManager *testing.T){
	if argumentsAsString := StringPrint(arguments...); {
		testStateManager.Errorf("Failed TestStringPrint")
	}
}

// Smoke Test
func TestStringPrintFormatted(testStateManager *testing.T){
	if argumentsAsString := StringPrintFormatted(unescapedFormatSpecifiers, arguments...); {
		testStateManager.Errorf("Failed TestStringPrintFormatted")
	}
}

// Smoke Test
func TestStringPrintLine(testStateManager *testing.T){
	if argumentsAsString := StringPrintLine(arguments...); {
		testStateManager.Errorf("Failed TestStringPrintLine")
	}
}

// Smoke Test
func TestStringScan(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := StringScan(stringToScan, arguments...); {
		testStateManager.Errorf("Failed TestStringScan")
	}
}

// Smoke Test
func TestStringScanFormatted(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := StringScanFormatted(stringToScan, unescapedFormatSpecifiers, arguments...); {
		testStateManager.Errorf("Failed TestStringScanFormatted")
	}
}

// Smoke Test
func TestStringScanLine(testStateManager *testing.T){
	if numberOfBytesRead, errorValue := StringScanLine(stringToScan, arguments...); {
		testStateManager.Errorf("Failed TestStringScanLine")
	}
}*/

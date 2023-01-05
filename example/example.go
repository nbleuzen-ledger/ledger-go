package main

// Sample code based on the APDU supported by the Boilerplate Nano app:
// https://github.com/LedgerHQ/app-boilerplate

import (
	"fmt"
	"os"

	"github.com/nbleuzen-ledger/ledger-go"
)

const (
	CLA                             = 0xE0
	insGetVersion                   = 0x03
	insGetAppName                   = 0x04
)

// VersionInfo contains app version information
type VersionInfo struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func (version VersionInfo) String() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}

// LedgerApp represents a connection to an app on a Ledger device
type LedgerApp struct {
	api     ledger_go.LedgerDevice
	name    string
	version VersionInfo
}

func main() {
	if err := exampleSample(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exampleSample() error {
	app := &LedgerApp{nil, "", VersionInfo{}}
	if err := app.Connect(); err != nil {
		return fmt.Errorf("Failed to connect to a Ledger device: %s", err)
	}
	defer app.Close()

	if err := app.GetAppName(); err != nil {
		if err.Error() == "[APDU_CODE_CLA_NOT_SUPPORTED] Class not supported" {
			err = fmt.Errorf("%s -> Are you sure the correct app is open?", err)
		}
		return fmt.Errorf("Failed to get app name: %s", err)
	}
	fmt.Printf("App name is: %s\n", app.name)

	if err := app.GetVersion(); err != nil {
		return fmt.Errorf("Failed to get app version: %s", err)
	}
	fmt.Printf("App version is: %s\n", app.version)

	return nil
}

// Creates a connection with the app
func (ledger *LedgerApp) Connect() error {
	ledgerAdmin := ledger_go.NewLedgerAdmin()

	if deviceCount := ledgerAdmin.CountDevices(); deviceCount == 0 {
		return fmt.Errorf("No Ledger device connected")
	} else if deviceCount > 1 {
		fmt.Printf("%d Ledger devices connected, will try to connect to the " +
				"first one in the following list:\n", deviceCount)
		ledgerAdmin.ListDevices()
	}

	ledgerDevice, err := ledgerAdmin.Connect(0)
	if err != nil {
		fmt.Printf("Failed to connect to device: %s\n", err)
		return err
	}
	ledger.api = ledgerDevice

	return nil
}

// Closes a connection with the app
func (ledger *LedgerApp) Close() {
	ledger.api.Close()
}

// Sets the name of the app
func (ledger *LedgerApp) GetAppName() error {
	message := []byte{CLA, insGetAppName, 0 /*P1*/, 0 /*P2*/, 0 /*Data length*/}
	response, sw, err := ledger.api.ExchangeNoCheck(message)
	if err != nil {
		fmt.Printf("insGetAppName -> SW: %04x\n", sw)
		return err
	}

	ledger.name = string(response[:])
	return nil
}

// Sets the version of the app
func (ledger *LedgerApp) GetVersion() error {
	message := []byte{CLA, insGetVersion, 0 /*P1*/, 0 /*P2*/, 0 /*Data length*/}
	response, err := ledger.api.Exchange(message)
	if err != nil {
		return err
	}

	if len(response) < 3 {
		return fmt.Errorf("Invalid response length")
	}

	ledger.version = VersionInfo{
		Major: response[0],
		Minor: response[1],
		Patch: response[2],
	}
	return nil
}

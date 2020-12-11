package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kor44/extcap"

	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

// Define all options
var (
	SnapLength = extcap.NewConfigIntegerOpt("snap-len", "packet snapshot length")
)

func main() {
	app := extcap.App{
		Usage:               "sample extcap application",
		HelpPage:            "Sample application to show how to use 'extcap' package. Uses gopacket package for basic capturing",
		GetInterfaces:       getAllInterfaces,
		GetDLT:              getDLT,
		GetAllConfigOptions: getAllConfigOptions,
		GetConfigOptions:    getConfigOptions,
		StartCapture:        startCapture,
	}

	app.Run(os.Args)
}

func getAllInterfaces() ([]extcap.CaptureInterface, error) {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		err = fmt.Errorf("Unable to get information about interfaces: %w", err)
		return nil, err
	}

	extIfaces := make([]extcap.CaptureInterface, len(ifaces))
	for i, iface := range ifaces {
		extIfaces[i] = extcap.CaptureInterface{
			Value:   iface.Name,
			Display: "extcapdump: " + iface.Description,
		}
	}

	return extIfaces, nil
}

func getDLT(iface string) (extcap.DLT, error) {
	inactiveHandler, err := pcap.NewInactiveHandle(iface)
	if err != nil {
		err = fmt.Errorf("Open interface '%s' error: %w", iface, err)
		return extcap.DLT{}, err
	}
	defer inactiveHandler.CleanUp()

	// Activate capture
	var handle *pcap.Handle
	if handle, err = inactiveHandler.Activate(); err != nil {
		err = fmt.Errorf("Unable to get DLT error: %s", err)
		return extcap.DLT{}, err
	}
	defer handle.Close()

	linkType := handle.LinkType()

	dlt := extcap.DLT{
		Number: int(linkType.LayerType()),
		Name:   linkType.String(),
	}

	return dlt, nil
}

func startCapture(iface string, pipe io.WriteCloser, filter string, opts map[string]interface{}) error {
	defer pipe.Close()

	// file, err := os.Open(fifo)
	// if err != nil {
	// 	return err
	// }
	w := pcapgo.NewWriter(pipe)

	inactiveHandler, err := pcap.NewInactiveHandle(iface)
	if err != nil {
		err = fmt.Errorf("Open interface '%s' error: %w", iface, err)
		return err
	}
	defer inactiveHandler.CleanUp()

	// // snap length
	// if err = inactiveHandler.SetSnapLen(-1); err != nil {
	// fmt.Fprintf(os.Stderr, "Set snap length error: %s\n", err)
	// os.Exit(-1)
	// }
	// if err = inactiveHandler.SetTimeout(-1); err != nil {
	// fmt.Fprintf(os.Stderr, "Set timeout error: %s\n", err)
	// os.Exit(-1)
	// }

	// Activate capture
	var handle *pcap.Handle
	if handle, err = inactiveHandler.Activate(); err != nil {
		err = fmt.Errorf("Unable to get DLT error: %s", err)
		return err
	}
	defer handle.Close()

	if err := w.WriteFileHeader(0, handle.LinkType()); err != nil {
		return fmt.Errorf("Can't write pcap file header: %s", err)
	}

	for data, ci, err := handle.ZeroCopyReadPacketData(); ; data, ci, err = handle.ZeroCopyReadPacketData() {
		if err != nil {
			return fmt.Errorf("Read packet error: %w", err)
		}
		w.WritePacket(ci, data)
	}

	return nil
}

func getConfigOptions(iface string) ([]extcap.ConfigOption, error) {
	opts := []extcap.ConfigOption{
		SnapLength,
	}

	return opts, nil
}

func getAllConfigOptions() []extcap.ConfigOption {
	opts := []extcap.ConfigOption{
		SnapLength,
	}
	return opts
}

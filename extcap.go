/*
Package extcap implements library to help create cli app for capture by Wireshark using extcap interface (https://www.wireshark.org/docs/wsdg_html_chunked/ChCaptureExtcap.html).
For flags parsing, package urfave/cli is used (https://github.com/urfave/cli)

For minimal application should be implemented following functions:
	GetInterfaces:
	GetDLT:
	StartCapture:

Full working example can be found in examples folder

*/
package extcap

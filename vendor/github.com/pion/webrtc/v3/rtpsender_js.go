// +build js,wasm

package webrtc

import "syscall/js"

// RTPSender allows an application to control how a given Track is encoded and transmitted to a remote peer
type RTPSender struct {
	// Pointer to the underlying JavaScript RTCRTPSender object.
	underlying js.Value
}

package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	DataDrivenScreenCloseReasonProgrammaticClose = iota
	DataDrivenScreenCloseReasonProgrammaticCloseAll
	DataDrivenScreenCloseReasonClientCanceled
	DataDrivenScreenCloseReasonUserBusy
	DataDrivenScreenCloseReasonInvalidForm
)

// ServerBoundDataDrivenScreenClosed is sent by the client when a data-driven UI screen is closed.
type ServerBoundDataDrivenScreenClosed struct {
	// FormID is the optional unique instance ID of the form that was closed.
	FormID protocol.Optional[uint32]
	// CloseReason is the reason the screen was closed. It is one of the DataDrivenScreenCloseReason constants.
	CloseReason uint8
}

// ID ...
func (*ServerBoundDataDrivenScreenClosed) ID() uint32 {
	return IDServerBoundDataDrivenScreenClosed
}

func (pk *ServerBoundDataDrivenScreenClosed) Marshal(io protocol.IO) {
	closeReason := closeReasonToString(pk.CloseReason)
	protocol.OptionalFunc(io, &pk.FormID, io.Uint32)
	io.String(&closeReason)
	closeReasonFromString(io, &pk.CloseReason, closeReason)
}

// closeReasonFromString looks up an close reason from a string and writes the result to x.
func closeReasonFromString(io protocol.IO, x *uint8, s string) {
	switch s {
	case "programmaticclose":
		*x = DataDrivenScreenCloseReasonProgrammaticClose
	case "programmaticcloseall":
		*x = DataDrivenScreenCloseReasonProgrammaticCloseAll
	case "clientcanceled":
		*x = DataDrivenScreenCloseReasonClientCanceled
	case "userbusy":
		*x = DataDrivenScreenCloseReasonUserBusy
	case "invalidform":
		*x = DataDrivenScreenCloseReasonInvalidForm
	default:
		io.InvalidValue(s, "closeReason", "unknown close reason")
	}
}

// closeReasonToString looks up an close reason constant and returns the string representation.
func closeReasonToString(x uint8) string {
	switch x {
	case DataDrivenScreenCloseReasonProgrammaticClose:
		return "programmaticclose"
	case DataDrivenScreenCloseReasonProgrammaticCloseAll:
		return "programmaticcloseall"
	case DataDrivenScreenCloseReasonClientCanceled:
		return "clientcanceled"
	case DataDrivenScreenCloseReasonUserBusy:
		return "userbusy"
	case DataDrivenScreenCloseReasonInvalidForm:
		return "invalidform"
	default:
		return "unknown"
	}
}

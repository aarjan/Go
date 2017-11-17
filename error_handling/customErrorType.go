//using custom error types

//In most of the cases, traditional error interface value provided by the errors package is
//enough for reporting and handling errors. However, sometimes the caller needs extra context
//in order to make a more informed error handling decisions. This is when custom error type make sense.

//from net package
package main

type OpError struct {
	//read or write operation
	Op string

	Net string //network type; tcp/udp

	Addr Addr //network address on which error occurred

	Err error //type of error that occurred
}

//a sample function that implements the custom type error

func Listen(net, laddr string) (Listener, error) {
	la, err := resolveAddr("listen", net, laddr, noDeadLine)
	if err != nil {
		return nil, &OpError{Op: "listen", Net: net, Addr: nil, Err: err}
	}

	var l Listener
	switch la := la.toAddr().(type) {
	case *TCPAddr:
		l, err = ListenTCP(net, la)
	case *UnixAddr:
		l, err = ListenUnix(net, la)
	default:
		return nil, &OpError{Op: "listen", Net: net, Addr: la, Err: &AddrError{Err: "unexpected address"}}
	}

	if err != nil {
		return nil, err //l is a non-nil interface containing nil pointer, so it is not returned
	}
	return l, nil
}

//implementation of error interface for OpError struct
func (e *OpError) Error() string {
	if e == nil {
		return "<nil>"
	}
	s := e.Op
	if e.Net != nil {
		s += "" + e.Net
	}
	if e.Addr != nil {
		s += "" + e.Addr.string()
	}
	s += ":" + e.Err.Error()

	return s
}

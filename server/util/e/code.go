package e

type code struct {
	Success          int
	AuthEmpty        int
	AuthFormatError  int
	AuthTokenError   int
	AuthTokenTimeOut int
	AuthError        int
}

var CODE code

func init() {
	CODE = code{
		Success:          200,
		AuthEmpty:        401,
		AuthFormatError:  402,
		AuthTokenError:   403,
		AuthTokenTimeOut: 404,
		AuthError:        405,
	}
}

func ParseCode(num int) string {
	var msg = ""
	switch num {
	case CODE.Success:
		msg = "OK"
	case CODE.AuthEmpty:
		msg = "Auth Empty"
	case CODE.AuthFormatError:
		msg = "Auth Format Error"
	case CODE.AuthTokenError:
		msg = "Auth Token Error"
	case CODE.AuthTokenTimeOut:
		msg = "Auth Token Time Out"
	case CODE.AuthError:
		msg = "Auth Error: Wrong Info"
	default:
		msg = "Unknown Error"
	}
	return msg
}

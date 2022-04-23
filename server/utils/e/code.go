package e

type code struct {
	Success          int
	AuthEmpty        int
	AuthFormatError  int
	AuthTokenError   int
	AuthTokenTimeOut int
}

var CODE code

func init() {
	CODE = code{
		Success:          200,
		AuthEmpty:        401,
		AuthFormatError:  402,
		AuthTokenError:   403,
		AuthTokenTimeOut: 404,
	}
}

func ParseCode(num int) string {
	var msg = ""
	switch num {
	case CODE.Success:
		msg = "Success"
	case CODE.AuthEmpty:
		msg = "Auth Empty"
	case CODE.AuthFormatError:
		msg = "Auth Format Error"
	case CODE.AuthTokenError:
		msg = "Auth Token Error"
	case CODE.AuthTokenTimeOut:
		msg = "Auth Token Time Out"
	default:
		msg = "Unknown Error"
	}
	return msg
}

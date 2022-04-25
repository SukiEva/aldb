package e

type code struct {
	Success           int
	CaptchaError      int
	CaptchaMismatch   int
	AuthBindError     int
	AuthEmpty         int
	AuthFormatError   int
	AuthTokenError    int
	AuthTokenTimeOut  int
	AuthError         int
	UserBindError     int
	UserAlreadyExists int
}

var CODE code

func init() {
	CODE = code{
		Success:           200,
		CaptchaError:      301,
		CaptchaMismatch:   302,
		AuthBindError:     400,
		AuthEmpty:         401,
		AuthFormatError:   402,
		AuthTokenError:    403,
		AuthTokenTimeOut:  404,
		AuthError:         405,
		UserBindError:     500,
		UserAlreadyExists: 501,
	}
}

func ParseCode(num int) string {
	var msg = ""
	switch num {
	case CODE.Success:
		msg = "OK"
	case CODE.CaptchaError:
		msg = "Fail to get captcha"
	case CODE.CaptchaMismatch:
		msg = "Captcha Mismatch"
	case CODE.AuthBindError:
		msg = "Auth Bind Error"
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
	case CODE.UserBindError:
		msg = "User Bind Error"
	case CODE.UserAlreadyExists:
		msg = "User Already Exists"
	default:
		msg = "Unknown Error"
	}
	return msg
}

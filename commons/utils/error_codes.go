package utils

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

const (
	ErrorCodeSuccess ErrorCode = "200"

	// SMS service
	ErrorCodeSMSInvalidPhoneNumber ErrorCode = "SMS400001"
	ErrorCodeSMSEmptyContent       ErrorCode = "SMS400002"
	ErrorCodeSMSSendFail           ErrorCode = "SMS400003"
	ErrorCodeSMSTypeTooManyTimes   ErrorCode = "SMS400004"
	ErrorCodeSMSServiceUnavailable ErrorCode = "SMS500001"
)

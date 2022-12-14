package errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Module constants definition.
const (
	ModuleCommon = 00
	ModuleUser   = 01
	ModuleRole   = 02
)

// Common module error codes definition.
var (
	InvalidRequestError    = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 1)
	FileTemplateError      = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 3)
	RequestTimeoutError    = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 4)
	CustomMessageError     = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 5)
	UnsupportedEntityError = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 6)
	UnauthorizedCodeError  = fmtErrorCode(http.StatusUnauthorized, ModuleCommon, 1)
	HeaderInvalidError     = fmtErrorCode(http.StatusUnauthorized, ModuleCommon, 2)
	HeaderNotProvideError  = fmtErrorCode(http.StatusUnauthorized, ModuleCommon, 2)
	ForbiddenError         = fmtErrorCode(http.StatusForbidden, ModuleCommon, 1)

	InternalServerError      = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 1)
	InternalMissingMetaError = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 2)
	ErrNoResponse            = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 3)
)

// User
var (
	DuplicateError  = fmtErrorCode(http.StatusBadRequest, ModuleUser, 1)
	PasswordInvalid = fmtErrorCode(http.StatusBadRequest, ModuleUser, 2)
)
var (
	RoleIsExistedError = fmtErrorCode(http.StatusBadRequest, ModuleRole, 1)
)

type (
	translatedMessages struct {
		VI string `json:"vi"`
		EN string `json:"en"`
	}
	detailCodeMap map[string]translatedMessages
	statusMap     map[string]detailCodeMap
)

var errorMessageMap map[string]statusMap

func init() {
	errorMessageMap = make(map[string]statusMap)
}

// ErrorMessagesFilePath is path of error_messages.json file.
const ErrorMessagesFilePath = "./error_messages.json"

// InitErrorMessagesResource loads error messages resource.
func InitErrorMessagesResource() error {
	buf, err := ioutil.ReadFile(ErrorMessagesFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &errorMessageMap)
	if err != nil {
		return err
	}
	return nil
}

// GetErrorMessage gets error message from errorMessageMap.
func GetErrorMessage(errCode ErrorCode, args ...interface{}) string {
	if len(args) > 0 {
		msg, ok := args[0].(string)
		if ok {
			return msg
		}
	}
	msg := http.StatusText(errCode.Status())
	if errorMessageMap == nil {
		return msg
	}
	modules, ok := errorMessageMap[fmt.Sprintf("%02d", errCode.Module())]
	if modules == nil || !ok {
		return msg
	}
	statuses, ok := modules[fmt.Sprintf("%d", errCode.Status())]
	if statuses == nil || !ok {
		return msg
	}
	detailCodes, ok := statuses[fmt.Sprintf("%02d", errCode.DetailCode())]
	if !ok {
		return msg
	}
	if detailCodes.VI != "" {
		msg = detailCodes.EN
	}
	return fmt.Sprintf(msg, args...)
}

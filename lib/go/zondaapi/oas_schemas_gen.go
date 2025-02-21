// Code generated by ogen, DO NOT EDIT.

package zondaapi

type CheckHealthOK struct {
	Revision string `json:"revision"`
}

// GetRevision returns the value of Revision.
func (s *CheckHealthOK) GetRevision() string {
	return s.Revision
}

// SetRevision sets the value of Revision.
func (s *CheckHealthOK) SetRevision(val string) {
	s.Revision = val
}

// Ref: #/components/schemas/error
type Error struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() float64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val float64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

func (*ErrorStatusCode) signUpRes() {}

type SignUpOK struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GetAccessToken returns the value of AccessToken.
func (s *SignUpOK) GetAccessToken() string {
	return s.AccessToken
}

// GetRefreshToken returns the value of RefreshToken.
func (s *SignUpOK) GetRefreshToken() string {
	return s.RefreshToken
}

// SetAccessToken sets the value of AccessToken.
func (s *SignUpOK) SetAccessToken(val string) {
	s.AccessToken = val
}

// SetRefreshToken sets the value of RefreshToken.
func (s *SignUpOK) SetRefreshToken(val string) {
	s.RefreshToken = val
}

func (*SignUpOK) signUpRes() {}

type SignUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *SignUpReq) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *SignUpReq) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *SignUpReq) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *SignUpReq) SetPassword(val string) {
	s.Password = val
}

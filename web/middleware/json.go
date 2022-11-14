package middleware

// Json general interface what represent the request payload
type Json interface {
	Validate() error
}

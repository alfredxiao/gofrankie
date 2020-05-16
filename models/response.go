package models

// PuppyObject represents a success response
type PuppyObject struct {
	Puppy bool `json:"puppy" binding:"required"`
}

// ErrorObject represents an error response
// TODO: Reconsider swagger doc which uses this object for 500 error,
//       but we are using it for 400 error at the moment
type ErrorObject struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
}

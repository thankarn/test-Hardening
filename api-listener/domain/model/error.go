package model

type ErrorProps struct {
	Error     error  `json:"error"`
	FileError string `json:"file_error"`
	LineError int    `json:"line_error"`
}

package models

type Request struct {
	Command string `json:"command"`
}

type Response struct {
	Output string `json:"output"`
}

package structure

type JsonResponse struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Data    []Movie `json:"data"`
}

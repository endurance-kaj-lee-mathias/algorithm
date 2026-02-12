package health

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type Status struct {
	Backend string `json:"backend"`
}

package types

type TaskRequest struct {
	Input string `json:"input"`
}

type TaskResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream,omitempty"`
}

type GenerateResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

// /api/tags
// POST без тела, ответ:
type TagsResponse struct {
	Models []string `json:"models"`
}

// /api/show
// POST {"model": "modelname"}
type ShowRequest struct {
	Model string `json:"model"`
}
type ShowResponse struct {
	Model   string            `json:"model"`
	Details map[string]string `json:"details"`
}

// /api/pull
// POST {"name": "modelname"}
type PullRequest struct {
	Name string `json:"name"`
}
type PullResponse struct {
	Status string `json:"status"`
}

// /api/create
// POST {"name": "modelname", "modelfile": "..."}
type CreateRequest struct {
	Name      string `json:"name"`
	ModelFile string `json:"modelfile"`
}
type CreateResponse struct {
	Status string `json:"status"`
}

// /api/delete
// POST {"model": "modelname"}
type DeleteRequest struct {
	Model string `json:"model"`
}
type DeleteResponse struct {
	Status string `json:"status"`
}

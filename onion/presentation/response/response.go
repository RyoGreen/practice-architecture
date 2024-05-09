package response

type JobResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type JobResponseWithSecret struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Content       string `json:"content"`
	SecretContent string `json:"secret_content"`
}

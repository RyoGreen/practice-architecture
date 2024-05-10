package request

type JobRequest struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Content       string `json:"content"`
	SecretContent string `json:"secret_content"`
}

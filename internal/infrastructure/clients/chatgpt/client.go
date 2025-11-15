package chatgpt

const (
	basePath = "v1/chat/completions"
	systemPrompt = "Você é um sistema que recebe um texto transcrito de um áudio e organiza, separando em parágrafos e corrigindo possíveis erros de concordância."
	resumeUserPrompt = "Preciso de um objeto com título sugerido, descrição sugerida, resumo breve e médio sobre a seguinte transcrição:"
	functionCallName = "resume"
)

type (
	ChatgptFunctionCallRequest struct {
		Model string `json:"model"`
		Messages []Message `json:"messages"`
	}

	Message struct {
		Role string `json:"role"`
		Content string `json:"content"`
	}
)
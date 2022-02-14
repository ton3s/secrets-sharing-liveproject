package types

// Request & Response structs
type CreateSecretRequest struct {
	Secret string `json:"plain_text"`
}

type CreateSecretResponse struct {
	Id string `json:"id"`
}

type GetSecretResponse struct {
	Data string `json:"data"`
}

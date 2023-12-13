package dto

type GenSessionRequestDTO struct {
	SessionID string `json:"session_id" conform:"trim" validate:"required,uuid4" error:"access_denied_1"`
	PublicKey string `json:"public_key" validate:"required" error:"access_denied_2"`
}

type GenSessionResponseDTO struct {
	Data string `json:"data"`
}

type GetKeyPairResponseDTO struct {
	PrivateKey string
	PublicKey  string
}

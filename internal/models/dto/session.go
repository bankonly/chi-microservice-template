package dto

type GenSessionRequestDTO struct {
	Data string `json:"data" validate:"required" error:"access_denied_2"`
	Enk  string `json:"enk" validate:"required" error:"access_denied_3"`
	IV   string `json:"iv" validate:"required" error:"access_denied_4"`
}

type GenSessionResponseDTO struct {
	SessionId     string `json:"session_id"`
	PlainSessinId string `json:"-"`
}

type GetKeyPairResponseDTO struct {
	PrivateKey string
	PublicKey  string
}

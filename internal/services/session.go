package services

import (
	"ecm-api-template/internal/caches"
	errorConf "ecm-api-template/internal/configs/error-conf"
	"ecm-api-template/internal/models/dto"

	"github.com/bankonly/go-pkg/v1/encryption"
)

type Session interface {
	GenSession(requestId, vector, token, enk string) (*dto.GenSessionResponseDTO, error)
	GetSession(requestId string)
}

type SessionOpts struct {
	sessionCache caches.Session
}

func NewSession(sessionCache caches.Session) Session {
	return &SessionOpts{
		sessionCache: sessionCache,
	}
}

func (opts *SessionOpts) GenSession(requestId, vector, token, enk string) (*dto.GenSessionResponseDTO, error) {
	sessionId, err := encryption.RSADecAESRandomKey(enk, token, vector)
	if err != nil {
		return nil, errorConf.ErrBadParameter
	}

	err = opts.sessionCache.Save(requestId, sessionId, "")
	if err != nil {
		return nil, errorConf.ErrBadParameter
	}

	res := dto.GenSessionResponseDTO{
		SessionId:     sessionId,
		PlainSessinId: sessionId,
	}

	return &res, nil
}

func (opts *SessionOpts) GetSession(requestId string) {

}

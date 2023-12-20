package services

import (
	"ecm-api-template/internal/caches"
	errorConf "ecm-api-template/internal/configs/error-conf"
	"ecm-api-template/internal/models/dto"

	"github.com/bankonly/go-pkg/v1/encryption"
)

type Session interface {
	GenSession(requestId, vector, encryptedSessionId, enk string) (*dto.GenSessionResponseDTO, error)
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

func (opts *SessionOpts) GenSession(requestId, vector, encryptedSessionId, enk string) (*dto.GenSessionResponseDTO, error) {
	sessionId, err := encryption.RSADecAESRandomKey(enk, encryptedSessionId, vector)
	if err != nil {
		return nil, errorConf.ErrBadParemeter
	}

	err = opts.sessionCache.Save(requestId, sessionId, "")
	if err != nil {
		return nil, errorConf.ErrBadParemeter
	}

	res := dto.GenSessionResponseDTO{
		SessionId:     sessionId,
		PlainSessinId: sessionId,
	}

	return &res, nil
}

func (opts *SessionOpts) GetSession(requestId string) {

}

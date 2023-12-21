package services

import (
	"ecm-api-template/internal/caches"
	errorConf "ecm-api-template/internal/configs/error-conf"
	"ecm-api-template/internal/models/dto"

	"github.com/bankonly/go-pkg/v1/encryption"
)

type SessionService interface {
	GenSession(requestId, vector, token, enk string) (*dto.GenSessionResponseDTO, error)
	GetSession(requestId string)
}

type SessionServiceOpts struct {
	sessionCache caches.SessionCache
}

func NewSessionService(sessionCache caches.SessionCache) SessionService {
	return &SessionServiceOpts{
		sessionCache: sessionCache,
	}
}

func (opts *SessionServiceOpts) GenSession(requestId, vector, token, enk string) (*dto.GenSessionResponseDTO, error) {
	sessionId, err := encryption.RSADecAESRandomKey(enk, token, vector)
	if err != nil {
		return nil, errorConf.ErrAccessDenied1
	}

	if sessionId == "" {
		return nil, errorConf.ErrAccessDenied2
	}

	err = opts.sessionCache.Save(requestId, sessionId, "")
	if err != nil {
		return nil, errorConf.ErrAccessDenied3
	}

	res := dto.GenSessionResponseDTO{
		SessionId:      sessionId,
		PlainSessionId: sessionId,
	}

	return &res, nil
}

func (opts *SessionServiceOpts) GetSession(requestId string) {

}

package services

import (
	"ecm-api-template/internal/caches"
	"ecm-api-template/internal/models/dto"
	"os"

	"github.com/bankonly/go-pkg/v1/stacktrace"
)

type Session interface {
	GenSession(requestId string, body *dto.GenSessionRequestDTO) (*dto.GenSessionResponseDTO, error)
	GetSession(requestId string)
	GetKeyPair() (*dto.GetKeyPairResponseDTO, error) // private key, public key
}

type SessionOpts struct {
	sessionCache caches.Session
}

func NewSession(sessionCache caches.Session) Session {
	return &SessionOpts{
		sessionCache: sessionCache,
	}
}

func (opts *SessionOpts) GenSession(requestId string, body *dto.GenSessionRequestDTO) (*dto.GenSessionResponseDTO, error) {
	err := opts.sessionCache.Save(requestId, body.SessionID, body.PublicKey)
	if err != nil {
		return nil, stacktrace.BadRequest("bad_request")
	}

	selfKey, err := opts.GetKeyPair()
	if err != nil {
		return nil, stacktrace.BadRequest("credential_failed")
	}

	res := dto.GenSessionResponseDTO{
		Data: selfKey.PublicKey,
	}
	return &res, nil
}
func (opts *SessionOpts) GetSession(requestId string) {}

// Get current key pair
func (opts *SessionOpts) GetKeyPair() (*dto.GetKeyPairResponseDTO, error) {
	pkFile, err := os.ReadFile(opts.privateKeyPath)
	if err != nil {
		return nil, err
	}

	pbkFile, err := os.ReadFile(opts.publicKeyPath)
	if err != nil {
		return nil, err
	}

	selfKey := dto.GetKeyPairResponseDTO{
		PrivateKey: string(pkFile),
		PublicKey:  string(pbkFile),
	}

	return &selfKey, nil
}

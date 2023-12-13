package services

type Services struct {
	Session Session
}

func New() *Services {
	return &Services{
		Session: NewSession(),
	}
}

package services

// IFactory interface is an abstraction of services runnable code
type IFactory interface {
	Run()
}

// ServiceType to be yield
type ServiceType string

// types available on servicetype
const (
	WriteXlsxResults = "writeXlsx"
)

// Factory to yield an instance for some service
func Factory(svc ServiceType) IFactory {
	switch svc {
	case WriteXlsxResults:
		return newXlsxWriter()
	default:
		return newXlsxWriter()
	}
}

package storage

// Service storage interface shared between
type Service interface {
	NewCode(code, url string) (string, error)
	GetURL(code string) (UrlCode, error)
	NewVisit(code string) error
	CodeInfo(code string) ([]CodeVisit, error)
	Close() error
}

// UrlCode struct to save
type UrlCode struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

// CodeVisit struct to save visits
type CodeVisit struct {
	Code string `json:"code"`
	Time string `json:"timestamp"`
}

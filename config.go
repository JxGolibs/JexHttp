package JexHttp

import "time"

type Config struct {
	Https          bool `default:"false"`
	CertFile       string
	KeyFile        string
	Listening      string        `default:":8080"`
	SessionExpires time.Duration `default:"60"`
}

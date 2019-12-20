package JexHttp

import "time"

type Config struct {
	Https     bool `default:"false"`
	CertFile  string
	KeyFile   string
	Listening string `default:":8080"`
	Session   struct {
		Cookie  string        `default:"JexGolib_session_id"`
		Expires time.Duration `default:"60"`  //有效时间
		SecureCookie    bool `default:"false"` 
	}
}

package auth

import (
	"errors"
	"time"

	"github.com/go-http-utils/cookie"
	uuid "github.com/satori/go.uuid"
)

type Session struct {
	Id      string
	Expired int64
	Option  *cookie.Options
}

var Map map[string]*Session

func init() {
	Map = map[string]*Session{}
}

func NewSession(options ...*cookie.Options) *Session { // maxAge unit: second. 0 is never expried.

	var option *cookie.Options

	if len(options) > 0 {
		option = options[0]
	} else {
		option = &cookie.Options{
			MaxAge:   1800, // optional
			Path:     "/",  // optional, default to "/"
			HTTPOnly: true, // optional, default to `true``
		}
	}

	uid := uuid.NewV4()

	s := &Session{
		Id:      uid.String(),
		Expired: int64(option.MaxAge) + time.Now().Unix(),
		Option:  option,
	}

	Map[uid.String()] = s

	return s
}

func Exists(id string) (bool, error) { // 判断是否存在或过期
	s, exists := Map[id]
	if !exists {
		return false, errors.New("not found")
	}

	if s.Expired < time.Now().Unix() {
		return false, errors.New("expired")
	}
	return true, nil
}

func (s *Session) Destroy() {
	defer destory(s.Id)
}

func (s *Session) Update() {
	s.Expired = int64(s.Option.MaxAge) + time.Now().Unix()
}

func destory(id string) {
	Map[id] = nil
}

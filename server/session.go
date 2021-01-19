package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

// sessionHandler implements
type sessionHandler struct {
	handler http.Handler
}

// Test to see if sessionHandler implements http.Handler
var _ http.Handler = &sessionHandler{}

// NewSessionHandler creates a new sessionHandler
func NewSessionHandler(h http.Handler) *sessionHandler {
	return &sessionHandler{handler: h}
}

// ServeHTTP wraps handlers to set "user" Cookie and logging
func (s *sessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	user := "unknown"
	newSession := false
	cookie, err := r.Cookie("user")
	if err != nil {
		n := newUserCookie()
		http.SetCookie(w, n)
		r.AddCookie(n) // Add it in request as well
		user = n.Value // n.Value is same as sanitized by AddCookie
		newSession = true
	} else {
		user = cookie.Value
	}

	// If path contains "/public/images" set cache enable
	enableImageCache(w, r.URL.Path)

	ctx := context.WithValue(r.Context(), "user", user)
	rc := r.WithContext(ctx)
	s.handler.ServeHTTP(w, rc)
	end := time.Since(start)
	logger.Info("Request",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("user", user),
		zap.String("host", r.Host),
		zap.String("IP", r.RemoteAddr),
		zap.Bool("new", newSession),
		zap.Int64("ms", end.Milliseconds()))
}

// newUserCookie
//  To encrypt or decrypt read
//  https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
func newUserCookie() *http.Cookie {
	return &http.Cookie{
		Name:     sessionIDKey,
		Value:    uuid.New().String(),
		Expires:  time.Now().Add(dayHrs), // One day
		MaxAge:   daySeconds,
		HttpOnly: true,
	}
}

// enableImageCache on images directory
func enableImageCache(w http.ResponseWriter, path string) {
	if path != "" && strings.Contains(path, "/images/") {
		age := fmt.Sprintf("max-age:%d, public", daySeconds)
		w.Header().Set("Cache-Control", age)
		now := time.Now().Format(http.TimeFormat)
		w.Header().Set("Last-Modified", now)
		expire := time.Now().Add(dayHrs).Format(http.TimeFormat)
		w.Header().Set("Expires", expire)
		logger.Debug("Cache-Control Set", zap.String("path", path))
	}
}

package middleware

import "net/http"

type Wrapper func(next http.Handler) http.Handler

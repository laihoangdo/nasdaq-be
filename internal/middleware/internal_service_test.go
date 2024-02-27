package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"nasdaqvfs/config"
	"nasdaqvfs/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMiddleware_ValidateInternalService(t *testing.T) {
	type input struct {
		key    string
		origin string
	}

	tcs := map[string]struct {
		givenInput input
		expBody    string
		expCode    int
	}{
		"success": {
			givenInput: input{
				key:    "api-key-1",
				origin: "api-origin-1",
			},
			expBody: `
			{
				"code": 0
			}
			`,
			expCode: http.StatusOK,
		},
		"error: empty origin": {
			givenInput: input{
				key: "api-key-1",
			},
			expBody: `
			{
				"code": 401,
				"message": "origin is empty"
			}
			`,
			expCode: http.StatusUnauthorized,
		},
		"error: empty api key": {
			givenInput: input{
				origin: "api-origin-1",
			},
			expBody: `
			{
				"code": 401,
				"message": "key is empty"
			}
			`,
			expCode: http.StatusUnauthorized,
		},
		"error: invalid api key": {
			givenInput: input{
				key:    "api-key-10",
				origin: "api-origin-1",
			},
			expBody: `
			{
				"code": 401,
				"message": "key is invalid"
			}
			`,
			expCode: http.StatusUnauthorized,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// GIVEN
			cfg := &config.Config{
				Logger: config.Logger{
					Level:             "debug",
					Development:       true,
					DisableCaller:     false,
					DisableStacktrace: false,
					Encoding:          "json",
				},
				InternalAPI: config.InternalAPIConfig{
					AcceptedKeys: []string{"api-key-1", "api-key-2"},
				},
			}

			appLogger := logger.NewApiLogger(cfg)
			appLogger.InitLogger()

			mw := NewMiddlewareManager(cfg, nil, appLogger, nil, nil)

			w := httptest.NewRecorder()
			c, e := gin.CreateTestContext(w)

			req, _ := http.NewRequest(http.MethodGet, "/", nil)

			req.Header.Set(internalAPIKeyHeader, tc.givenInput.key)
			req.Header.Set(internalAPIOriginHeader, tc.givenInput.origin)

			c.Request = req

			// WHEN
			e.Use(mw.ValidateInternalService())
			e.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
				})
			})
			e.ServeHTTP(w, req)

			// THEN
			require.Equal(t, tc.expCode, w.Code)
			require.JSONEq(t, tc.expBody, w.Body.String())
		})
	}
}

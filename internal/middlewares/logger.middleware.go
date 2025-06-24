package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"time"

	"system-management-pg/global"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// LogApiSuccess represents the structure for successful API log
type LogApiSuccess struct {
	IDUserGuest  string                 `json:"id_user_guest"`
	UserAgent    string                 `json:"userAgent"`
	ClientIP     string                 `json:"clientIp"`
	Time         time.Time              `json:"time"`
	Duration     int64                  `json:"duration"`
	Message      string                 `json:"message"`
	BodyRequest  string                 `json:"bodyRequest"`
	BodyResponse string                 `json:"bodyResponse"`
	Method       string                 `json:"method"`
	Params       map[string]interface{} `json:"params"`
	Path         string                 `json:"path"`
	StatusCode   int                    `json:"statusCode"`
}

// LogApiError represents the structure for error API log
type LogApiError struct {
	IDUserGuest  string                 `json:"id_user_guest"`
	UserAgent    string                 `json:"userAgent"`
	ClientIP     string                 `json:"clientIp"`
	Time         time.Time              `json:"time"`
	Duration     int64                  `json:"duration"`
	Message      string                 `json:"message"`
	BodyRequest  string                 `json:"bodyRequest"`
	BodyResponse string                 `json:"bodyResponse"`
	Method       string                 `json:"method"`
	Params       map[string]interface{} `json:"params"`
	Path         string                 `json:"path"`
	StatusCode   int                    `json:"statusCode"`
}

const (
	IndexLogApiSuccess = "log-api-system-management-success"
	IndexLogApiError   = "log-api-system-management-error"
)

// SaveLogApiSuccess saves successful API log to Elasticsearch
func SaveLogApiSuccess(logData LogApiSuccess) {
	body, err := json.Marshal(logData)
	if err != nil {
		global.Logger.Error("Failed to marshal success log", zap.Error(err))
		return
	}

	_, err = global.EsClient.Index(
		IndexLogApiSuccess,
		bytes.NewReader(body),
		global.EsClient.Index.WithDocumentID(uuid.New().String()),
	)
	if err != nil {
		global.Logger.Error("Failed to save success log to Elasticsearch", zap.Error(err))
	}
}

// SaveLogApiError saves error API log to Elasticsearch
func SaveLogApiError(logData LogApiError) {
	body, err := json.Marshal(logData)
	if err != nil {
		global.Logger.Error("Failed to marshal error log", zap.Error(err))
		return
	}

	_, err = global.EsClient.Index(
		IndexLogApiError,
		bytes.NewReader(body),
		global.EsClient.Index.WithDocumentID(uuid.New().String()),
	)
	if err != nil {
		global.Logger.Error("Failed to save error log to Elasticsearch", zap.Error(err))
	}
}

// formatDate formats time to match the NestJS format
func formatDate(t time.Time) string {
	return t.Format("15:04:05 - 02/01/2006")
}

// LogApiMiddleware logs API requests and responses
func LogApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip logging for specific paths
		path := c.Request.URL.Path
		if path == "/metrics" || path == "/api/metrics" || path == "/api/v1/metrics" ||
			path == "/api/v1/employees/register-face" || path == "/api/v1/payment" ||
			strings.HasPrefix(path, "/api/v1/upload/view-image") {
			c.Next()
			return
		}

		startTime := time.Now()
		idUserGuest := c.GetHeader("id_user_guest")
		idUserGuestNew := "Guest-" + uuid.New().String()
		if idUserGuest == "" || idUserGuest == "undefined" {
			idUserGuest = idUserGuestNew
			c.Header("id_user_guest", idUserGuest)
		}

		userAgent := c.GetHeader("user-agent")
		clientIP := c.ClientIP()
		method := c.Request.Method
		params := make(map[string]interface{})
		for k, v := range c.Request.URL.Query() {
			if len(v) > 0 {
				params[k] = v[0]
			}
		}

		// Read request body
		var bodyRequest string
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil && len(bodyBytes) > 0 {
					bodyRequest = string(bodyBytes) // Keep as JSON string
			} else {
				bodyRequest = "No data"
			}
			// Restore body for downstream handlers
			c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		} else {
			bodyRequest = "No data"
		}

		// Capture response
		writer := &bodyLogWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
		c.Writer = writer

		// Process request
		c.Next()

		duration := time.Since(startTime).Milliseconds()
		statusCode := c.Writer.Status()
		message := "" // Customize based on your needs, e.g., retrieve from a response metadata if available

		// Capture response body as JSON string
		var bodyResponse string
		responseBytes := writer.body.Bytes()
		if len(responseBytes) > 0 {
			bodyResponse = string(responseBytes) // Keep as JSON string
		} else {
			bodyResponse = "No data"
		}

		// Log success
		if statusCode < 400 {
			SaveLogApiSuccess(LogApiSuccess{
				IDUserGuest:  idUserGuest,
				UserAgent:    userAgent,
				ClientIP:     clientIP,
				Time:         time.Now(),
				Duration:     duration,
				Message:      message,
				BodyRequest:  bodyRequest,
				BodyResponse: bodyResponse,
				Method:       method,
				Params:       params,
				Path:         path,
				StatusCode:   statusCode,
			})
		} else {
			// Log error
			errorMessage := "Unknown error"
			if len(responseBytes) > 0 {
				var errorResp map[string]interface{}
				if json.Unmarshal(responseBytes, &errorResp) == nil {
					if msg, ok := errorResp["message"].(string); ok {
						errorMessage = msg
					}
				}
			}

			SaveLogApiError(LogApiError{
				IDUserGuest:  idUserGuest,
				UserAgent:    userAgent,
				ClientIP:     clientIP,
				Time:         time.Now(),
				Duration:     duration,
				Message:      message,
				BodyRequest:  bodyRequest,
				BodyResponse: errorMessage,
				Method:       method,
				Params:       params,
				Path:         path,
				StatusCode:   statusCode,
			})
		}
	}
}

// bodyLogWriter captures response body
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
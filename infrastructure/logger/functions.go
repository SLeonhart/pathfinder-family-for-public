package logger

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	//RequestIDKey uuid для Jin
	RequestIDKey = "X-Request-ID"
)

type key int

const (
	keyRequestID key = iota
	keyStartTime key = iota
)

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// Get получить Logger
func Get() *logrus.Logger {
	return logger
}

// TraceWithFields послать Trace сообщение с полями
func TraceWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.TraceLevel, fields, message)
}

// DebugWithFields послать Debug сообщение с полями
func DebugWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.DebugLevel, fields, message)
}

// InfoWithFields послать Info сообщение с полями
func InfoWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.InfoLevel, fields, message)
}

// WarnWithFields послать Warn сообщение с полями
func WarnWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.WarnLevel, fields, message)
}

// ErrorWithFields послать Error сообщение с полями
func ErrorWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.ErrorLevel, fields, message)
}

// FatalWithFields послать Fatal сообщение с полями
func FatalWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.FatalLevel, fields, message)
}

// PanicWithFields послать Panic сообщение с полями
func PanicWithFields(requestID string, fields map[string]interface{}, message string) {
	logWithFields(requestID, logrus.PanicLevel, fields, message)
}

// TracefWithFields послать форматированное Trace сообщение с полями
func TracefWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	TraceWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// DebugfWithFields послать форматированное Debug сообщение с полями
func DebugfWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	DebugWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// InfofWithFields послать форматированное Info сообщение с полями
func InfofWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	InfoWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// WarnfWithFields послать форматированное Warn сообщение с полями
func WarnfWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	WarnWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// ErrorfWithFields послать форматированное Error сообщение с полями
func ErrorfWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	ErrorWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// FatalfWithFields послать форматированное Fatal сообщение с полями
func FatalfWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	FatalWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// PanicfWithFields послать форматированное Panic сообщение с полями
func PanicfWithFields(requestID string, fields map[string]interface{}, format string, args ...interface{}) {
	PanicWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// Tracef послать Trace сообщение с интерфесом
func Tracef(requestID string, format string, args ...interface{}) {
	TracefWithFields(requestID, nil, format, args...)
}

// Debugf послать Debug сообщение с интерфесом
func Debugf(requestID string, format string, args ...interface{}) {
	DebugfWithFields(requestID, nil, format, args...)
}

// Infof послать Info сообщение с интерфесом
func Infof(requestID string, format string, args ...interface{}) {
	InfofWithFields(requestID, nil, format, args...)
}

// Warnf послать Warn сообщение с интерфесом
func Warnf(requestID string, format string, args ...interface{}) {
	WarnfWithFields(requestID, nil, format, args...)
}

// Errorf послать Error сообщение с интерфесом
func Errorf(requestID string, format string, args ...interface{}) {
	ErrorfWithFields(requestID, nil, format, args...)
}

// Fatalf послать Fatal сообщение с интерфесом
func Fatalf(requestID string, format string, args ...interface{}) {
	FatalfWithFields(requestID, nil, format, args...)
}

// Panicf послать Panic сообщение с интерфесом
func Panicf(requestID string, format string, args ...interface{}) {
	PanicfWithFields(requestID, nil, format, args...)
}

// Trace послать Trace сообщение
func Trace(requestID string, message string) {
	TraceWithFields(requestID, nil, message)
}

// Debug послать Debug сообщение
func Debug(requestID string, message string) {
	DebugWithFields(requestID, nil, message)
}

// Info послать Info сообщение
func Info(requestID string, message string) {
	InfoWithFields(requestID, nil, message)
}

// Warn послать Warn сообщение
func Warn(requestID string, message string) {
	WarnWithFields(requestID, nil, message)
}

// Error послать Error сообщение
func Error(requestID string, message string) {
	ErrorWithFields(requestID, nil, message)
}

// Fatal послать Fatal сообщение
func Fatal(requestID string, message string) {
	FatalWithFields(requestID, nil, message)
}

// Panic послать Panic сообщение
func Panic(requestID string, message string) {
	PanicWithFields(requestID, nil, message)
}

// ErrorWithErr передача Error сообщение с ошибкой
func ErrorWithErr(requestID string, message string, err error) {
	fields := make(map[string]interface{}, 0)
	fields["err"] = err
	ErrorWithFields(requestID, fields, message)
}

// WarnWithErr передача Warn сообщение с ошибкой
func WarnWithErr(requestID string, message string, err error) {
	fields := make(map[string]interface{}, 0)
	fields["err"] = err
	WarnWithFields(requestID, fields, message)
}

// PanicWithErr передача Panic сообщение с ошибкой
func PanicWithErr(requestID string, message string, err error) {
	fields := make(map[string]interface{}, 0)
	fields["err"] = err
	PanicWithFields(requestID, fields, message)
}

// FatalWithErr передача Fatal сообщение с ошибкой
func FatalWithErr(requestID string, message string, err error) {
	fields := make(map[string]interface{}, 0)
	fields["err"] = err
	FatalfWithFields(requestID, fields, message)
}

// WarnfWithFieldsAndErr послать форматированное Warn сообщение с полями и ошибкой
func WarnfWithFieldsAndErr(requestID string, fields map[string]interface{}, err error, format string, args ...interface{}) {
	fields["err"] = err
	WarnWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// ErrorfWithFieldsAndErr послать форматированное Error сообщение с полями и ошибкой
func ErrorfWithFieldsAndErr(requestID string, fields map[string]interface{}, err error, format string, args ...interface{}) {
	fields["err"] = err
	ErrorWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// FatalfWithFieldsAndErr послать форматированное Fatal сообщение с полями и ошибкой
func FatalfWithFieldsAndErr(requestID string, fields map[string]interface{}, err error, format string, args ...interface{}) {
	fields["err"] = err
	FatalWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// PanicfWithFieldsAndErr послать форматированное Panic сообщение с полями и ошибкой
func PanicfWithFieldsAndErr(requestID string, fields map[string]interface{}, err error, format string, args ...interface{}) {
	fields["err"] = err
	PanicWithFields(requestID, fields, fmt.Sprintf(format, args...))
}

// WarnWithFieldsAndErr послать Warn сообщение с полями и ошибкой
func WarnWithFieldsAndErr(requestID string, fields map[string]interface{}, message string, err error) {
	fields["err"] = err
	logWithFields(requestID, logrus.WarnLevel, fields, message)
}

// ErrorWithFieldsAndErr послать Error сообщение с полями и ошибко
func ErrorWithFieldsAndErr(requestID string, fields map[string]interface{}, message string, err error) {
	fields["err"] = err
	logWithFields(requestID, logrus.ErrorLevel, fields, message)
}

// FatalWithFieldsAndErr послать Fatal сообщение с полями и ошибко
func FatalWithFieldsAndErr(requestID string, fields map[string]interface{}, message string, err error) {
	fields["err"] = err
	logWithFields(requestID, logrus.FatalLevel, fields, message)
}

// PanicWithFieldsAndErr послать Panic сообщение с полями и ошибко
func PanicWithFieldsAndErr(requestID string, fields map[string]interface{}, message string, err error) {
	fields["err"] = err
	logWithFields(requestID, logrus.PanicLevel, fields, message)
}

func logWithFields(requestID string, level logrus.Level, fields map[string]interface{}, message string) {
	_logger := Get()
	_logger.SetFormatter(&logrus.JSONFormatter{})
	if fields == nil {
		fields = make(map[string]interface{}, 0)
	}

	if requestID != "" {
		fields["request.id"] = requestID
	}

	if level == logrus.DebugLevel || level == logrus.TraceLevel || level == logrus.ErrorLevel || level == logrus.FatalLevel || level == logrus.PanicLevel {
		fields["stacktrace"] = getStackTrace()
	}
	fields["goid"] = goid()
	entry := _logger.WithFields(fields)
	switch level {
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		entry.Logger.Out = os.Stderr
	default:
		entry.Logger.Out = os.Stdout
	}
	entry.Log(level, message)
}

func getStackTrace() string {
	i := 5
	fullFuncName := ""
	for {
		pc := make([]uintptr, 15)
		n := runtime.Callers(i, pc)
		i++
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		if len(frame.Function) == 0 {
			break
		}
		funcSplit := strings.Split(frame.Function, "/")
		funcName := funcSplit[len(funcSplit)-1]
		if strings.Contains(funcName, "log-formatter") {
			continue
		}
		if strings.Contains(funcName, "runtime") {
			continue
		}
		if len(fullFuncName) == 0 {
			fullFuncName = fmt.Sprintf("[%v]", funcName)
		} else {
			fullFuncName = fmt.Sprintf("[%v].%v", funcName, fullFuncName)
		}
	}
	return fullFuncName
}

// CreateFieldsFromGinContext создать context для GIN
func CreateFieldsFromGinContext(ctx *gin.Context) map[string]interface{} {
	if ctx == nil {
		return logrus.Fields{}
	}

	fields := make(map[string]interface{}, 0)

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	fields["request.method"] = ctx.Request.Method
	fields["request.url"] = ctx.Request.URL
	fields["request.header"] = ctx.Request.Header
	fields["request.body"] = string(body)
	fields["request.id"] = getRequestID(ctx)

	return fields
}

// CreateRequestIDField ?
func CreateRequestIDField(ctx context.Context) string {
	requestID := ctx.Value("X-Request-ID")
	if requestID == nil {
		requestID = ctx.Value("request.id")
		if requestID == nil {
			return ""
		}
	}

	return requestID.(string)
}

func sendMessageUndetectedLevel(requestID string, fields map[string]interface{}, message string, loglvl string) {
	level, err := logrus.ParseLevel(loglvl)
	if err != nil {
		level = logrus.TraceLevel
	}

	switch level {
	case logrus.DebugLevel:
		DebugWithFields(requestID, fields, message)
	case logrus.TraceLevel:
		TraceWithFields(requestID, fields, message)
	case logrus.WarnLevel:
		WarnWithFields(requestID, fields, message)
	case logrus.ErrorLevel:
		ErrorWithFields(requestID, fields, message)
	case logrus.FatalLevel:
		FatalWithFields(requestID, fields, message)
	case logrus.PanicLevel:
		PanicWithFields(requestID, fields, message)
	case logrus.InfoLevel:
		InfoWithFields(requestID, fields, message)
	}

}

// CreateContextRequest создание context для запросов
func CreateContextRequest() context.Context {
	ctx := context.Background()

	reqId := uuid.New().String()
	ctx = context.WithValue(ctx, keyRequestID, reqId)
	ctx = context.WithValue(ctx, RequestIDKey, reqId)

	start := time.Now().Format("2006-01-02T15:04:05-0700")
	ctx = context.WithValue(ctx, keyStartTime, start)
	return ctx
}

// CopyRequestIdFromCtx копирование идентификатора запроса из контекста в контекст
func CopyRequestIdFromCtx(toCtx context.Context, fromCtx context.Context) context.Context {
	toCtx = context.WithValue(toCtx, keyRequestID, CreateRequestIDField(fromCtx))
	toCtx = context.WithValue(toCtx, RequestIDKey, CreateRequestIDField(fromCtx))
	return toCtx
}

// CopyRequestIdFromGinCtx копирование идентификатора запроса из контекста gin в контекст
func CopyRequestIdFromGinCtx(toCtx context.Context, fromCtx *gin.Context) context.Context {
	toCtx = context.WithValue(toCtx, keyRequestID, getRequestID(fromCtx))
	toCtx = context.WithValue(toCtx, RequestIDKey, getRequestID(fromCtx))
	return toCtx
}

// LogRestyResponse логирование Resty ответа
func LogRestyResponse(requestID string, response *resty.Response, prefix string, loglvl string) {
	if response == nil || response.Request == nil {
		return
	}

	fields := make(map[string]interface{}, 10)

	fields[prefix+"request."+"method"] = response.Request.Method
	fields[prefix+"request."+"url"] = response.Request.URL

	// reqBody, _ := response.Request.RawRequest.GetBody()
	// byteReqBody, _ := ioutil.ReadAll(reqBody)

	// fields[prefix+"Request."+"header"] = response.Request.Header
	// fields[prefix+"request."+"body"] = string(byteReqBody)
	fields[prefix+"response."+"code"] = response.StatusCode()
	fields[prefix+"response."+"duration"] = response.Time().Seconds()
	fields[prefix+"response."+"body"] = string(response.Body())

	sendMessageUndetectedLevel(requestID, fields, "Request complete", loglvl)
}

// LogHTTPResponse логирование HTTP ответа
func LogHTTPResponse(response *http.Response, prefix string, loglvl string) {
	if response == nil || response.Request == nil {
		return
	}

	var requestID string
	fields := make(map[string]interface{}, 10)

	fields[prefix+"request."+"method"] = response.Request.Method
	fields[prefix+"request."+"url"] = response.Request.URL

	if requestID = fmt.Sprintf("%v", response.Request.Context().Value(keyRequestID)); requestID == "<nil>" {
		requestID = ""
		Warnf(requestID, "requestID is empty")
	}

	// fields[prefix+"Request."+"header"] = response.Request.Header
	fields[prefix+"request."+"body"] = response.Request.Body
	fields[prefix+"response."+"code"] = response.StatusCode

	requestStart, err := time.Parse("2006-01-02T15:04:05-0700", fmt.Sprintf("%v", response.Request.Context().Value(keyStartTime)))
	if err != nil {
		WarnWithErr(requestID, "Cant parce requestStartDate", err)
	} else {
		fields[prefix+"response."+"duration"] = time.Since(requestStart).Seconds()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ErrorWithErr(requestID, "Cant parce responce body", err)
	} else {
		fields[prefix+"response."+"body"] = string(body)
	}

	sendMessageUndetectedLevel(requestID, fields, "Request complete", loglvl)

}

func getRequestID(c *gin.Context) interface{} {
	if c == nil {
		return ""
	}
	requestID, _ := c.Get(RequestIDKey)
	return requestID
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// LogGinRequest логирование ответов Gin
func LogGinRequest(c *gin.Context) {

	if c == nil || c.IsAborted() {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	start := time.Now()
	requestID := uuid.New().String()

	c.Set(RequestIDKey, requestID)

	fields := CreateFieldsFromGinContext(c)
	InfoWithFields(requestID, fields, "Start request")

	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w

	c.Next()

	end := time.Now()
	latency := end.Sub(start)
	if latency > time.Minute {
		latency = latency - latency%time.Second
	}
	statusCode := c.Writer.Status()
	body := w.body.String()

	fields["response.code"] = statusCode
	fields["response.latency"] = latency.Seconds()
	fields["response.body"] = body

	InfoWithFields(requestID, fields, "Request complete")
}

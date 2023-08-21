### Директория для gin-миддлваров

Миддлвары реализуются в виде функций, кроме или вместо параметра ctx, могут быть переданы любые свои, которые требуются.

func Middleware(ctx context.Context) func(*gin.Context) {
return func(c *gin.Context) {
c.Next()
}
}

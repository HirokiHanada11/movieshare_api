package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)


func TestLoginHandler(t *testing.T) {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	req, _ := http.NewRequest("GET", "/", nil)

	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQwMTU0NmJkMWRhMzA0ZDc2NGNmZWUzYTJhZTVjZDBlNGY2ZjgyN2IiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoidGVzdHVzZXIiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbXNoYXJlLWF1dGhlbnRpY2F0aW9uIiwiYXVkIjoibXNoYXJlLWF1dGhlbnRpY2F0aW9uIiwiYXV0aF90aW1lIjoxNjQyODQ3MzYxLCJ1c2VyX2lkIjoiQ0tPYzN2UWJtcmVQc0c2SWlET0w2ak1pOXptMSIsInN1YiI6IkNLT2MzdlFibXJlUHNHNklpRE9MNmpNaTl6bTEiLCJpYXQiOjE2NDI4NDczNjEsImV4cCI6MTY0Mjg1MDk2MSwiZW1haWwiOiJ0ZXN0MkB0ZXN0LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJ0ZXN0MkB0ZXN0LmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.M195L2W4hKLLSw1kCqYWEMN9LBDxi7POhjhoGdOdALTLDyUucTbLFvN0Z5kLtI8Nx3_0o229UZVIp7Dvq27x9Wp7xd_8zeYBlRvNVb7-rhyMaTXFLajtlobjw2VYRhyJJbmzP-OJt1atxegOaLSoyJ8WrtQUpmI4qLu5WJtZd7RsW7cNHrvY_AWCLU_GB7ZW6HPTuw8N2SeBnMwg_6F255HkdVznluJCsrRTt0aj9UOi2w4P61h_aIcDsuL8yGzurVTYUC0F8P7-4oG7igoeshrRmZ4snAXLj5uFAfZstV8GlM1-QwaeHGDRh2nF69rz_vj6B7H8UE_C_bFqZc0cKQ"

	req.Header.Add("idToken", idToken)

	ginContext.Request = req

	// err := LoginHandler(ginContext)
	// if err != nil {
	// 	t.Logf("Error: %v", err)
	// 	t.Fail()
	// }
}

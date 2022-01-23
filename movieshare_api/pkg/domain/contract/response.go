package contract

import "firebase.google.com/go/auth"

type (
	VerificationResponse struct {
		Status string `json:"status"`
		User auth.UserInfo `json:"user,omitempty"`
	}
)
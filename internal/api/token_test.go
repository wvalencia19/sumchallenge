package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var secretKey = "secret"

func TestValidToken(t *testing.T) {
	tests := []struct {
		name       string
		expiration time.Duration
		wantErr    bool
	}{
		{
			name:       "valid token",
			expiration: 1 * time.Minute,
			wantErr:    false,
		},
		{
			name:       "expiredToken",
			expiration: 1 * time.Millisecond,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username := "username"
			token, err := GenerateToken(username, secretKey, tt.expiration)
			assert.NoError(t, err)
			time.Sleep(1 * time.Second)
			err = ValidToken(token, secretKey)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
		})
	}
}

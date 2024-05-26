package auth

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_authDto_Login(t *testing.T) {
	type args struct {
		Email    string
		Password string
	}
	tests := []struct {
		name    string
		args    args
		want    LoginDto
		wantErr bool
	}{
		{
			name: "test 1 success",
			args: args{
				Email:    "email@email.com",
				Password: "saidjasidisjd",
			},
			want: LoginDto{
				Email:    "email@email.com",
				Password: "saidjasidisjd",
			},
		},
		{
			name: "test with no email",
			args: args{
				Email:    "",
				Password: "saidjasidisjd",
			},
			wantErr: true,
		},
		{
			name: "test with password too short",
			args: args{
				Email:    "email@email.com",
				Password: "12345",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {

				gin.SetMode(gin.TestMode)
				_ = gin.Default()

				jsonBody, err := json.Marshal(tt.args)

				req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
				req.Header.Set("Content-Type", "application/json")

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = req

				dto := NewAuthDto()

				got, err := dto.Login(c)
				if (err != nil) != tt.wantErr {
					t.Errorf("CreateUserDto() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("CreateUserDto() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

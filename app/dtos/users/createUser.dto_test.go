package users

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCreateUserDto(t *testing.T) {
	type args struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		IsAdmin  bool   `json:"isAdmin"`
		GroupId  uint   `json:"groupId"`
	}
	tests := []struct {
		name    string
		args    args
		want    CreateUserDto
		wantErr bool
	}{
		{
			name: "Test CreateUserDto",
			args: args{
				Email:    "email@email.com",
				Password: "password",
				Name:     "name",
				IsAdmin:  true,
				GroupId:  1,
			},
			want: CreateUserDto{
				Email:    "email@email.com",
				Password: "password",
				Name:     "name",
				IsAdmin:  true,
				GroupID:  1,
			},
		},
		{
			name: "Test CreateUserDto with empty email",
			args: args{
				Password: "password",
				Name:     "name",
				IsAdmin:  true,
				GroupId:  1,
			},
			want:    CreateUserDto{},
			wantErr: true,
		},
		{
			name: "Test CreateUserDto with invalid email",
			args: args{
				Email:    "email",
				Password: "password",
				Name:     "name",
				IsAdmin:  true,
				GroupId:  1,
			},
			want:    CreateUserDto{},
			wantErr: true,
		},
		{
			name: "Test CreateUserDto with empty password",
			args: args{
				Email:   "email@email.com",
				Name:    "name",
				IsAdmin: true,
				GroupId: 1,
			},
			want:    CreateUserDto{},
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

				dto := NewUserDto()

				got, err := dto.CreateUser(c)
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

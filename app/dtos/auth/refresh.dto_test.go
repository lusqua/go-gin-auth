package auth

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lusqua/gin-auth/app/usecases"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_authDto_Refresh(t *testing.T) {
	type args struct {
		JTI string `json:"jti"`
	}

	sucessJTI := usecases.GenerateRandomString(32)

	tests := []struct {
		name    string
		args    args
		want    RefreshDto
		wantErr bool
	}{
		{
			name: "test 1 success",
			args: args{
				JTI: sucessJTI,
			},
			want: RefreshDto{
				JTI: sucessJTI,
			},
		},
		{
			name: "test with invalid JTI",
			args: args{
				JTI: "asdhaudas",
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

				a := &authDto{}
				got, err := a.Refresh(c)
				if (err != nil) != tt.wantErr {
					t.Errorf("Refresh() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Refresh() got = %v, want %v", got, tt.want)
				}

				if usecases.ValidateRandomString(got.JTI) == false {
					t.Errorf("JTI %s is invalid", got.JTI)
				}
			},
		)
	}
}

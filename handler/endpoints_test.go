package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helpers"
	"github.com/SawitProRecruitment/UserService/helpers/pgerrcode"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// Test rules:
// 1. Phone numbers must be at minimum 10 characters and maximum 13 characters.
// 2. Phone numbers must start with the Indonesia country code “+62”.
// 3. Full name must be at minimum 3 characters and maximum 60 characters.
// 4. Passwords must be minimum 6 characters and maximum 64 characters,
// containing at least 1 capital characters AND 1 number AND 1 special (non
// alpha-numeric) characters.
// 5. Otherwise, return 400 Bad Requests with the error message containing ALL fields that
// failed the validation and which rule they failed at.
func TestRegistration(t *testing.T) {
	type expectation struct {
		input      generated.UserRegistrationRequest
		output     interface{}
		statusCode int
		repoReturn []interface{}
	}

	var (
		id       = float32(2)
		ctrl     = gomock.NewController(t)
		errCodes helpers.ErrorCodes
	)

	expectations := []expectation{
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "AA",
				Password:    "",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`Phone numbers must start with the Indonesia country code "+62". Phone numbers must be at minimum 10 characters and maximum 13 characters. Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters. Error codes:%w`, errCodes.LoginError()).Error(),
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{},
				nil,
			},
		},
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "+62AA",
				Password:    "",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`Phone numbers must be at minimum 10 characters and maximum 13 characters. Phone numbers must be a number. Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters. Error codes:%w`, errCodes.LoginError()).Error(),
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{},
				nil,
			},
		},
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "+6288888888888",
				Password:    "T",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 number AND 1 special (non alpha-numeric) characters. Error codes:%w`, errCodes.LoginError()).Error(),
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{},
				nil,
			},
		},
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "+6288888888888",
				Password:    "T1",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 special (non alpha-numeric) characters. Error codes:%w`, errCodes.LoginError()).Error(),
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{},
				nil,
			},
		},
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "+6288888888888",
				Password:    "T1!",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`Passwords must be minimum 6 characters and maximum 64 characters. Error codes:%w`, errCodes.LoginError()).Error(),
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{},
				nil,
			},
		},
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "+6288888888888",
				Password:    "T1!foo",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`User already registered. Error codes:%w`, errCodes.LoginError()).Error(),
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{},
				pq.Error{
					Code: pgerrcode.UniqueViolation, // simulate the phone number is registered
				},
			},
		},
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "+6288888888888",
				Password:    "T1!foo",
			},
			output: generated.UserRegistrationResponse{
				UserId: &id,
			},
			statusCode: http.StatusBadRequest,
			repoReturn: []interface{}{
				generated.UserRegistrationResponse{
					UserId: &id,
				},
				nil,
			},
		},
	}

	e := echo.New()

	for i := 0; i < len(expectations); i++ {
		expectation := expectations[i]
		rec := httptest.NewRecorder()

		payload := bytes.NewBuffer([]byte(fmt.Sprintf(`{"phone_number":"%s", "name":"%s", "password":"%s"}`, expectation.input.PhoneNumber, expectation.input.Name, expectation.input.Password)))
		req := httptest.NewRequest(http.MethodPost, "/register", payload)
		c := e.NewContext(req, rec)

		repo := repository.NewMockRepositoryInterface(ctrl)
		repo.EXPECT().SetProfile(c, gomock.Any).Return(expectation.repoReturn...)

		s := NewServer(NewServerOptions{
			Repository: repo,
		})

		s.Register(c)

		assert.Equal(t, expectation.statusCode, rec.Result().StatusCode)

		var result interface{}
		if rec.Result().StatusCode == http.StatusOK {
			var _result generated.UserRegistrationResponse
			if json.Unmarshal(rec.Body.Bytes(), &_result) != nil {
				t.Fatal()
			}
			result = _result
		} else if rec.Result().StatusCode != http.StatusOK {
			var _result generated.MessageResponse
			if json.Unmarshal(rec.Body.Bytes(), &_result) != nil {
				t.Fatal()
			}
			result = _result
		}

		assert.Equal(t, expectation.output, result)
	}
}

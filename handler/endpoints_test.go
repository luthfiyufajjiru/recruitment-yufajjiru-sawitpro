package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helpers"
	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
	"github.com/SawitProRecruitment/UserService/helpers/pgerrcode"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const keyStr = "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"

type expectation[T generated.UserRegistrationRequest | generated.UserLoginRequest] struct {
	input      T
	output     interface{}
	statusCode int
	repoReturn []interface{}
}

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
	var (
		id   = 1
		ctrl = gomock.NewController(t)
	)

	expectations := []expectation[generated.UserRegistrationRequest]{
		{
			input: generated.UserRegistrationRequest{
				Name:        "fulan",
				PhoneNumber: "AA",
				Password:    "",
			},
			output: generated.MessageResponse{
				Message: fmt.Errorf(`Phone numbers must start with the Indonesia country code "+62". Phone numbers must be at minimum 10 characters and maximum 13 characters. Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters. Error codes:%w`, errorIndex.LoginError).Error(),
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
				Message: fmt.Errorf(`Phone numbers must be at minimum 10 characters and maximum 13 characters. Phone numbers must be a number. Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters. Error codes:%w`, errorIndex.LoginError).Error(),
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
				Message: fmt.Errorf(`Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 number AND 1 special (non alpha-numeric) characters. Error codes:%w`, errorIndex.LoginError).Error(),
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
				Message: fmt.Errorf(`Passwords must be minimum 6 characters and maximum 64 characters, containing at least 1 special (non alpha-numeric) characters. Error codes:%w`, errorIndex.LoginError).Error(),
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
				Message: fmt.Errorf(`Passwords must be minimum 6 characters and maximum 64 characters. Error codes:%w`, errorIndex.LoginError).Error(),
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
				Message: fmt.Errorf(`User already registered. Error codes:%w`, errorIndex.LoginError).Error(),
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
		req.Header.Set("content-type", "application/x-www-form-urlencoded")

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

func TestLogin(t *testing.T) {
	const timeConstant = "10s"
	helpers.InitializeJWT(keyStr, keyStr, timeConstant, timeConstant) // init JWT module

	var (
		err                error
		dummyToken         = "dummyToken" // we are not checking token with excpectation, but just test the response
		correctPhoneNumber = "+6288888888888"
		ctrl               = gomock.NewController(t)
	)

	expectations := []expectation[generated.UserLoginRequest]{
		{
			input: generated.UserLoginRequest{
				PhoneNumber: "AA",
				Password:    "",
			},
			output: generated.MessageResponse{
				Message: helpers.DRForbidden,
			},
			statusCode: http.StatusForbidden,
			repoReturn: []interface{}{
				nil,
			},
		},
		{
			input: generated.UserLoginRequest{
				PhoneNumber: correctPhoneNumber,
				Password:    "T1!foo",
			},
			output: generated.MessageResponse{
				Message: helpers.DRForbidden,
			},
			statusCode: http.StatusForbidden,
			repoReturn: []interface{}{
				sql.ErrNoRows,
			},
		},
		{
			input: generated.UserLoginRequest{
				PhoneNumber: correctPhoneNumber,
				Password:    "T1!foo",
			},
			output: generated.JWTTokens{
				AccessToken:  &dummyToken,
				RefreshToken: &dummyToken,
			},
			statusCode: http.StatusOK,
			repoReturn: []interface{}{
				nil,
			},
		},
	}

	e := echo.New()

	for i := 0; i < len(expectations); i++ {
		expectation := expectations[i]
		rec := httptest.NewRecorder()

		payload := bytes.NewBuffer([]byte(fmt.Sprintf(`{"phone_number":"%s", "password":"%s"}`, expectation.input.PhoneNumber, expectation.input.Password)))
		req := httptest.NewRequest(http.MethodPost, "/login", payload)
		req.Header.Set("content-type", "application/x-www-form-urlencoded")

		c := e.NewContext(req, rec)

		repo := repository.NewMockRepositoryInterface(ctrl)
		repo.EXPECT().ComparePassword(c, expectation.input.PhoneNumber, expectation.input.Password).Return(expectation.repoReturn...)

		s := NewServer(NewServerOptions{
			Repository: repo,
		})

		s.Login(c)

		assert.Equal(t, expectation.statusCode, rec.Result().StatusCode)

		switch rec.Result().StatusCode {
		case http.StatusOK:
			var JWTTokens generated.JWTTokens
			err = json.Unmarshal(rec.Body.Bytes(), &JWTTokens)
			assert.Nil(t, err)
			assert.NotNil(t, JWTTokens.AccessToken)
			assert.NotNil(t, JWTTokens.RefreshToken)
			_, err = helpers.GetClaims(*JWTTokens.AccessToken)
			assert.Nil(t, err)
			newAccessToken, err := helpers.RefreshToken(*JWTTokens.RefreshToken)
			assert.NotEqual(t, "", newAccessToken)
			assert.Nil(t, err)
		case http.StatusForbidden:
			var msg generated.MessageResponse
			err = json.Unmarshal(rec.Body.Bytes(), &msg)
			assert.Nil(t, err)
			assert.NotZero(t, msg)
			assert.Contains(t, helpers.DRForbidden, msg.Message)
		case http.StatusInternalServerError:
			bodyResp := rec.Body.String()
			t.Fatal(bodyResp)
		default:
			t.Fatal("not handled response")
		}
	}
}

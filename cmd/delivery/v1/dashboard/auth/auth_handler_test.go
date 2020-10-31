package dashboard

import (
	"arka/cmd/entity"
	"arka/cmd/mocks"
	"arka/pkg/server"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var configOption = server.Options{
	ListenAddress: ":4000",
	MaxConnection: 100,
	ReadTimeout:   time.Duration(30 * time.Second),
	WriteTimeout:  time.Duration(30 * time.Second),
	Timeout:       time.Duration(30 * time.Second),
}

var mockService = new(mocks.AuthService)

func TestLogin(t *testing.T) {
	mockUser := entity.User{
		ID:          "ini id",
		Username:    "user1",
		Email:       "coba@mail.com",
		FirstName:   "first_name",
		LastName:    "last_name",
		Password:    "okokok",
		RoleID:      "ok",
		PhoneNumber: "0813456789",
		CreatedAt:   time.Now(),
		CreatedBy:   "o",
		UpdatedAt:   nil,
		UpdatedBy:   nil,
	}
	loginPayload := &entity.Authorization{
		User:    mockUser,
		Refresh: "refresh",
		Token:   "token",
	}

	t.Run("success", func(t *testing.T) {
		mockService.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(loginPayload, nil).Once()
		bodyLogin := map[string]interface{}{
			"email":       "coba@mail.com",
			"phoneNumber": "0813456789",
			"password":    "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("success-with-email", func(t *testing.T) {
		mockService.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(loginPayload, nil).Once()
		bodyLogin := map[string]interface{}{
			"email":    "coba@mail.com",
			"password": "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("success-with-phone-number", func(t *testing.T) {
		mockService.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(loginPayload, nil).Once()
		bodyLogin := map[string]interface{}{
			"phone_number": "0813456789",
			"password":     "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New("error")).Once()
		bodyLogin := map[string]interface{}{
			"email":       "coba@mail.com",
			"phoneNumber": "0813456789",
			"password":    "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.NotEqual(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error-email", func(t *testing.T) {
		bodyLogin := map[string]interface{}{
			"email":    "cobamail.com",
			"password": "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error-phone-number", func(t *testing.T) {
		bodyLogin := map[string]interface{}{
			// "pho":    "081292929323",
			"password": "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("error-decode", func(t *testing.T) {
		bodyLogin := map[string]interface{}{
			// "pho":    "081292929323",
			// "password": "okokok",
		}

		body, _ := json.Marshal(bodyLogin)
		server := server.New(&configOption)
		router := server.Router()

		authHandler := NewAuthDashboard(mockService)

		router.POST("/login", authHandler.Login)

		req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(body)))

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertExpectations(t)
	})

}

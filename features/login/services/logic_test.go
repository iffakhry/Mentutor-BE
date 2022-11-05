package services

import (
	"testing"

	"be12/mentutor/features/login"
	mocks "be12/mentutor/mocks/features/login"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewDataInterface(t)
	t.Run("Success login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(login.Core{ID: 2, Email: "fatur69@gmail.com", Password: "$2a$10$fk68mY5i/hFtQLhtaLS6L.LVNyIWoCgQ3CUdD2ySbYwHWbQulzUUu"}, nil).Once()
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "Fatur123$"}
		res, _, err := srv.Login(input)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed cointain space", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fatur6 9@gmail.com", Password: "Fatur123$"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "contain space")
		repo.AssertExpectations(t)
	})
	t.Run("Failed length not valid", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fat", Password: "Fatur123$"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "length not valid")
		repo.AssertExpectations(t)
	})
	t.Run("Failed not contain @", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fataasdfasf", Password: "Fatur123$"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not contain (@) or (.)")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "fatur"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "string not as expected")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "FATUR"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "string not as expected")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "Fatur"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "string not as expected")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "Fatur123"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "string not as expected")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "Ft3$"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "string too short or too long")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(login.Core{Password: "$2a$10$fk68mY5i/hFtQLhtaLS6L.LVNyIWoCgQ3CUdD2ySbYwHWbQulzUUu"}, nil).Once()
		srv := New(repo)
		input := login.Core{Email: "fatur699@gmail.com", Password: "Fatursdf123$"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong email or password")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(login.Core{ID: 23, Password: "Admin123$", Role: "admin"}, nil).Once()
		srv := New(repo)
		input := login.Core{Email: "admin.fatur@gmail.com", Password: "Admasi123$"}
		res, _ ,err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong email or password")
		repo.AssertExpectations(t)
	})
	t.Run("Failed string not as expected", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(login.Core{ID: 23, Password: "$2a$10$fk68mY5i/hFtQLhtaLS6L.LVNyIWoCgQ3CUdD2ySbYwHWbQulzUUu", Role: "mentee"}, nil).Once()
		srv := New(repo)
		input := login.Core{Email: "fatur69@gmail.com", Password: "Admasi123$"}
		res, _ ,err := srv.Login(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong email or password")
		repo.AssertExpectations(t)
	})
}

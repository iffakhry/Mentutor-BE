package services

import (
	"be12/mentutor/features/mentor"
	mocks "be12/mentutor/mocks/features/mentor"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFailedUpdateUser(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Length name not valid", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).
		Return(mentor.UserCore{
			IdUser: 1,
			Name: "Nur Fatchurohman",
			Email: "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role: "mentor",
		}, errors.New("user not found")).Once()

		input := mentor.UserCore{
			IdUser: 1,
			Name: "58hgdghdghhdG",
			Email: "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role: "mentee",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
func TestSuccessUpdateUser(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Success update user mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).
		Return(mentor.UserCore{
				IdUser: 1,
				Name: "Nur Fatchurohman",
				Email: "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role: "mentee",
			}, nil).Once()
		repo.On("EditProfileMentee", mock.Anything).
		Return(mentor.UserCore{
			IdUser: 1,
			Name: "Nur Fatchurohman",
			Email: "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role: "mentee",
		}, nil).Once()
		input := mentor.UserCore{
			IdUser: 1,
				Name: "Nur Fatchurohman",
				Email: "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role: "mentee",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentee")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Success update user mentor", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).
		Return(mentor.UserCore{
				IdUser: 1000,
				Name: "Nur Fatchurohman",
				Email: "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role: "mentor",
			}, nil).Once()
		repo.On("EditProfileMentor", mock.Anything).
		Return(mentor.UserCore{
			IdUser: 1000,
			Name: "Nur Fatchurohman",
			Email: "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role: "mentor",
		}, nil).Once()
		input := mentor.UserCore{
			IdUser: 1000,
				Name: "Nur Fatchurohman",
				Email: "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role: "mentor",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed get user mentor", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).
		Return(mentor.UserCore{}, errors.New("user not found")).Once()

		input := mentor.UserCore{
			IdUser: 1000,
				Name: "Nur Fatchurohman",
				Email: "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role: "mentor",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed get user mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).
		Return(mentor.UserCore{}, errors.New("user not found")).Once()

		input := mentor.UserCore{
			IdUser: 1,
				Name: "Nur Fatchurohman",
				Email: "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role: "mentee",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}


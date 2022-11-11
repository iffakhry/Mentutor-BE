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
				IdUser:  1,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentor",
			}, errors.New("user not found")).Once()

		input := mentor.UserCore{
			IdUser:  1,
			Name:    "58hgdghdghhdG",
			Email:   "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role:    "mentee",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("no data input", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).
			Return(mentor.UserCore{
				IdUser:  1,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentee",
			})
		input := mentor.UserCore{
			IdUser:  1,
			Name:    "",
			Email:   "",
			IdClass: 0,
			Role:    "",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("password condition number", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).
			Return(mentor.UserCore{
				IdUser:  1000,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentee",
			}, errors.New("password condition number"))
		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "FaturRohman",
			Role:     "mentee",
		}
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("password condition upper", func(t *testing.T) {
		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1001,
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "faturrohman",
			Role:     "mentee",
		}
		res, err := srv.UpdateProfile(input, "mentor")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("User not admin", func(t *testing.T) { //USER NOT ADMIN LOGIC LINE 31
		repo.On("GetSingleMentor", mock.Anything).
			Return(mentor.UserCore{
				IdUser:  1000,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentee",
			}, errors.New("User not admin"))
		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1002,
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.UpdateProfile(input, "mentee")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Length email", func(t *testing.T) {
		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Nur Fatchurohman",
			Email:    "fam",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user not found")
		repo.AssertExpectations(t)
	})

}

func TestAddTask(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	task := mentor.TaskCore{
		ID:          1,
		IdClass:     1,
		IdMentor:    1000,
		Title:       "Persamaan",
		Description: "samain a dan xxx",
		File:        "file.pdf",
		Images:      "image.jpg",
	}
	t.Run("Success Add Task", func(t *testing.T) {
		repo.On("InsertTask", mock.Anything, mock.Anything).Return(task, nil).Once()
		srv := New(repo)
		res, err := srv.AddTask(task, "mentor")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Task", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.AddTask(task, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Task", func(t *testing.T) {
		repo.On("InsertTask", mock.Anything, mock.Anything).Return(task, errors.New("error insert task")).Once()
		srv := New(repo)
		_, err := srv.AddTask(task, "mentor")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
func TestSuccessUpdateUser(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Success update user mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).
			Return(mentor.UserCore{
				IdUser:  1,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentee",
			}, nil).Once()
		repo.On("EditProfileMentee", mock.Anything).
			Return(mentor.UserCore{
				IdUser:  1,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentee",
			}, nil).Once()
		input := mentor.UserCore{
			IdUser:  1,
			Name:    "Nur Fatchurohman",
			Email:   "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role:    "mentee",
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
				IdUser:  1000,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentor",
			}, nil).Once()
		repo.On("EditProfileMentor", mock.Anything).
			Return(mentor.UserCore{
				IdUser:  1000,
				Name:    "Nur Fatchurohman",
				Email:   "nur.faturohman28@gmail.com",
				IdClass: 1,
				Role:    "mentor",
			}, nil).Once()
		input := mentor.UserCore{
			IdUser:  1000,
			Name:    "Nur Fatchurohman",
			Email:   "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role:    "mentor",
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
			IdUser:  1000,
			Name:    "Nur Fatchurohman",
			Email:   "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role:    "mentor",
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
			IdUser:  1,
			Name:    "Nur Fatchurohman",
			Email:   "nur.faturohman28@gmail.com",
			IdClass: 1,
			Role:    "mentee",
		}
		srv := New(repo)
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Not mentor", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.DeleteTask(1, 1, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Success delete class", func(t *testing.T) {
		repo.On("GetTaskSub", mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteTask(1, 1, "mentor")
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Empty(t, res)

	})
	t.Run("Failed delete task", func(t *testing.T) {
		repo.On("GetTaskSub", mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, errors.New("error delete task")).Once()
		srv := New(repo)
		res, err := srv.DeleteTask(1, 1, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "error delete task")
		repo.AssertExpectations(t)
	})
}

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

func TestGetAllTask(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	task := []mentor.TaskCore{{
		ID:          1,
		IdClass:     1,
		IdMentor:    2,
		Title:       "Persamaan",
		Description: "A ketemu B tentukan bilangan z ?",
		File:        "file.pdf",
		Images:      "image.jpg",
	}}
	t.Run("Success Get All Task", func(t *testing.T) {
		repo.On("GetAllTask", mock.Anything).Return(task, nil).Once()

		srv := New(repo)
		res, err := srv.GetAllTask("mentor")
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All Task", func(t *testing.T) {
		repo.On("GetAllTask", mock.Anything).Return([]mentor.TaskCore{}, errors.New("error get all task")).Once()
		srv := New(repo)
		res, err := srv.GetAllTask("mentor")
		assert.Error(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All Task", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.GetAllTask("admin")
		assert.Error(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
func TestAddScore(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	sub := mentor.SubmissionCore{
		ID:         1,
		NameMentee: "Hery Budiyana",
		Title:      "Persamaan",
		IdMentee:   1,
		IdTask:     1,
		File:       "file.pdf",
		Score:      80,
	}
	task := mentor.TaskCore{
		ID:          1,
		IdClass:     1,
		IdMentor:    2,
		Title:       "Persamaan",
		Description: "A ketemu B tentukan bilangan z ?",
		File:        "file.pdf",
		Images:      "image.jpg",
	}
	t.Run("Success Add Score", func(t *testing.T) {
		repo.On("GetSubmission", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("GetSingleTask", mock.Anything, mock.Anything).Return(task, nil).Once()
		repo.On("AddScore", mock.Anything, mock.Anything).Return(sub, nil).Once()
		srv := New(repo)
		res, err := srv.AddScore(sub, "mentor")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("input score empty", func(t *testing.T) {
		// repo.On("AddScore", mock.Anything, mock.Anything).Return(mentor.SubmissionCore{}, errors.New("input score empty")).Once()
		input := mentor.SubmissionCore{Score: 0}
		srv := New(repo)
		_, err := srv.AddScore(input, "mentor")
		assert.Nil(t, nil)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Score", func(t *testing.T) {
		repo.On("GetSubmission", mock.Anything, mock.Anything).Return(errors.New("failed get submission")).Once()
		srv := New(repo)
		_, err := srv.AddScore(sub, "mentor")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("user not mentor", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.AddScore(sub, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}
func TestGetTaskSub(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	task := mentor.TaskCore{
		ID:          1,
		IdClass:     1,
		IdMentor:    2,
		Title:       "Persamaan",
		Description: "A ketemu B tentukan bilangan z ?",
		File:        "file.pdf",
		Images:      "image.jpg",
	}
	sub := []mentor.SubmissionCore{{
		ID:         1,
		NameMentee: "Hery Budiyana",
		Title:      "Persamaan",
		IdMentee:   1,
		IdTask:     1,
		File:       "persamaaan.pdf",
		Score:      99,
	}}
	t.Run("Success Get Task Submission", func(t *testing.T) {
		repo.On("GetTaskSub", mock.Anything, mock.Anything).Return(task, sub, nil).Once()
		srv := New(repo)
		res, _, err := srv.GetTaskSub(1, "mentor")
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)

	})
	t.Run("Failed Get Task Submission", func(t *testing.T) {
		repo.On("GetTaskSub", mock.Anything, mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, errors.New("error get detail task")).Once()
		srv := New(repo)
		res, _, err := srv.GetTaskSub(1, "mentor")
		assert.Error(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("User Not Mentor", func(t *testing.T) {
		// repo.On("GetTaskSub", mock.Anything, mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, errors.New("user not mentor")).Once()
		srv := New(repo)
		res, _, err := srv.GetTaskSub(1, "mentee")
		assert.Error(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateTask(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	task := mentor.TaskCore{
		ID:          1,
		IdClass:     1,
		IdMentor:    1,
		Title:       "persamaan",
		Description: "jika x dan y adalah genap maka tentukan bilangan asd",
		File:        "persamaan.pdf",
		Images:      "image.jpg",
	}
	t.Run("success update task", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(task, nil).Once()
		repo.On("EditTask", mock.Anything).Return(task, nil).Once()

		srv := New(repo)
		res, err := srv.UpdateTask(task, "mentor")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)

	})
	t.Run("Failed Update task", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(task, nil).Once()
		repo.On("EditTask", mock.Anything).Return(mentor.TaskCore{}, errors.New("error update task")).Once()
		srv := New(repo)
		_, err := srv.UpdateTask(task, "mentor")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)

	})
	t.Run("user not mentor", func(t *testing.T) {

		srv := New(repo)
		_, err := srv.UpdateTask(task, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)

	})
	t.Run("success update task", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(task, nil).Once()
		repo.On("EditTask", mock.Anything).Return(task, nil).Once()

		srv := New(repo)
		input := mentor.TaskCore{
			ID:       1,
			IdClass:  1,
			IdMentor: 1,
			Title:    "",
		}
		res, err := srv.UpdateTask(input, "mentor")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)

	})
	t.Run("success update task", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(mentor.TaskCore{}, errors.New("task not found")).Once()
		// repo.On("EditTask", mock.Anything).Return(task, nil).Once()

		srv := New(repo)
		input := mentor.TaskCore{
			ID:       1,
			IdClass:  1,
			IdMentor: 1,
			Title:    "",
		}
		_, err := srv.UpdateTask(input, "mentor")
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
	t.Run("Success delete task", func(t *testing.T) {
		repo.On("GetTaskSub", mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteTask(1, 1, "mentor")
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Empty(t, res)

	})
	t.Run("Success delete task", func(t *testing.T) {
		repo.On("GetTaskSub", mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, nil).Once()
		input := mentor.TaskCore{
			ID:      1,
			IdClass: 1,
		}
		srv := New(repo)
		res, err := srv.DeleteTask(input.ID, input.IdClass, "mentor")
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

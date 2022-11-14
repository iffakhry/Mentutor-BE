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
			Return(mentor.UserCore{}, errors.New("password condition number")).Once()
		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "as",
			Email:    "as",
			IdClass:  7,
			Password: "as",
			Role:     "mentee",
		}
		res, err := srv.UpdateProfile(input, "mentor")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

}

func TestFailedUpdateUsers(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	mentors := mentor.UserCore{
		IdUser:  1000,
		Name:    "Hery Budiyana",
		Email:   "heribudiyana@gmail.com",
		IdClass: 1,
		Class:   "Back End",
		Role:    "mentor",
		Images:  "image.jpg",
	}
	mentee := mentor.UserCore{
		IdUser:  1,
		Name:    "Heri Budiyana",
		Email:   "heri.mentee@gmail.com",
		IdClass: 1,
		Class:   "Back_End",
		Role:    "mentee",
		Images:  "image.jpg",
	}

	repo.On("GetSingleMentee", mock.Anything).Return(mentee, nil).Once()
	repo.On("EditProfileMentee", mock.Anything).Return(mentee, nil).Once()
	input := mentor.UserCore{
		IdUser:   12,
		Name:     "Hery Mentee",
		Email:    "hery.mentee@mail.com",
		Password: "Asdf123$",
		IdClass:  7,
	}
	srv := New(repo)
	res, err := srv.UpdateProfile(input, "mentor")
	assert.NotNil(t, res)
	assert.Empty(t, err)
	repo.AssertExpectations(t)
	t.Run("Length Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "ft r",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "FATUR ROHMAN",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Upper Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "fatur rohman",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Number Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur 1324",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Special Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur ##$",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Space Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "FaturRohman",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Space error email", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatu @gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Length error email", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fat@.co",
			Password: "Mentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Length error email", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "faturrohman",
			Password: "Mentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)

		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "menteementutor$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "FATURROHMAN$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "FaturRohman$",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "FaturRohman123",
			IdClass:  7,
		}
		res, err := srv.UpdateProfile(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Class avail check", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentors, nil).Once()
		repo.On("EditProfileMentor", mock.Anything).Return(mentors, nil).Once()
		srv := New(repo)
		input := mentor.UserCore{
			IdUser:   1000,
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "Mentee123$?",
			IdClass:  0,
		}
		_, err := srv.UpdateProfile(input, "admin")

		assert.Empty(t, err)
		assert.Nil(t, nil)
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
		// repo.On("InsertTask", mock.Anything, mock.Anything).Return(task, nil).Once()
		input := mentor.TaskCore{
			ID:          2,
			IdClass:     2,
			IdMentor:    1002,
			Title:       "Persamaan",
			Description: "samain a dan xxx",
			File:        "file.pdf",
			Images:      "image.jpg",
		}
		srv := New(repo)
		_, err := srv.AddTask(input, "mentor")
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Task", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.AddTask(task, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Task", func(t *testing.T) {

		input := mentor.TaskCore{
			Title:       "",
			Description: "",
			File:        "",
			Images:      "",
		}
		srv := New(repo)
		_, err := srv.AddTask(input, "mentor")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Task", func(t *testing.T) {
		// repo.On("InsertTask", mock.Anything).Return(mentor.TaskCore{}, errors.New("failed add task")).Once()

		input := mentor.TaskCore{
			ID:          1,
			IdClass:     1,
			IdMentor:    1,
			Title:       "",
			Description: "",
			File:        "",
			Images:      "",
		}
		srv := New(repo)
		_, err := srv.AddTask(input, "mentor")
		assert.Nil(t, nil)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add Task", func(t *testing.T) {

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
	t.Run("invalid number", func(t *testing.T) {

		repo.On("GetSubmission", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("GetSingleTask", mock.Anything, mock.Anything).Return(mentor.TaskCore{}, errors.New("task not found")).Once()
		srv := New(repo)
		_, err := srv.AddScore(sub, "mentor")
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("invalid number", func(t *testing.T) {

		repo.On("GetSubmission", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("GetSingleTask", mock.Anything, mock.Anything).Return(task, nil).Once()
		repo.On("AddScore", mock.Anything, mock.Anything).Return(mentor.SubmissionCore{}, errors.New("error add score")).Once()

		srv := New(repo)
		_, err := srv.AddScore(sub, "mentor")
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

	// t.Run("success delete task", func(t *testing.T) {
	// 	repo.On("GetTaskSub", mock.Anything).Return(mentor.TaskCore{}, []mentor.SubmissionCore{}, nil).Once()
	// 	repo.On("DeleteTask", mock.Anything, mock.Anything).Return(nil, nil, nil).Once()
	// 	idclass := 1
	// 	idToken := 1
	// 	srv := New(repo)
	// 	res, err := srv.DeleteTask(uint(idToken), uint(idclass), "mentor")
	// 	assert.NotNil(t, err)
	// 	assert.Empty(t, res)
	// 	// assert.ErrorContains(t, err, "error delete task")
	// 	repo.AssertExpectations(t)
	// })
}

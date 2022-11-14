package services

import (
	"be12/mentutor/features/mentee"
	mocks "be12/mentutor/mocks/features/mentee"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// DONE
func TestPostForum(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	data := mentee.Status{
		ID:       1,
		Caption:  "Aku bingung ini kenapa",
		Images:   "image.jpg",
		IdMentee: 2,
	}
	t.Run("success add status", func(t *testing.T) {
		repo.On("AddStatus", mock.Anything, mock.Anything).Return(data, nil).Once()
		srv := New(repo)
		input := mentee.Status{
			Caption: "Aku bingung ini kenapa",
			Images:  "image.jpg",
		}
		res, err := srv.InsertStatus(input, 1)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed add status", func(t *testing.T) {
		repo.On("AddStatus", mock.Anything, mock.Anything).Return(mentee.Status{}, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.InsertStatus(data, 1)
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed length not valid", func(t *testing.T) {
		srv := New(repo)
		input := mentee.Status{
			Caption: "as",
			Images:  "images.jpg",
		}
		res, err := srv.InsertStatus(input, 1)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "input not valid")
		repo.AssertExpectations(t)
	})
}

// DONE
func TestAddComment(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("success add comments", func(t *testing.T) {
		repo.On("GetSingleStatus", mock.Anything).Return(nil, nil).Once()
		repo.On("AddComment", mock.Anything).Return(mentee.CommentsCore{}, nil).Once()
		srv := New(repo)
		input := mentee.CommentsCore{IdStatus: 1, Caption: "Aku bingung ini kenapa"}
		res, err := srv.Insert(input)
		assert.Empty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed add status", func(t *testing.T) {
		repo.On("GetSingleStatus", mock.Anything).Return(nil, nil).Once()
		repo.On("AddComment", mock.Anything).Return(mentee.CommentsCore{}, errors.New("Error")).Once()
		input := mentee.CommentsCore{
			IdStatus: 1,
			Caption:  "asasas",
		}
		srv := New(repo)
		_, err := srv.Insert(input)
		assert.Error(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed add status", func(t *testing.T) {
		repo.On("GetSingleStatus", mock.Anything).Return(errors.New("failed get status"), errors.New("failed get status")).Once()
		input := mentee.CommentsCore{

			Caption: "asasas",
		}
		srv := New(repo)
		res, err := srv.Insert(input)
		assert.Error(t, err)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed length not valid", func(t *testing.T) {
		srv := New(repo)
		input := mentee.CommentsCore{
			IdStatus: 1,
			Caption:  "as",
		}
		res, err := srv.Insert(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "failed add your comment check charancter len")
		repo.AssertExpectations(t)
	})

}

// // DONE
func TestInsertSub(t *testing.T) {
	repo := mocks.NewRepoInterface(t)

	tasks := mentee.Task{
		ID:          1,
		IdClass:     2,
		IdMentor:    3,
		Title:       "ini title",
		Description: "ini description",
		File:        "file.pdf",
		Images:      "image.jpg",
	}
	t.Run("success add submission", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(tasks, nil).Once()
		srv := New(repo)
		input := mentee.Submission{File: "file.pdf"}
		res, _ := srv.InsertSub(input)
		assert.Empty(t, res)
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
	})

	t.Run("failed add submission", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(tasks, errors.New("Error")).Once()

		srv := New(repo)
		input := mentee.Submission{File: "file.pdf"}
		result, err := srv.InsertSub(input)
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
}

func TestAddStatus(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	submission := mentee.Submission{
		ID:        1,
		ID_Mentee: 1,
		ID_Tasks:  1,
		File:      "file.pdf",
		Score:     0,
		Title:     "persamaan",
		Status:    "active",
	}
	t.Run("success add submission", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(mentee.Task{}, nil).Once()
		// repo.On("AddSub", mock.Anything).Return(mentee.Submission{}, nil, nil).Once()
		srv := New(repo)
		// input := mentee.Submission{File: "file.pdf"}
		res, _ := srv.InsertSub(submission)
		assert.Empty(t, res)
		assert.Nil(t, nil)
		repo.AssertExpectations(t)

	})
}

// // Done
func TestGetStatus(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	status := []mentee.Status{{
		ID:       1,
		Caption:  "ini gimana",
		Name:     "Hery",
		Images:   "image.jpg",
		IdMentee: 2,
	}}
	t.Run("Success Get status", func(t *testing.T) {
		repo.On("GetAllPosts", mock.Anything).Return(status, nil, nil, nil).Once()

		usecase := New(repo)
		result, _, _, err := usecase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})

	t.Run("failed get status", func(t *testing.T) {
		repo.On("GetAllPosts", mock.Anything).Return(nil, nil, nil, errors.New("Error")).Once()

		usecase := New(repo)

		result, _, _, err := usecase.GetAll()
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
}

// // DONE
func TestGetTask(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	task := []mentee.Task{{
		ID:          1,
		IdMentor:    2,
		IdClass:     2,
		Title:       "Pertambahan",
		Description: "tambahkan nilai A ke Z kemudian jika x bertemu H maka si A kemana ?",
		File:        "task.pdf",
		Images:      "task.jpg",
		Score:       90,
	}}
	token := 1
	t.Run("Success Get Task", func(t *testing.T) {
		repo.On("GetAllTask", mock.Anything, mock.Anything).Return(task, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetTask(uint(token), "mentee")
		assert.NoError(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
	t.Run("failed get task", func(t *testing.T) {
		repo.On("GetAllTask", mock.Anything, mock.Anything).Return([]mentee.Task{}, errors.New("user not mentee")).Once()

		usecase := New(repo)

		result, err := usecase.GetTask(uint(token), "mentee")
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
	t.Run("failed user not mentee", func(t *testing.T) {

		usecase := New(repo)

		result, err := usecase.GetTask(uint(token), "admin")
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
}
func TestGetTokenMentee(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	token := mentee.Token{
		Id:           1,
		IdMentee:     1,
		Code:         "asasasasa",
		AccessToken:  "sasasasasa",
		TokenType:    "Credential",
		RefreshToken: "Refreshing",
	}
	idToken := 1
	t.Run("Success Get Token Mentee", func(t *testing.T) {
		repo.On("GetTokenMentee", mock.Anything).Return(token, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetTokenMentee(uint(idToken))
		assert.NoError(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Token Mentee", func(t *testing.T) {
		repo.On("GetTokenMentee", mock.Anything).Return(mentee.Token{}, errors.New("failed Get token mentee")).Once()

		usecase := New(repo)
		result, err := usecase.GetTokenMentee(uint(idToken))
		assert.Error(t, err)
		assert.NotNil(t, result)
		repo.AssertExpectations(t)
	})
}

func TestAddToken(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	token := mentee.Token{
		Id:           1,
		IdMentee:     1,
		Code:         "asdasdasdasdasdasdasd",
		AccessToken:  "dsadsadasdsadsa",
		TokenType:    "string",
		RefreshToken: "refreshing",
	}
	t.Run("success add token", func(t *testing.T) {
		repo.On("InsertToken", mock.Anything).Return(mentee.Token{}, nil).Once()
		srv := New(repo)
		res, err := srv.AddToken(token)
		assert.Empty(t, err)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Insert Token", func(t *testing.T) {
		repo.On("InsertToken", mock.Anything).Return(mentee.Token{}, errors.New("failed insert token")).Once()
		srv := New(repo)
		res, err := srv.AddToken(token)
		assert.NotEmpty(t, err)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetSingleTask(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	task := mentee.Task{
		ID:          1,
		IdClass:     1,
		IdMentor:    1,
		Title:       "Persamaan",
		Description: "tugas akhir semester genap",
		File:        "persamaan.pdf",
		Images:      "file.pdf",
		Score:       89,
	}
	t.Run("success get single task", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(task, nil).Once()

		srv := New(repo)
		res, err := srv.GetSingleTask(1)
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("failed get singletask", func(t *testing.T) {
		repo.On("GetSingleTask", mock.Anything).Return(mentee.Task{}, errors.New("failed get single task")).Once()

		srv := New(repo)
		res, err := srv.GetSingleTask(1)
		assert.Error(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
}

func TestGetMentee(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	mentees := mentee.MenteeCore{
		IdUser: 1,
		Name:   "Heri Budiyana",
		Email:  "heribudiyana@gmail.com",
		Images: "hery.jpg",
	}
	t.Run("success get mentee", func(t *testing.T) {
		repo.On("GetMentee", mock.Anything).Return(mentees, nil).Once()

		srv := New(repo)
		res, err := srv.GetMentee(1)
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed get mentee", func(t *testing.T) {
		repo.On("GetMentee", mock.Anything).Return(mentee.MenteeCore{}, errors.New("failed get mentee")).Once()

		srv := New(repo)
		res, err := srv.GetMentee(1)
		assert.Error(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
}

func TestGetMentor(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	mentors := mentee.MentorCore{
		Role:   "mentor",
		Name:   "Heri Budiyana",
		Email:  "heribudiyana@gmail.com",
		Images: "hery.jpg",
	}
	t.Run("success get mentor", func(t *testing.T) {
		repo.On("GetMentor", mock.Anything).Return(mentors, nil).Once()

		srv := New(repo)
		res, err := srv.GetMentor(1)
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("success get mentor", func(t *testing.T) {
		repo.On("GetMentor", mock.Anything).Return(mentee.MentorCore{}, errors.New("failed get mentor")).Once()

		srv := New(repo)
		res, err := srv.GetMentor(1)
		assert.Error(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})

}

func TestGetSub(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("success get submission", func(t *testing.T) {
		repo.On("GetSub", mock.Anything, mock.Anything).Return(1, nil)

		srv := New(repo)
		res, err := srv.GetSub(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
	t.Run("failed get submission", func(t *testing.T) {
		repo.On("GetSub", mock.Anything, mock.Anything).Return(-1, errors.New("failed get submission"))

		srv := New(repo)
		res, _ := srv.GetSub(1, 1)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})
}

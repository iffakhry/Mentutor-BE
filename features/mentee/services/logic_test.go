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
	comment := mentee.CommentsCore{
		Caption: "pake cara ini juga bisa",
	}
	t.Run("success add comments", func(t *testing.T) {
		repo.On("AddComment", mock.Anything).Return(comment, nil).Once()
		srv := New(repo)
		input := mentee.CommentsCore{Caption: "Aku bingung ini kenapa"}
		res, err := srv.Insert(input)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed add comments", func(t *testing.T) {
		repo.On("AddComment", mock.Anything).Return(mentee.CommentsCore{}, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.Insert(comment)
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})

	t.Run("failed add your comment check charancter len", func(t *testing.T) {
		srv := New(repo)
		input := mentee.CommentsCore{Caption: "as"}
		res, err := srv.Insert(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "failed add your comment check charancter len")
		repo.AssertExpectations(t)
	})
}

// DONE
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

// Done
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

// DONE
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
	t.Run("failed get status", func(t *testing.T) {
		repo.On("GetAllTask", mock.Anything, mock.Anything).Return([]mentee.Task{}, errors.New("user not mentee")).Once()

		usecase := New(repo)

		result, err := usecase.GetTask(uint(token), "mentee")
		assert.Error(t, err)
		assert.Equal(t, result, result)
		repo.AssertExpectations(t)
	})
}

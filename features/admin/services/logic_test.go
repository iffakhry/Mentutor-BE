package services

import (
	"be12/mentutor/features/admin"
	mocks "be12/mentutor/mocks/features/admin"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// DONE
func TestAddUser(t *testing.T) {

	repo := mocks.NewRepoInterface(t)
	t.Run("Success Register", func(t *testing.T) {
		repo.On("InsertMentee", mock.Anything).
			Return(admin.UserCore{
				Name:     "Nur Fatchurohman",
				Email:    "fatur@gmail.com",
				IdClass:  7,
				Class:    "Back End",
				Password: "Fatur123$",
				Role:     "mentee"}, nil).
			Once()
		repo.On("GetClass", mock.Anything).
			Return(admin.ClassCore{
				IdClass:   7,
				ClassName: "Back End",
				Status:    "active",
			}, nil).Once()
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123?$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("User not admin", func(t *testing.T) { //USER NOT ADMIN LOGIC LINE 31
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "mentee")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Length email", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fam",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "length email not valid")
		repo.AssertExpectations(t)
	})
	t.Run("@ email", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fasdfasasgasam",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("space email", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur @gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("length name", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nun",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("name condition upper", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "NUR FATCHUROHMAN",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	// DONE
	t.Run("name condition lower", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "nur fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	}) // BATAS
	t.Run("name condition space", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "nurfatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("name condition number", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "nuR fatchurohma25n",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("name condition schar", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur fatchurohma@#n",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("password condition number", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "FaturRohman",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("name condition schar", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur fatchu3rohma@#n",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "F2das3",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("password condition upper", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "faturrohman",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("password condition lower", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "FATURROHMAN",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("password condition schar", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "FaturRohman123",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed insert mentee", func(t *testing.T) {
		repo.On("InsertMentee", mock.Anything).
			Return(admin.UserCore{}, errors.New("failed insert")).
			Once()
		repo.On("GetClass", mock.Anything).
			Return(admin.ClassCore{
				IdClass:   7,
				ClassName: "Back End",
				Status:    "active",
			}, nil).Once()
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("password condition lenght", func(t *testing.T) {
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohmann",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Ftr23?",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Class avail false", func(t *testing.T) {
		repo.On("GetClass", mock.Anything).
			Return(admin.ClassCore{}, errors.New("input class not valid")).Once()
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  54,
			Password: "Fatur123$",
			Role:     "mentee",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "input class not valid")
		repo.AssertExpectations(t)
	})

	t.Run("Failed insert mentor", func(t *testing.T) {
		repo.On("InsertMentor", mock.Anything).
			Return(admin.UserCore{
				IdUser: 1,
				Name:   "Nur Fatchurohman",
				Email:  "fatur@gmail.com",
			}, nil).
			Once()
		repo.On("GetClass", mock.Anything).
			Return(admin.ClassCore{
				IdClass:   7,
				ClassName: "Back End",
				Status:    "active",
			}, nil).Once()
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentor",
		}
		res, err := srv.AddUser(input, "admin")
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed insert mentor", func(t *testing.T) {
		repo.On("InsertMentor", mock.Anything).
			Return(admin.UserCore{}, errors.New("failed insert")).
			Once()
		repo.On("GetClass", mock.Anything).
			Return(admin.ClassCore{
				IdClass:   7,
				ClassName: "Back End",
				Status:    "active",
			}, nil).Once()
		srv := New(repo)
		input := admin.UserCore{
			Name:     "Nur Fatchurohman",
			Email:    "fatur@gmail.com",
			IdClass:  7,
			Password: "Fatur123$",
			Role:     "mentor",
		}
		res, err := srv.AddUser(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestSuccessGetAllUser(t *testing.T) {

	repo := mocks.NewRepoInterface(t)
	t.Run("User not admin", func(t *testing.T) {
		srv := New(repo)
		resMentee, resMentor, err := srv.GetAllUser("mentee")
		assert.Empty(t, resMentee)
		assert.Empty(t, resMentor)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Success Get All User", func(t *testing.T) {
		repo.On("GetAllUser", mock.Anything).Return(
			[]admin.UserCore{
				admin.UserCore{
					IdUser: 1,
					Name:   "Fatur",
				},
			},
			[]admin.UserCore{
				admin.UserCore{
					IdUser: 2,
					Name:   "Rohman",
				},
			}, nil,
		)
		srv := New(repo)
		resMentee, resMentor, err := srv.GetAllUser("admin")
		assert.NotEmpty(t, resMentee)
		assert.NotEmpty(t, resMentor)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestFailedGetAllUser(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Failed get all user", func(t *testing.T) {
		t.Run("Failed get all admin", func(t *testing.T) {
			repo.On("GetAllUser").Return([]admin.UserCore{}, []admin.UserCore{}, errors.New("user not found"))
			srv := New(repo)
			resMentee, resMentor, err := srv.GetAllUser("admin")
			assert.Empty(t, resMentee)
			assert.Empty(t, resMentor)
			assert.NotNil(t, err)
			repo.AssertExpectations(t)
		})
	})
}

func TestSuccessAddNewClass(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	class := admin.ClassCore{ClassName: "Front End", Status: "active"}

	t.Run("User not admin", func(t *testing.T) {
		srv := New(repo)
		input := admin.ClassCore{ClassName: "backend"}
		res, err := srv.AddNewClass(input, "mentee")
		assert.NotNil(t, res, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed add new class", func(t *testing.T) {
		repo.On("InsertNewClass", mock.Anything, mock.Anything).Return(admin.ClassCore{}, errors.New("input not valid")).Once()
		srv := New(repo)
		_, err := srv.AddNewClass(class, "admin")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("input not empty", func(t *testing.T) {
		repo.On("InsertNewClass", mock.Anything, mock.Anything).Return(admin.ClassCore{}, errors.New("input not valid")).Once()
		srv := New(repo)
		res, _ := srv.AddNewClass(class, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, res, 0)
		repo.AssertExpectations(t)

	})
	t.Run("contain special character", func(t *testing.T) {
		srv := New(repo)
		input := admin.ClassCore{ClassName: ""}
		res, _ := srv.AddNewClass(input, "admin")
		assert.NotNil(t, res, 0)
		repo.AssertExpectations(t)

	})
	t.Run("length name not valid", func(t *testing.T) {
		repo.On("InsertNewClass", mock.Anything, mock.Anything).Return(admin.ClassCore{}, errors.New("input not valid")).Once()

		srv := New(repo)
		input := admin.ClassCore{
			IdClass:      1,
			Status:       "adctive",
			ClassName:    "Back End",
			TotalStudent: 20,
		}
		_, err := srv.AddNewClass(input, "admin")
		// assert.Empty(t, err)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "inpdut not valid")
		repo.AssertExpectations(t)
	})
	t.Run("input not valid", func(t *testing.T) {
		repo.On("InsertNewClass", mock.Anything, mock.Anything).Return(admin.ClassCore{IdClass: 0}, errors.New("input not valid")).Once()

		srv := New(repo)
		input := admin.ClassCore{
			IdClass:      0,
			Status:       "active",
			ClassName:    "Back End",
			TotalStudent: 20,
		}
		res, err := srv.AddNewClass(input, "admin")
		assert.Empty(t, res)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "inpdut not valid")
		repo.AssertExpectations(t)
	})
	t.Run("input not valid", func(t *testing.T) {
		srv := New(repo)
		input := admin.ClassCore{
			IdClass:   0,
			ClassName: "1",
		}
		res, err := srv.AddNewClass(input, "admin")
		assert.Empty(t, res, err)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "length name not valid")
		repo.AssertExpectations(t)
	})

}

func TestFailedAddNewClass(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Failed add new class", func(t *testing.T) {
		repo.On("InsertNewClass", mock.Anything, mock.Anything).Return(admin.ClassCore{}, errors.New("input not valid")).Once()
		srv := New(repo)
		input := admin.ClassCore{ClassName: "backend"}
		_, err := srv.AddNewClass(input, "admin")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestSuccessGetClass(t *testing.T) {

	repo := mocks.NewRepoInterface(t)
	t.Run("Success Get New Class", func(t *testing.T) {
		repo.On("GetAllClass", mock.Anything).Return([]admin.ClassCore{
			admin.ClassCore{
				IdClass:   1,
				ClassName: "Backend",
				Status:    "active",
			},
		}, nil)
		srv := New(repo)
		res, err := srv.GetAllClass("admin")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestFailedGetClass(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("User not admin", func(t *testing.T) {
		srv := New(repo)
		res, err := srv.GetAllClass("mentee")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Add New Class", func(t *testing.T) {
		repo.On("GetAllClass", mock.Anything).Return([]admin.ClassCore{}, errors.New("error get all class"))
		srv := New(repo)
		res, err := srv.GetAllClass("admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	class := admin.ClassCore{IdClass: 1, ClassName: "Back End", Status: "active", TotalStudent: 20}
	mentor := admin.UserCore{
		IdUser:   1,
		Name:     "Hery Budiyana",
		Email:    "hery@gmail.com",
		IdClass:  1,
		Class:    "Back End",
		Password: "Asdf123$",
		Role:     "mentor",
		Images:   "Hery.jpg"}
	t.Run("Not admin", func(t *testing.T) {
		// DONE
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetClass", mock.Anything).Return(class, nil).Once()
		repo.On("EditUserMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "mentee")
		assert.NotNil(t, res)
		assert.Empty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Success update mentor", func(t *testing.T) {
		repo.On("EditUserMentee", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetSingleMentee", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetClass", mock.Anything).Return(class, nil).Once()

		srv := New(repo)

		res, err := srv.UpdateUserAdmin(mentor, "admin")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)

	})

	t.Run("Length Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "ft r",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "FATUR ROHMAN",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Upper Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "fatur rohman",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Number Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur 1324",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Special Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur ##$",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Space Char Name", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "FaturRohman",
			Email:    "fatu@gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Space error email", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatu @gmail.com",
			Password: "MEentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Length error email", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fat@.co",
			Password: "Mentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Length error email", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "faturrohman",
			Password: "Mentee123$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)

		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "menteementutor$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "FATURROHMAN$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "FaturRohman$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "FaturRohman123",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Char error password", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Fatur Rohman",
			Email:    "fatur@gmail.com",
			Password: "Fan1$",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	//	DARI SINI
	t.Run("Class avail check", func(t *testing.T) {
		// repo.On("GetClass", mock.Anything).Return(admin.ClassCore{}, errors.New("class not found")).Once()
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()
		repo.On("EditUserMentor", mock.Anything).Return(mentor, nil).Once()
		// DONE
		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "Mentee123$?",
			IdClass:  0,
		}
		_, err := srv.UpdateUserAdmin(input, "admin")

		assert.Empty(t, err)
		assert.Nil(t, nil)
		repo.AssertExpectations(t)
	})
	t.Run("Failed edit mentor", func(t *testing.T) {
		// repo.On("EditUserMentor", mock.Anything).Return(admin.UserCore{}, errors.New("error edit mentor")).Once()
		repo.On("GetSingleMentor", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetClass", mock.Anything).Return(class, errors.New("class not found")).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1000,
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "Mentee123$?",
			IdClass:  7,
		} // DONE
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed edit mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetClass", mock.Anything).Return(class, errors.New("class not found")).Once()

		// repo.On("EditUserMentee", mock.Anything).Return(admin.UserCore{}, errors.New("error edit mentor")).Once()
		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1,
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "Mentee123$?",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Success update mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetClass", mock.Anything).Return(class, errors.New("class not found")).Once()

		// repo.On("EditUserMentee", mock.Anything).Return(admin.UserCore{
		// 	IdUser:   1,
		// 	Name:     "Mentee Admin",
		// 	Password: "Mentee123$",
		// 	IdClass:  7,
		// }, nil).Once()

		srv := New(repo)
		input := admin.UserCore{
			IdUser:   1,
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "Mentee123$?",
			IdClass:  7,
		}
		res, _ := srv.UpdateUserAdmin(input, "admin")
		assert.Nil(t, nil)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Success update mentee", func(t *testing.T) {
		// repo.On("EditUserMentee", mock.Anything).Return(mentor, errors.New("user not found")).Once()
		repo.On("GetSingleMentee", mock.Anything).Return(mentor, nil).Once()
		repo.On("GetClass", mock.Anything).Return(class, errors.New("class not found")).Once()

		srv := New(repo)
		input := admin.UserCore{
			Name:     "Mentee Admin",
			Email:    "fatu@gmail.com",
			Password: "Mentee123$?",
			IdClass:  7,
		}
		res, err := srv.UpdateUserAdmin(input, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Not Admin", func(t *testing.T) {
		srv := New(repo)
		err := srv.DeleteUser(1, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Success delete mentee", func(t *testing.T) {
		repo.On("DeleteUserMentee", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteUser(1, "admin")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed delete mentee", func(t *testing.T) {
		repo.On("DeleteUserMentee", mock.Anything).Return(errors.New("error delete user")).Once()
		srv := New(repo)
		err := srv.DeleteUser(1, "admin")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Success delete mentor", func(t *testing.T) {
		repo.On("DeleteUserMentor", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteUser(1000, "admin")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed delete mentee", func(t *testing.T) {
		repo.On("DeleteUserMentor", mock.Anything).Return(errors.New("error delete user")).Once()
		srv := New(repo)
		err := srv.DeleteUser(1000, "admin")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed delete mentee", func(t *testing.T) {
		repo.On("DeleteUserMentor", mock.Anything).Return(errors.New("error delete user")).Once()
		srv := New(repo)
		err := srv.DeleteUser(1000, "admin")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestSingleUser(t *testing.T) {
	repo := mocks.NewRepoInterface(t)

	t.Run("Success get mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).Return(admin.UserCore{IdUser: 1}, nil).Once()
		srv := New(repo)
		res, err := srv.GetSingleUser(1, "admin")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed get mentee", func(t *testing.T) {
		repo.On("GetSingleMentee", mock.Anything).Return(admin.UserCore{}, errors.New("error get mentee")).Once()
		srv := New(repo)
		res, err := srv.GetSingleUser(1, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Success get mentor", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(admin.UserCore{IdUser: 1}, nil).Once()
		srv := New(repo)
		res, err := srv.GetSingleUser(1000, "admin")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed get menotor", func(t *testing.T) {
		repo.On("GetSingleMentor", mock.Anything).Return(admin.UserCore{}, errors.New("error get mentor")).Once()
		srv := New(repo)
		res, err := srv.GetSingleUser(1000, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateClass(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	class := admin.ClassCore{IdClass: 1, ClassName: "Front End", Status: "active", TotalStudent: 20}
	t.Run("Success Update class", func(t *testing.T) {
		srv := New(repo)
		repo.On("GetSingleClass", mock.Anything).Return(class, nil).Once()
		repo.On("EditClass", mock.Anything).Return(class, nil).Once()
		res, err := srv.UpdateClass(class, "admin")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Not Admin", func(t *testing.T) {
		srv := New(repo)
		input := admin.ClassCore{IdClass: 7, ClassName: "Front End"}
		res, err := srv.UpdateClass(input, "mentee")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update class", func(t *testing.T) {
		srv := New(repo)
		repo.On("GetSingleClass", mock.Anything).Return(class, nil).Once()
		repo.On("EditClass", mock.Anything).Return(admin.ClassCore{}, errors.New("error update class")).Once()

		res, err := srv.UpdateClass(class, "admin")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteClass(t *testing.T) {
	repo := mocks.NewRepoInterface(t)
	t.Run("Not Admin", func(t *testing.T) {
		srv := New(repo)
		err := srv.DeleteClass(1, "mentee")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Success delete class", func(t *testing.T) {
		repo.On("DeleteClass", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteClass(1, "admin")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed delete class", func(t *testing.T) {
		repo.On("DeleteClass", mock.Anything).Return(errors.New("error delet user")).Once()
		srv := New(repo)
		err := srv.DeleteClass(1, "admin")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

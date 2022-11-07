package services

import (
	"be12/mentutor/features/mentor"
	"errors"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type mentorUsecase struct {
	mentorRepo mentor.RepoInterface
}

func New(data mentor.RepoInterface) mentor.UsecaseInterface {
	return &mentorUsecase{
		mentorRepo: data,
	}
}

func roleCheck(role string) bool {
	if role == "mentor" {
		return true
	}
	return false
}

func (mu *mentorUsecase) UpdateProfile(input mentor.UserCore, role string) (mentor.UserCore, error) {

	var user mentor.UserCore
	if input.IdUser < 1000 {
		res, err := mu.mentorRepo.GetSingleMentee(input.IdUser)
		if err != nil {
			return mentor.UserCore{}, errors.New("user not found")
		}
		user = res
	} else if input.IdUser >= 1000 {
		res, err := mu.mentorRepo.GetSingleMentor(input.IdUser)
		if err != nil {
			return mentor.UserCore{}, errors.New("user not found")
		}
		user = res
	}

	// CEK KONDISI NAMA
	if input.Name != "" {
		if len(input.Name) < 5 == true || len(input.Name) > 50 == true{
			return mentor.UserCore{}, errors.New("length name not valid")
		}
		var upper, lower, number, sChar, space int
		for _, v := range input.Name {
			if unicode.IsUpper(v) == true {
				upper+=1
			} else if unicode.IsLower(v) ==  true {
				lower += 1
			} else if unicode.IsNumber(v) == true {
				number+=1
			} else if unicode.IsPunct(v){
				sChar+=1
			} else if unicode.IsSpace(v){
				space+=1
			}
		}
		if upper < 1 {
			return mentor.UserCore{}, errors.New("input name not valid")
		} else if lower < 1 {
			return mentor.UserCore{}, errors.New("input name not valid")
		} else if number > 0 {
			return mentor.UserCore{}, errors.New("input name not valid")
		} else if sChar > 0 {
			return mentor.UserCore{}, errors.New("input name not valid")
		} else if space < 1 {
			return mentor.UserCore{}, errors.New("input name not valid")
		}
	} else {
		input.Name = user.Name
	}

	// CEK KONDISI EMAIL
	if input.Email != "" {
		for _, v := range input.Email {
			if unicode.IsSpace(v) {
				return mentor.UserCore{}, errors.New("contain space")
			}
		}
		if len(input.Email) < 8 || len(input.Email) > 40  {
			return mentor.UserCore{}, errors.New("length email not valid")
		} else if strings.Contains(input.Email, "@") == false && strings.Contains(input.Email, ".") == false {
			return mentor.UserCore{}, errors.New("not contain (@) and (.)")
		} 
	} else {
		input.Email = user.Email
	}

	// CEK KONDISI PASSWORD
	if input.Password != ""{
		var sChar = "@#$%^&*<>:;'[]{}|`~!"
		var passUpper, passLower, passNumber, specialChar int
		for _, v := range input.Password {
			if unicode.IsUpper(v) == true {
				passUpper += 1
			} else if unicode.IsLower(v) == true {
				passLower += 1
			} else if unicode.IsNumber(v) == true {
				passNumber += 1
			} else if unicode.IsPunct(v) == true {
				specialChar += 1
			} else if strings.Contains(sChar, string(v)) == true {
				specialChar += 1
			}
		}
		if passUpper < 1 {
			return mentor.UserCore{}, errors.New("string not as expected")
		} else if passLower < 1 {
			return mentor.UserCore{}, errors.New("string not as expected")
		} else if passNumber < 1 {
			return mentor.UserCore{}, errors.New("string not as expected")
		} 
		if specialChar == 0 {
			return mentor.UserCore{}, errors.New("string not as expected")
		} else if len(input.Password) < 8 || len(input.Password) > 30 {
			return mentor.UserCore{}, errors.New("string too short or too long")
		}

		generate , _:= bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		input.Password = string(generate)
	}

	// CEK GAMBAR 
	if input.Images == "" {
		input.Images = user.Images
	}

	if input.IdUser < 1000 {
		res, err := mu.mentorRepo.EditProfileMentee(input)
		if err != nil {
			return mentor.UserCore{}, errors.New("error update user")
		}
		return res, nil
	} else if input.IdUser >= 1000 {
		res, err := mu.mentorRepo.EditProfileMentor(input)
		if err != nil {
			return mentor.UserCore{}, errors.New("error update user")
		}
		return res, nil
	}
	return mentor.UserCore{}, errors.New("error update user")
}

func (mu *mentorUsecase) AddTask(input mentor.TaskCore, role string) (mentor.TaskCore, error) {
	if err := roleCheck(role); err != true {
		return mentor.TaskCore{}, errors.New("user not mentor")
	}

	res, err := mu.mentorRepo.InsertTask(input)
	if err != nil {
		return mentor.TaskCore{}, errors.New("error insert task")
	}
	return res, nil
}

func (mu *mentorUsecase) GetAllTask(role string) ([]mentor.TaskCore, error) {
	if err := roleCheck(role); err != true {
		return []mentor.TaskCore{}, errors.New("user not mentor")
	}

	res, err := mu.mentorRepo.GetAllTask()
	if err != nil {
		return []mentor.TaskCore{}, errors.New("error get all task")
	}
	return res, nil
}

func (mu *mentorUsecase) GetTaskSub(id uint, role string) (mentor.TaskCore, []mentor.SubmissionCore, error) {
	if err := roleCheck(role); err != true {
		return mentor.TaskCore{}, []mentor.SubmissionCore{}, errors.New("user not mentor")
	}

	resTask, resSub, err := mu.mentorRepo.GetTaskSub(id)
	if err != nil {
		return mentor.TaskCore{}, []mentor.SubmissionCore{}, errors.New("error get detail task")
	}
	return resTask, resSub, nil
}
	
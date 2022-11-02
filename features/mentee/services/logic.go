package services

import (
	"be12/mentutor/features/mentee"
)

type MenteeUsecase struct {
	menteeData mentee.RepoInterface
}

func New(data mentee.RepoInterface) mentee.UseCaseInterface {
	return &MenteeUsecase{
		menteeData: data,
	}
}

func (mu *MenteeUsecase) UpdateProfile(id uint, data mentee.MenteeCore) (mentee.MenteeCore, error) {
	res, err := mu.menteeData.EditProfile(id, data)
	if err != nil {
		return mentee.MenteeCore{}, err
	}
	return res, nil
}
package services

import (
	"be12/mentutor/features/mentee"
	"errors"
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

func (mu *MenteeUsecase) InsertStatus(data mentee.Status, token int) (mentee.Status, error) {

	data, err := mu.menteeData.AddStatus(data, token)
	if err != nil {
		return mentee.Status{}, err
	}
	return data, nil
}

func (mu *MenteeUsecase) GetAll() ([]mentee.Status, error) {
	dataAll, err := mu.menteeData.GetAllPosts()
	if err != nil {
		return nil, errors.New("failed get all data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data is still empty")
	}
	return dataAll, nil
}

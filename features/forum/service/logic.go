package service

import (
	"be12/mentutor/features/forum"
	"errors"
)

type postUsecase struct {
	postData forum.DataInterface
}

func New(pd forum.DataInterface) forum.UseCaseInterface {
	return &postUsecase{
		postData: pd,
	}
}

func (ps *postUsecase) AddPost(data forum.Status, token int) (int, error) {

	add, err := ps.postData.Insert(data, token)
	if err != nil || add == 0 {
		return -1, err
	} else {
		return 1, nil
	}

}

func (ps *postUsecase) AddStatus(data forum.Status, token int) (forum.Status, error) {

	_, err := ps.postData.InsertStatus(data, token)
	if err != nil {
		return forum.Status{}, err
	} else {
		return data, nil
	}

}

func (ps *postUsecase) GetAllPosts() ([]forum.Status, error) {
	dataAll, err := ps.postData.GetAll()
	if err != nil {
		return nil, errors.New("failed get all data")
	} else if len(dataAll) == 0 {
		return nil, errors.New("data is still empty")
	} else {
		return dataAll, nil
	}
}

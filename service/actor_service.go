package service

import "dvdrental/repository"

type ActorService interface {
	Create()
	Update()
	Delete()
	Find()
	FindAll()
}

type actorService struct {
	repository repository.ActorRepository
}

func NewActorService(repo repository.ActorRepository) ActorService {
	return &actorService{
		repository: repo,
	}
}

func (service *actorService) Create() {
	panic("NY Implemented")
}

func (service *actorService) Update() {
	panic("NY Implemented")
}

func (service *actorService) Delete() {
	panic("NY Implemented")
}

func (service *actorService) Find() {
	panic("NY Implemented")
}

func (service *actorService) FindAll() {
	panic("NY Implemented")
}

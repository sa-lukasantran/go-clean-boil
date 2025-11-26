package usecase

import (
	"go-basesource/src/adapters"
	"go-basesource/src/entity"
)

type ExampleUseCase struct {
	db      *adapters.Database
	Record  entity.Example
	Records entity.Examples
	Total   int64
}

func NewExampleUseCase() *ExampleUseCase {
	db := adapters.GetDatabase()
	return &ExampleUseCase{db: db}
}

func (uc *ExampleUseCase) Save() error {
	return uc.db.Create(&uc.Record).Error
}

func (uc *ExampleUseCase) FromExampleID(ID string) error {
	return uc.db.First(&uc.Record, ID).Error
}

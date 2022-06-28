package rates

import (
	ratesdomain "github.com/c3llus/simple-service/domain/rates"
)

type Usecase struct {
	rd ratesdomain.DomainItf
}

func InitUsecase(rd ratesdomain.DomainItf) *Usecase {
	return &Usecase{
		rd: rd,
	}
}

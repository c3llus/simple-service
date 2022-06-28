package rates

import "github.com/c3llus/cdk/app/resource"

type DomainItf interface {
}

type DBResItf interface {
}

func InitDomain(db resource.SQLDB) Domain {
	return Domain{
		DB: DBRes{
			DB: db,
		},
	}
}

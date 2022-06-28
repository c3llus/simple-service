package rates

import "github.com/c3llus/cdk/app/resource"

type Domain struct {
	DB DBResItf
}

type DBRes struct {
	DB resource.SQLDB
}

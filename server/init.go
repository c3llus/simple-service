package http

import "github.com/c3llus/cdk/app/resource"

type App interface {
	Resource() resource.Resource
}

func Init(app App) {
	_, _ = app.Resource().GetSQLDB("postgres")
}

package _interface

import "github.com/labstack/echo/v4"

type IRecommendFoodHandler interface {
	Recommend(c echo.Context) error
}
type ISelectFoodHandler interface {
	Select(c echo.Context) error
}

type IHistoryFoodHandler interface {
	History(c echo.Context) error
}

type IMetaFoodHandler interface {
	Meta(c echo.Context) error
}
type IRankingFoodHandler interface {
	Ranking(c echo.Context) error
}
type IImageUploadFoodHandler interface {
	ImageUpload(c echo.Context) error
}
type IEmptyImageFoodHandler interface {
	EmptyImage(c echo.Context) error
}

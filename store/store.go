package store

import (
	"github.com/labstack/echo"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
)

//go:generate mockgen -destination store_mock.go -package store github.com/leoryu/leo-ryu.herokuapp.com/store Store

type Store interface {
	SavePaper(paper *model.Paper) error
	ModifyPaper(paper *model.Paper) error
	GetPaper(id string) (*model.Paper, error)
	DeletePaper(id string) error
}

func SavePaper(c echo.Context, paper *model.Paper) error {
	return FromContext(c).SavePaper(paper)
}

func ModifyPaper(c echo.Context, paper *model.Paper) error {
	return FromContext(c).ModifyPaper(paper)
}

func GetPaper(c echo.Context, id string) (*model.Paper, error) {
	return FromContext(c).GetPaper(id)
}

func DeletePaper(c echo.Context, id string) error {
	return FromContext(c).DeletePaper(id)
}
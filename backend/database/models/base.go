package models

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

func IsNotFound(err error) {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); !ok && err != nil {
		color.Red(fmt.Sprintf("error :%v \n ", err))
	}
}

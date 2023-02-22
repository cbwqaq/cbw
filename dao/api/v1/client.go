package v1

import (
	"GINCHAT/driver"
	"context"
	"fmt"
)

type Client interface {
	Create(ctx context.Context, user Student) error
}

type clientImpl struct {
}

func Create(user Student) {
	var u Student
	driver.DB().Debug().Where("studentId = ?", 12).Find(&u)

	fmt.Println("success", u)

}

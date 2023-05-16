package db

import (
	"fmt"
	"github.com/asdine/storm/v3/q"
	"golang.org/x/crypto/bcrypt"
	"kubepi_password/internal/core"
	"kubepi_password/internal/model"
	"os"
	"time"
)

func GetByNameOrEmail(el string) (*model.User, error) {
	var us model.User
	query := core.DB.Select(q.Or(q.Eq("Name", el), q.Eq("Email", el)))

	if err := query.First(&us); err != nil {
		return nil, err
	}
	return &us, nil
}
func UpdatePassword() error {
	cu, err := GetByNameOrEmail(core.Username)
	if cu == nil {
		fmt.Printf("找不到用户:%s\n", core.Username)
		os.Exit(2)
	}
	if err != nil {
		return err
	}
	bs, err := bcrypt.GenerateFromPassword([]byte(core.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cu.Authenticate.Password = string(bs)
	cu.UpdateAt = time.Now()
	err = core.DB.Update(cu)
	if err != nil {
		fmt.Printf("用户:%s\n密码:%s\n状态:更新失败\n", core.Username, core.Password)
	} else {
		fmt.Printf("用户:%s\n密码:%s\n状态:更新成功\n", core.Username, core.Password)
	}
	cu, err = GetByNameOrEmail(core.Username)
	if err := bcrypt.CompareHashAndPassword([]byte(cu.Authenticate.Password), []byte(core.Password)); err != nil {
		fmt.Println("校验失败")
	} else {
		fmt.Println("校验成功")
	}
	return err
}

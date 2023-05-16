package cmd

import (
	"flag"
	"fmt"
	"github.com/asdine/storm/v3"
	"go.etcd.io/bbolt"
	"kubepi_password/internal/core"
	"kubepi_password/internal/db"
	"os"
	"time"
)

func init() {
	flag.StringVar(&core.Username, "username", "", "指定用户名")
	flag.StringVar(&core.Password, "password", "123456", "指定新的密码")
	flag.StringVar(&core.Dbpath, "dbpath", "/var/lib/kubepi/db/kubepi.db", "指定数据库路径")
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func Execute() {
	flag.Parse()
	var err error
	if !PathExists(core.Dbpath) {
		fmt.Printf("%s file does not exist!\n", core.Dbpath)
		os.Exit(2)
	}
	core.DB, err = storm.Open(core.Dbpath, storm.BoltOptions(0600, &bbolt.Options{Timeout: 3 * time.Second}))
	if err != nil {
		fmt.Println("不能打开数据库!,必须关闭kubepi服务确保文件没有被占用", err)
		os.Exit(2)
	}
	db.UpdatePassword()
	defer core.DB.Close()
}

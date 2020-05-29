package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 15:34
 * @Title: 数据库配置
 * --- --- ---
 * @Desc:
 */

/**
初始化连接数据库(mySql)
*/
func New(opt *Options) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", opt.User, opt.Password, opt.Host, opt.Db))
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}
	if opt.Debug {
		db = db.Debug()
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(opt.MaxIdleConns)
	db.DB().SetMaxOpenConns(opt.MaxOpenConns)

	return db, nil
}

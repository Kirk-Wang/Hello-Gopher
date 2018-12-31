/**
 * 获取数据库连接
 */
package datasource

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/Kirk-Wang/Hello-Gopher/iris-1/conf"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// 主库，单例
func InstanceMaster() *xorm.Engine {
	// 有直接返回
	if masterEngine != nil {
		return masterEngine
	}
	// 互斥锁，因为会有并发的问题
	lock.Lock()
	defer lock.Unlock()

	// 还要验证一下
	// 例如：同时10个进来，都是不存在masterEngine的
	// 其中 1 个锁住了在操作，9 个在等待，
	// 这个解锁后，其它的会继续执行，所以要再判断一次
	if masterEngine != nil {
		return masterEngine
	}
	// 拿到配置信息
	c := conf.MasterDbConfig
	// 拼接连接串
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	// xorm 去连接
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return engine
}

// 从库，单例
func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}
	c := conf.SlaveDbConfig
	engine, err := xorm.NewEngine(conf.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err != nil {
		log.Fatal("dbhelper", "DbInstanceMaster", err)
		return nil
	}
	engine.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = engine
	return engine
}

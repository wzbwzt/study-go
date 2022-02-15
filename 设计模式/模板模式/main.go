package main

import "fmt"

//模板模式(使用较多):有一个结构定义了业务步骤，并允许实现“类”不改变业务结构的情况下，重新定义某些步骤

//譬如结合redis/文件/etcd，获取缓存这个狗功能，实现商品缓存这件事
// 1.根据id判断redis/文件/etcd是否有值，有取redis/文件/etcd,没有取数据库
// 2.按照格式输出缓存，并记录该商品id的访问次数

type ICach interface {
	Get(int64) interface{}
	OutPut(int64) interface{}
}

//##############################################################################
type Cache struct { //总"类"
	c ICach
}

func (this *Cache) Get(id int64) interface{} {
	return nil
}

func (this *Cache) OutPut(id int64) interface{} {
	return map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": this.c.Get(id),
	}
}

//##############################################################################

//重写子类1
type RedisCache struct {
	*Cache
}

func (this *RedisCache) Get(id int64) interface{} {
	return fmt.Sprintf("get data %v from redis", id)
}

func NewRedisCache() *RedisCache {
	this := &RedisCache{}
	this.Cache = &Cache{c: this}
	return this
}

//重写子类2
type FileCache struct {
	*Cache
}

func (this *FileCache) Get(id int64) interface{} {
	return fmt.Sprintf("get data %v from file", id)
}

func NewFileCache() *FileCache {
	this := &FileCache{}
	this.Cache = &Cache{c: this}
	return this
}

//##############################################################################
func main() {
	cache := NewRedisCache()
	fmt.Println(cache.OutPut(12))

	cache2 := NewFileCache()
	fmt.Println(cache2.OutPut(12))
}

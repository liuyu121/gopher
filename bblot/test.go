package main

import (
	"fmt"
	"log"

	bolt "github.com/liuyu121/bbolt"
)

var world = []byte("world")

func main() {
	// 打开一个文件，获取 db 对象，所有操作都是围绕这个 db 对象进行
	db, err := bolt.Open("bblot.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭 db
	defer db.Close()

	// 存储数据
	key := []byte("hello")
	value := []byte("hello world")

	key2 := []byte("name")
	value2 := []byte("liuduoyu")
	value2_2 := []byte("xxxxx")

	//keys := []string{"hello", "name"}
	//values := []string{"hello world", "liuduoyu"}

	// 插入数据
	err = db.Update(func(tx *bolt.Tx) error {
		// 创建一个名称为 hello 的 bucket
		bucket, err := tx.CreateBucketIfNotExists(world)
		fmt.Println(bucket)
		if err != nil {
			return err
		}

		// 一个 bucket 可以写入多个 k- pair，说明 bucket
		err = bucket.Put(key, value)
		err = bucket.Put(key2, value2)
		err = bucket.Put(key2, value2_2)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// 查找关键字
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		fmt.Println(bucket)

		if bucket == nil {
			return fmt.Errorf("bucket %q not fount", world)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		val2 := bucket.Get(key2)
		fmt.Println(string(val2))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

package interfacefunccode

/*
接口型函数的使用(重要重要重要重要重要重要重要重要重要重要重要重要重要)，步骤：
1 定义接口
2 定义函数类型实现接口
3 实现接口中的函数
*/
//接口
type Getter interface {
	Get(key string) ([]byte, error)
}

//函数类型实现接口
type Getterfunc func(key string) ([]byte, error)

//实现接口中的方法******************重要
func (f Getterfunc) Get(key string) ([]byte, error) {
	return f(key)
}

//使用接口函数
func GetFromSource(getter Getter, key string) []byte {
	buff, err := getter.Get(key)
	if err == nil {
		return buff
	}
	return nil
}

/*************************2 方式一：GetterFunc 类型的函数作为参数*************/
// 2.1 使用匿名函数调用接口函数
func Interfacefunstd() []byte {
	//匿名函数调用
	return GetFromSource(Getterfunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	}), "hello")
}

//2.2 使用普通函数调用接口函数
func test(key string) ([]byte, error) {
	return []byte(key), nil
}
func Interfacefunstd2() []byte {
	//匿名函数调用
	return GetFromSource(Getterfunc(test), "hello2")
}

/**********************3 方式二：实现了 Getter 接口的结构体作为参数***************/
type DB struct{ url string }

func (db *DB) Query(sql string, args ...string) string {
	// ...
	return "hello"
}

func (db *DB) Get(key string) ([]byte, error) {
	// ...
	v := db.Query("SELECT NAME FROM TABLE WHEN NAME= ?", key)
	return []byte(v), nil
}

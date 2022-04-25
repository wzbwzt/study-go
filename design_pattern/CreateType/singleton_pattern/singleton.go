package singleton

// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = newSingleton()
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}

func newSingleton() *Singleton {
	return &Singleton{}
}

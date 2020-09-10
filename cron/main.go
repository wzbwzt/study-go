package main

import (
	"fmt"
	"os"
	"sync"
	"github.com/pkg/errors"
	cron "github.com/robfig/cron/v3"
)

// Crontab crontab manager
type Crontab struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

// NewCrontab new crontab
func NewCrontab() *Crontab {
	return &Crontab{
		inner: cron.New(),
		ids:   make(map[string]cron.EntryID),
	}
}

// IDs ...
func (c *Crontab) IDs() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIDs := make([]string, 0, len(c.ids))
	invalidIDs := make([]string, 0)
	for sid, eid := range c.ids {
		if e := c.inner.Entry(eid); e.ID != eid {
			invalidIDs = append(invalidIDs, sid)
			continue
		}
		validIDs = append(validIDs, sid)
	}
	for _, id := range invalidIDs {
		delete(c.ids, id)
	}
	return validIDs
}

// Start start the crontab engine
func (c *Crontab) Start() {
	c.inner.Start()
}

// Stop stop the crontab engine
func (c *Crontab) Stop() {
	c.inner.Stop()
}

// DelByID remove one crontab task
func (c *Crontab) DelByID(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid, ok := c.ids[id]
	if !ok {
		return
	}
	c.inner.Remove(eid)
	delete(c.ids, id)
}

// AddByID add one crontab task
// id is unique
// spec is the crontab expression
func (c *Crontab) AddByID(id string, spec string, cmd cron.Job) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.Errorf("crontab id exists")
	}
	eid, err := c.inner.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// AddByFunc add function as crontab task
func (c *Crontab) AddByFunc(id string, spec string, f func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.Errorf("crontab id exists")
	}
	eid, err := c.inner.AddFunc(spec, f)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// IsExists check the crontab task whether existed with job id
func (c *Crontab) IsExists(jid string) bool {
	_, exist := c.ids[jid]
	return exist
}

//代码实现很简单,每个函数的作用都可以参考注释.下面简单实用一下上面的封装:

type testTask struct {
}

func (t *testTask) Run() {
	fmt.Println("hello world")
}

//func taskTick() (err error) {
//	fmt.Println("Joel...")
//	return errors.New("task failed!!!")
//}
func main() {
	crontab := NewCrontab()
	// 实现接口的方式添加定时任务
	task := &testTask{}
	err := crontab.AddByID("1", "* * * * *", task)
	if err != nil {
		fmt.Printf("error to add crontab task:%s", err)
		os.Exit(-1)
	}

	// 添加函数作为定时任务
	//taskFunc := func() {
	//	fmt.Println("hello world")
	//}
	//if err := crontab.AddByFunc("2", "* * * * *", taskFunc); err != nil {
	//	fmt.Printf("error to add crontab task:%s", err)
	//	os.Exit(-1)
	//}
	crontab.Start()
	select {}

}


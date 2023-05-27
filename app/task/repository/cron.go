package repository

import (
	"fmt"
	"github.com/pkg/errors"
	cron "github.com/robfig/cron/v3"
	"sync"
)

var SelfCron *Crontab

// Crontab crontab manager
type Crontab struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

func InitCrontab() {
	SelfCron = NewCrontab()
	//初始化
	//任务定时器列表
	var taskRule = map[int]string{
		1: "0 0 2 * * ?",
	}
	for i, v := range taskRule {
		//如果存在先删除
		if SelfCron.IsExists(fmt.Sprint(i)) {
			SelfCron.DelById(fmt.Sprint(i))
		}
		award := &Award{}
		award.Times = i
		err := SelfCron.AddById(fmt.Sprint(i), v, award)
		if err != nil {
			fmt.Println(err)
		}
	}

	SelfCron.Start()
	select {}
}

// NewCrontab new crontab
func NewCrontab() *Crontab {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return &Crontab{
		inner: cron.New(cron.WithParser(secondParser), cron.WithChain()),
		ids:   make(map[string]cron.EntryID),
	}
}

// Ids ...
func (c *Crontab) Ids() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIds := make([]string, 0, len(c.ids))
	invalidIds := make([]string, 0)
	for sid, eid := range c.ids {
		if e := c.inner.Entry(eid); e.ID != eid {
			invalidIds = append(invalidIds, sid)
			continue
		}
		validIds = append(validIds, sid)
	}
	for _, id := range invalidIds {
		delete(c.ids, id)
	}
	return validIds
}

// Start start the crontab engine
func (c *Crontab) Start() {
	c.inner.Start()
}

// Stop stop the crontab engine
func (c *Crontab) Stop() {
	c.inner.Stop()
}

// DelById remove one crontab task
func (c *Crontab) DelById(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid, ok := c.ids[id]
	if !ok {
		return
	}
	c.inner.Remove(eid)
	delete(c.ids, id)
}

// AddById add one crontab task
// id is unique
// spec is the crontab expression
func (c *Crontab) AddById(id string, spec string, cmd cron.Job) error {
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

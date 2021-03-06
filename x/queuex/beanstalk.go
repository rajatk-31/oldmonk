// Package queuex provides primitives for communicating with all queue
package queuex

import (
	"github.com/beanstalkd/go-beanstalk"
	oldmonkv1 "github.com/evalsocket/oldmonk/pkg/apis/oldmonk/v1"
	"strconv"
)

// BeanstalkController hold configration and client for beanstalkd queue
type BeanstalkController struct {
	Config *oldmonkv1.ListOptions
	Client *beanstalk.Tube
}

// NewBeanstalkClient create a new BeanstalkController object
// It returns the BeanstalkController
func NewBeanstalkClient(config *oldmonkv1.ListOptions) *BeanstalkController {
	c, err := beanstalk.Dial("tcp", config.Uri)
	if err != nil {
		logger.Error("unable to dial", err)
		return &BeanstalkController{}
	}
	ts := &beanstalk.Tube{
		Conn: c,
		Name: config.Tube,
	}
	beanstalkController := BeanstalkController{
		Client: ts,
		Config: config,
	}
	return &beanstalkController
}

// GetCount count the number of message in a tube
// It returns the number of Messages in a tube
func (b *BeanstalkController) GetCount() int32 {
	states, err := b.Client.Conn.Stats()
	if err != nil {
		logger.Error("error in getting state", err)
		return 0
	}
	i, err := strconv.Atoi(states[b.Config.Key])
	if err != nil {
		logger.Error("error in getting converting string to int32", err)
	}
	return int32(i)
}

// Close will close  beanstalk connection
// It returns the error
func (r *BeanstalkController) Close() error {
	return nil
}

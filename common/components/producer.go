package components

import (
	"encoding/json"

	"math/rand"

	"github.com/bitly/go-nsq"
	"github.com/hjhcode/deploy-web/common/g"
)

type SendMess struct {
	SubmitType string `json:"submit_type"`
	SubmitId   int64  `json:"submit_id"`
}

func Send(topic string, sendMess *SendMess) bool {
	ipList := g.Conf().Nsq.Address
	msg, err := json.Marshal(sendMess)
	if err != nil {
		return false
	}
	result := false
	for _, i := range rand.Perm(len(ipList)) {
		if flag, _ := postNsq(ipList[i], topic, msg); flag == true {
			result = true
			break
		}
	}
	return result
}

func postNsq(address, topic string, msg []byte) (bool, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(address, config)
	if err != nil {
		return false, err
	}
	err = producer.Publish(topic, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}

package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

// 连接信息
const MQURL = "amqp://develop:123456@127.0.0.1:5672/task_1130"

// RabbitMQ 结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机名称
	Exchange string
	// bind Key 名称
	Key string
	// 链接信息
	MqUrl string
	sync.Mutex
}

// 创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange: exchange,
		Key: key,
		MqUrl: MQURL,
	}
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// failOnErr 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}

// 创建简单模式下的RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

// 直接模式下的生产队列
func (r *RabbitMQ) PublishSimple(message string) error {
	r.Lock()
	defer r.Unlock()
	// 1、申请队列，如果不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		return err
	}

	// 调用channel 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
		)
	if err != nil {
		return err
	}
	return nil
}

// simple 模式下的消费者
func (r *RabbitMQ) ConsumeSimple() {
	// 1、申请队列，如果队列不存在会创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		fmt.Println(err)
	}

	// 消费者流量控制
	// 当前消费者一次能接受的最大消息数量、服务器传递的最大容量（以八位字节为单位）、如果设置为true 对channel可用
	err = r.channel.Qos(1, 0, false)
	if err != nil {
		fmt.Println(err)
	}

	//接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		//这里要改掉，我们用手动应答
		false, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", d.Body)
			//message := &datamodels.Message{}
			//err :=json.Unmarshal([]byte(d.Body),message)
			//if err !=nil {
			//	fmt.Println(err)
			//}
			////插入订单
			//_,err=orderService.InsertOrderByMessage(message)
			//if err !=nil {
			//	fmt.Println(err)
			//}
			//
			////扣除商品数量
			//err = productService.SubNumberOne(message.ProductID)
			//if err !=nil {
			//	fmt.Println(err)
			//}
			//如果为true表示确认所有未确认的消息，
			//为false表示确认当前消息
			_ = d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}



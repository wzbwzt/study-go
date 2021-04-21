# 消息分发
使用任务队列的优点之一是能够轻松并行化工作。如果我们的工作正在积压，我们可以增加更多的工人，这样就可以轻松扩展。
默认情况下，RabbitMQ将按顺序将每个消息发送给下一个消费者。平均而言，每个消费者都会收到相同数量的消息。
这种分发消息的方式称为轮询。

也可以设置告诉RabbitMQ不要一次向一个worker发出多个消息。或者，换句话说，在处理并确认前一条消息之前，不要向worker发送新消息。
相反，它将把它发送给下一个不忙的worker
将预取计数设置为1
```go
err = ch.Qos(
  1,     // prefetch count
  0,     // prefetch size
  false, // global
)

```
>如果所有的worker都很忙,可能队列会满,需要增加worker或者其他测略


# 消息确认
work 完成任务可能需要耗费几秒钟，如果一个worker在任务执行过程中宕机了该怎么办呢？(其通道已关闭，连接已关闭或TCP连接丢失)
在这种情况下，如果你终止一个worker那么你就可能会丢失这个任务，我们还将丢失所有已经交付给这个worker的尚未处理的消息。
为了确保消息永不丢失，RabbitMQ支持 消息确认。消费者发送回一个确认（acknowledgement），以告知RabbitMQ已经接收，处理了特定的消息，
并且RabbitMQ可以自由删除它。
在返回一个Delivery的通道时可以设置自动回复确认ack信息：
```go
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
```

也可以手动回复一个ack消息
```go
amqp.Delivery.Ack(false) // 手动传递消息确认
```
>消息确认必须在接收消息的同一通道（Channel）上发送。尝试使用不同的通道（Channel）进行消息确认将导致通道级协议异常。有关更多信息，请参阅确认的文档指南
>https://www.rabbitmq.com/confirms.html

# 消息持久化
如果RabbitMQ服务器停止运行，我们的任务仍然会丢失,因为默认情况下，rabbitmq是不会将数据持久化到磁盘
当RabbitMQ退出或崩溃时，它将忘记队列和消息，除非您告诉它不要这样做。要确保消息不会丢失，需要做两件事：我们需要将队列和消息都标记为持久的。
```go
q, err := ch.QueueDeclare(
	"hello", // name
	true,    // durable
	false,   // delete when unused
	false,   // exclusive
	false,   // no-wait
	nil,     // arguments
)

```
>注意： RabbitMQ不允许你使用不同的参数重新定义现有队列，并将向任何尝试重新定义的程序返回错误。如果需要更改已经声明的的队列为未持久化的，最快速的方法就是声明一个具有不同名称的队列一个新的队列

>将消息标记为持久性并不能完全保证消息不会丢失。尽管它告诉RabbitMQ将消息保存到磁盘上，但是RabbitMQ接受了一条消息并且还没有保存它时，仍然有一个很短的时间窗口。而且，RabbitMQ并不是对每个消息都执行fsync(2)——它可能只是保存到缓存中，而不是真正写入磁盘。持久性保证不是很强，但是对于我们的简单任务队列来说已经足够了。如果您需要更强有力的担保，那么您可以使用https://www.rabbitmq.com/confirms.html


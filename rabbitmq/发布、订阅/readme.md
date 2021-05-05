# 发布订阅
与工作队列恰恰相反：
工作队列是每条消息只有一个消费者，不存在一个任务被多个worker领取；
发布订阅是将向多个消费者传递一个消息,一条消息被多次消费

# Exchanges（交换器）
完整的消息传递模型:
核心思想是生产者从不将任何消息直接发送到队列。实际上，生产者经常甚至根本不知道是否将消息传递到任何队列。
相反，生产者只能将消息发送到交换器。交换器是非常简单的东西。一方面，它接收来自生产者的消息，另一方面，将它们推入队列。
交换器必须确切知道如何处理接收到的消息。它应该被附加到特定的队列吗？还是应该将其附加到许多队列中？或者它应该被丢弃。这些规则由交换器的类型定义。

交换器类型:
direct, topic, headers 和 fanout。

```go
//创建一个fanout(扇出)交换器,名称是logs
//fanout（扇出）交换器非常简单。它只是将接收到的所有消息广播到它知道的所有队列中。
err = ch.ExchangeDeclare(
  "logs",   // name
  "fanout", // type
  true,     // durable
  false,    // auto-deleted
  false,    // internal
  false,    // no-wait
  nil,      // arguments
)

//发布到定义的交换器中
body := bodyFrom(os.Args)
err = ch.Publish(
  "logs", // exchange
  "",     // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing{
          ContentType: "text/plain",
          Body:        []byte(body),
  })
```

## 临时队列
我们传递一个空字符串作为队列名称时，我们将使用随机生成的名称创建的一个临时队列:
```go
q, err := ch.QueueDeclare(
  "",    // 空字符串作为队列名称
  false, // 非持久队列
  false, // delete when unused
  true,  // 独占队列（当前声明队列的连接关闭后即被删除）
  false, // no-wait
  nil,   // arguments
)

```

## 查看交换器列表
```sh
sudo rabbitmqctl list_exchanges
```
>在此列表中，将有一些`amq.*`交换器和一个默认的（未命名）交换器。这些是默认创建的，但是你现在不太可能需要使用它们。
>默认交换器:由空字符串（`""`）标识

```go
 err = ch.Publish(
   "",     // exchange
   q.Name, // routing key
   false,  // mandatory
   false,  // immediate
   amqp.Publishing{
     ContentType: "text/plain",
     Body:        []byte(body),
 })
```

在这里，我们使用默认或无名称的交换器：消息将以`route_key`参数指定的名称路由到队列（如果存在）

# 绑定
我们已经创建了一个扇出交换器和一个队列。现在我们需要告诉交换器将消息发送到我们的队列。交换器和队列之间的关系称为*绑定*。
```go
err = ch.QueueBind(
  q.Name, // queue name
  "",     // routing key :交换器类型为fanout可以忽略routing_key
  "logs", // exchange
  false,
  nil,
)

```

## 罗列绑定关系列表
```sh
rabbitmqctl list_bindings
```


# 路由
>当我们需要使它能够只订阅消息的一个子集。例如，我们将只能将关键错误消息定向到日志文件（以节省磁盘空间），同时仍然能够在控制台上打印所有日志消息。
>需要路由来实现


## direct交换器
direct交换器背后的路由算法很简单——消息进入其binding key与消息的routing key完全匹配的队列。
发布
```go
//定义交换器
err = ch.ExchangeDeclare(
  "logs_direct", // name
  "direct",      // type
  true,          // durable
  false,         // auto-deleted
  false,         // internal
  false,         // no-wait
  nil,           // arguments
)
failOnError(err, "Failed to declare an exchange")

body := bodyFrom(os.Args)
err = ch.Publish(
  "logs_direct",         // exchange
  severityFrom(os.Args), // routing key
  false, // mandatory
  false, // immediate
  amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte(body),
})

```

订阅
>为感兴趣的每种严重性（日志级别）创建一个新的绑定
```go
q, err := ch.QueueDeclare(
  "",    // name
  false, // durable
  false, // delete when unused
  true,  // exclusive
  false, // no-wait
  nil,   // arguments
)
failOnError(err, "Failed to declare a queue")

if len(os.Args) < 2 {
  log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
  os.Exit(0)
}
// 建立多个绑定关系
for _, s := range os.Args[1:] {
  log.Printf("Binding queue %s to exchange %s with routing key %s",
     q.Name, "logs_direct", s)
  err = ch.QueueBind(
    q.Name,        // queue name
    s,             // routing key
    "logs_direct", // exchange
    false,
    nil)
  failOnError(err, "Failed to bind a queue")
}

```

## topic交换器
发送到topic交换器的消息不能具有随意的routing_key——它必须是单词列表，以点分隔;可以包含任意多个单词，最多255个字节
topic交换器背后的逻辑类似于direct交换器——用特定路由键发送的消息将传递到所有匹配绑定键绑定的队列;
绑定键有两个重要的特殊情况： - *（星号）可以代替一个单词。 - ＃（井号）可以替代零个或多个单词。

>topic交换器功能强大，可以像其他交换器一样运行。
>当队列用“#”（井号）绑定键绑定时，它将接收所有消息，而与路由键无关，就像在fanout交换器中一样。
>当在绑定中不使用特殊字符“*”（星号）和“#”（井号）时，topic交换器的行为就像direct交换器一样。

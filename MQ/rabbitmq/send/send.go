// send.go - Создаёт очередь RabbitMQ и отправляет в неё сообщение
package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://USERNAME:PASSWD@192.168.0.156:5672")
	failOnError(err, "Не удалось подключиться к RabbitMQ: ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Не удалось открыть канал: ")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-go", // name, имя очереди
		true,       // durable, сохранять ли сообщения, если кластер был перезапущен
		false,      // autoDelete, если нет получателя, тогда удалять сообщения
		false,      // exclusive, выдать ошибку, если будут подключаться другие получатели
		false,      // noWait, не ждать пока очередь будет успешно установлена
		nil)        // args, аргументы, к примеру для отбора из очереди по хэдеру

	failOnError(err, "Не удалось создать очередь: ")

	t := time.Now().Format("15:04:05")
	body := "This is my message! " + t // само сообщение (тут должна быть полезная нагрузка)

	err = ch.Publish(
		"",
		q.Name, // имя очереди, в которую отправлять сообщения
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain", // тип сообщения
			Body:        []byte(body), // формат тела сообщения
		})
	failOnError(err, "Не удалось опубликовать сообщение: ")
}

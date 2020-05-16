// receive.go - подключается к очереди RabbitMQ и получает сообщения
package main

import (
	"log"

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

	msgs, err := ch.Consume(
		q.Name, // имя очереди
		"",     // consumer, ?
		true,   // auto-acknowledge, автоматически принимает сообщения
		false,  // exclusive, заперетить другие косумеры
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Не удалось зарегистровать получатель: ")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Получено сообщение: %s", d.Body)
		}
	}()

	log.Printf(" [*] Ожидаю сообщения. Для выхода нажмиите CTRL+C")
	<-forever
}

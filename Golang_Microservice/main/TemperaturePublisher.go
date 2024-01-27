package main

import (
	"fmt"
	"golang_weather_fetcher/myproject/fetcher"

	"github.com/streadway/amqp"
)

const (
	rabbitMQURL                      = "amqp://guest:guest@localhost:5672/" // Replace with your RabbitMQ server URL
	queueName                        = "temperature_data"
	rabbitMqDown                     = "RabbitMq down"
	rabbitMqUp                       = "RabbitMq up"
	unableToCreateQueue              = "unable to create queue"
	unableToPublishData              = "unable tp publis data"
	successfullyPublishedDataToQueue = "successfully published data to queue"
)

type TemperaturePublisher struct {
	JsonData string
}

func main() {
	tf := &fetcher.TemperatureFetcher{}
	cities := []string{"hyderabad", "bangalore"}
	ApiKey := "955c361fea4b58ee9b3f7bde50ce5237"

	tf.ConstructData(cities, ApiKey)

	tp := &TemperaturePublisher{}
	tp.JsonData = tf.FetchTemperature()

	fmt.Println("jsonData to be stored in rabbitMq", tp.JsonData)

	status := tp.PingRabbitMq()
	if status == rabbitMqDown {
		fmt.Println(rabbitMqDown)
	} else {
		errToPublish := tp.publishToRabbitMq(tp.JsonData, queueName)
		fmt.Print(errToPublish)
	}
	return
}

func (tp *TemperaturePublisher) publishToRabbitMq(JsonData string, queueName string) string {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	channel, errConnectingToChannel := conn.Channel()
	if errConnectingToChannel != nil {
		fmt.Println(errConnectingToChannel.Error())
	}

	defer channel.Close()

	queue, errorCreatingQueue := channel.QueueDeclare(
		queueName, // Queue name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)

	if errorCreatingQueue != nil {
		fmt.Println(errorCreatingQueue.Error())
		return unableToCreateQueue
	}

	errorPublishingQueue := channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(JsonData),
		},
	)

	if errorPublishingQueue != nil {
		fmt.Println(errorPublishingQueue.Error())
		return errorPublishingQueue.Error()
	}else{
		tp.consumeFromRabbitMq(channel)
	}
	
	return successfullyPublishedDataToQueue
}

func (tp *TemperaturePublisher) PingRabbitMq() string {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		fmt.Println(err.Error())
		return rabbitMqDown
	}
	defer conn.Close()

	return rabbitMqUp
}

func (tp *TemperaturePublisher) consumeFromRabbitMq(channel *amqp.Channel) {
	consumedMessages, errorWHileConsuming := channel.Consume(
		queueName, // Queue name
		"",        // Consumer
		true,      // Auto-Acknowledge
		false,     // Exclusive
		false,     // No-Local
		false,     // No-Wait
		nil,       // Args
	)
	if errorWHileConsuming != nil {
		fmt.Print("Failed to register a consumer: %v", errorWHileConsuming)
	}

	fmt.Println("Waiting for messages. To exit press CTRL+C")
	for msg := range consumedMessages {
		fmt.Printf("Received a message: %s\n", msg.Body)
	}
	defer channel.Close()
	return
}

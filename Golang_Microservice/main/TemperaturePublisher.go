package main





import (
	"fmt"
	"golang_weather_fetcher/myproject/fetcher"

	"golang_weather_fetcher/myproject/constants"
	"golang_weather_fetcher/myproject/purger"
	"golang_weather_fetcher/myproject/Data"
	"github.com/streadway/amqp"

)

type TemperaturePublisher struct {
	JsonData string
}

func main() {
	fmt.Println("Running RealTime")

	tf := &fetcher.TemperatureFetcher{}
	cities := Data.GetCities()
	ApiKey := "955c361fea4b58ee9b3f7bde50ce5237"

	tf.ConstructData(cities, ApiKey)

	tp := &TemperaturePublisher{}
	tp.JsonData = tf.FetchTemperature()

	fmt.Println("jsonData to be stored in rabbitMq", tp.JsonData)

	status := tp.PingRabbitMq()

	if status == constants.RabbitMqDown {
		fmt.Println(constants.RabbitMqDown)
		return
	} else {
		fmt.Println(constants.RabbitMqUp)
		// below line to remove data from queue
		purge, error := purger.RemoveDataInQueue(constants.QueueName)
		if error == nil {
			fmt.Println(purge)
			errToPublish := tp.publishToRabbitMq(tp.JsonData, constants.QueueName)
			fmt.Print(errToPublish)
		}

	}
}

// test

// func main() {
// 	fmt.Println("Running TestData")
// 	tf := &fetcher.TemperatureFetcher{}
// 	cities := []string{"hyderabad", "bangalore"}
// 	ApiKey := "955c361fea4b58ee9b3f7bde50ce5237"

// 	tf.ConstructData(cities, ApiKey)

// 	tp := &TemperaturePublisher{}
// 	tp.JsonData = `{
// 		"New York" : {"temp": 277.66, "humidity": 52, "fetched_time": 1706981278, "fetched_time_str": "2024-02-03 22:57:58"},
// 		"Tokyo" : {"temp": 278.86, "humidity": 58, "fetched_time": 1706981278, "fetched_time_str": "2024-02-03 22:57:58"},
// 		"London" : {"temp": 286.06, "humidity": 85, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Paris" : {"temp": 283.38, "humidity": 88, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Beijing" : {"temp": 265.09, "humidity": 33, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Moscow" : {"temp": 270.46, "humidity": 95, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Sydney" : {"temp": 294.78, "humidity": 83, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Cairo" : {"temp": 289.57, "humidity": 33, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Mumbai" : {"temp": 302.14, "humidity": 34, "fetched_time": 1706981279, "fetched_time_str": "2024-02-03 22:57:59"},
// 		"Rio de Janeiro" : {"temp": 301.60, "humidity": 78, "fetched_time": 1706981280, "fetched_time_str": "2024-02-03 22:58:00"},
// 		"Istanbul" : {"temp": 279.74, "humidity": 65, "fetched_time": 1706981280, "fetched_time_str": "2024-02-03 22:58:00"},
// 		"Berlin" : {"temp": 282.76, "humidity": 93, "fetched_time": 1706981280, "fetched_time_str": "2024-02-03 22:58:00"},
// 		"Johannesburg" : {"temp": 294.70, "humidity": 63, "fetched_time": 1706981280, "fetched_time_str": "2024-02-03 22:58:00"},
// 		"Toronto" : {"temp": 274.36, "humidity": 66, "fetched_time": 1706981280, "fetched_time_str": "2024-02-03 22:58:00"},
// 		"Buenos Aires" : {"temp": 309.54, "humidity": 56, "fetched_time": 1706981280, "fetched_time_str": "2024-02-03 22:58:00"},
// 		"Seoul" : {"temp": 276.11, "humidity": 75, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Los Angeles" : {"temp": 285.77, "humidity": 75, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Rome" : {"temp": 287.11, "humidity": 47, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Dubai" : {"temp": 294.15, "humidity": 43, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Bangkok" : {"temp": 300.85, "humidity": 89, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Nairobi" : {"temp": 295.77, "humidity": 53, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Amsterdam" : {"temp": 282.59, "humidity": 85, "fetched_time": 1706981281, "fetched_time_str": "2024-02-03 22:58:01"},
// 		"Singapore" : {"temp": 299.70, "humidity": 83, "fetched_time": 1706981282, "fetched_time_str": "2024-02-03 22:58:02"},
// 		"Mexico City" : {"temp": 293.88, "humidity": 11, "fetched_time": 1706981282, "fetched_time_str": "2024-02-03 22:58:02"},
// 		"Stockholm" : {"temp": 278.67, "humidity": 61, "fetched_time": 1706981282, "fetched_time_str": "2024-02-03 22:58:02"},
// 		"Athens" : {"temp": 283.51, "humidity": 67, "fetched_time": 1706981282, "fetched_time_str": "2024-02-03 22:58:02"},
// 		"Vancouver" : {"temp": 279.81, "humidity": 89, "fetched_time": 1706981282, "fetched_time_str": "2024-02-03 22:58:02"},
// 		"Cape Town" : {"temp": 301.67, "humidity": 58, "fetched_time": 1706981282, "fetched_time_str": "2024-02-03 22:58:02"},
// 		"Prague" : {"temp": 281.31, "humidity": 80, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Jakarta" : {"temp": 300.08, "humidity": 83, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Vienna" : {"temp": 283.66, "humidity": 72, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Auckland" : {"temp": 289.51, "humidity": 66, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Chicago" : {"temp": 276.45, "humidity": 77, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Barcelona" : {"temp": 289.05, "humidity": 71, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Oslo" : {"temp": 276.73, "humidity": 68, "fetched_time": 1706981283, "fetched_time_str": "2024-02-03 22:58:03"},
// 		"Warsaw" : {"temp": 281.80, "humidity": 75, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Budapest" : {"temp": 283.45, "humidity": 70, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Riyadh" : {"temp": 287.23, "humidity": 27, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Helsinki" : {"temp": 274.92, "humidity": 85, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Delhi" : {"temp": 289.20, "humidity": 72, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Santiago" : {"temp": 305.77, "humidity": 22, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Manila" : {"temp": 299.64, "humidity": 78, "fetched_time": 1706981284, "fetched_time_str": "2024-02-03 22:58:04"},
// 		"Brasília" : {"temp": 297.66, "humidity": 65, "fetched_time": 1706981285, "fetched_time_str": "2024-02-03 22:58:05"},
// 		"Oslo" : {"temp": 276.73, "humidity": 68, "fetched_time": 1706981285, "fetched_time_str": "2024-02-03 22:58:05"},
// 		"Warsaw" : {"temp": 281.80, "humidity": 75, "fetched_time": 1706981285, "fetched_time_str": "2024-02-03 22:58:05"},
// 		"Budapest" : {"temp": 283.45, "humidity": 70, "fetched_time": 1706981285, "fetched_time_str": "2024-02-03 22:58:05"},
// 		"Riyadh" : {"temp": 287.23, "humidity": 27, "fetched_time": 1706981285, "fetched_time_str": "2024-02-03 22:58:05"},
// 		"Helsinki" : {"temp": 274.92, "humidity": 85, "fetched_time": 1706981285, "fetched_time_str": "2024-02-03 22:58:05"},
// 		"Delhi" : {"temp": 289.20, "humidity": 72, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Santiago" : {"temp": 305.77, "humidity": 22, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Manila" : {"temp": 299.64, "humidity": 78, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Brasília" : {"temp": 297.66, "humidity": 65, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Oslo" : {"temp": 276.73, "humidity": 68, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Warsaw" : {"temp": 281.80, "humidity": 75, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Budapest" : {"temp": 283.45, "humidity": 70, "fetched_time": 1706981286, "fetched_time_str": "2024-02-03 22:58:06"},
// 		"Riyadh" : {"temp": 287.23, "humidity": 27, "fetched_time": 1706981287, "fetched_time_str": "2024-02-03 22:58:07"},
// 		"Helsinki" : {"temp": 274.92, "humidity": 85, "fetched_time": 1706981287, "fetched_time_str": "2024-02-03 22:58:07"},
// 		"Delhi" : {"temp": 289.20, "humidity": 72, "fetched_time": 1706981287, "fetched_time_str": "2024-02-03 22:58:07"}
// 	 }
// 	 `;

// 	fmt.Println("jsonData to be stored in rabbitMq", tp.JsonData)

// 	status := tp.PingRabbitMq()

// 	if status == constants.RabbitMqDown {
// 		fmt.Println(constants.RabbitMqDown)
// 		return
// 	} else {
// 		fmt.Println(constants.RabbitMqUp)
// 		// below line to remove data from queue
// 		purge, error := purger.RemoveDataInQueue(constants.QueueName)
// 		if error == nil {
// 			fmt.Println(purge)
// 			errToPublish := tp.publishToRabbitMq(tp.JsonData, constants.QueueName)
// 			fmt.Print(errToPublish)
// 		}

// 	}
// }

func (tp *TemperaturePublisher) publishToRabbitMq(JsonData string, queueName string) string {
	conn, err := amqp.Dial(constants.RabbitMQURL)
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
		return constants.UnableToCreateQueue
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
	} else {
		// tp.consumeFromRabbitMq(channel)
	}

	return constants.SuccessfullyPublishedDataToQueue
}

func (tp *TemperaturePublisher) PingRabbitMq() string {
	conn, err := amqp.Dial(constants.RabbitMQURL)
	if err != nil {
		fmt.Println(err.Error())
		return constants.RabbitMqDown
	}
	defer conn.Close()

	return constants.RabbitMqUp
}

// func (tp *TemperaturePublisher) consumeFromRabbitMq(channel *amqp.Channel) {
// 	consumedMessages, errorWHileConsuming := channel.Consume(
// 		constants.QueueName, // Queue name
// 		"",                  // Consumer
// 		true,                // Auto-Acknowledge
// 		false,               // Exclusive
// 		false,               // No-Local
// 		false,               // No-Wait
// 		nil,                 // Args
// 	)
// 	if errorWHileConsuming != nil {
// 		fmt.Print("Failed to register a consumer: ", errorWHileConsuming)
// 	}

// 	fmt.Println("Waiting for messages. To exit press CTRL+C")
// 	for msg := range consumedMessages {
// 		fmt.Printf("Received a message: %s\n", msg.Body)
// 	}
// 	defer channel.Close()
// 	return
// }

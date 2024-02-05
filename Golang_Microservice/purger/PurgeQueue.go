package purger








import (
	"fmt"
	"log"

	"golang_weather_fetcher/myproject/constants"
	"github.com/streadway/amqp"
)

func PingRabbitMq() string {
	conn, err := amqp.Dial(constants.RabbitMQURL)
	if err != nil {
		fmt.Println(err.Error())
		return constants.RabbitMqDown
	}
	defer conn.Close()

	return constants.RabbitMqUp
}

func RemoveDataInQueue(queueName string) (string, error){
	// tp := &main.TemperaturePublisher{}
	status := PingRabbitMq()

	if status == constants.RabbitMqUp {
		conn, err := amqp.Dial(constants.RabbitMQURL)
		if err != nil {
			fmt.Println(err.Error())
			return "",err
		}
		defer conn.Close()

		channel, errConnectingToChannel := conn.Channel()
		if errConnectingToChannel != nil {
			fmt.Println(errConnectingToChannel.Error())
			return "",err
		}

		// Purge the queue

		_, err = channel.QueuePurge(queueName, false)
		if err != nil {
			log.Fatalf("Failed to purge the queue: %s", err)
			return "",err
		}

		
	}
	return queueName +" purged successfully", nil
}








package constants




const (
	RabbitMQURL                      = "amqp://guest:guest@localhost:5672/" // Replace with your RabbitMQ server URL
	QueueName                        = "temperature_data"
	RabbitMqDown                     = "RabbitMq down"
	RabbitMqUp                       = "RabbitMq up"
	UnableToCreateQueue              = "unable to create queue"
	UnableToPublishData              = "unable tp publis data"
	SuccessfullyPublishedDataToQueue = "successfully published data to queue"
);
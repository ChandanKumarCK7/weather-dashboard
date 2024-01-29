package com.Consumers;




import com.enums.RabbitQueueEnums;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.annotation.PostConstruct;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.boot.configurationprocessor.json.JSONException;
import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.boot.configurationprocessor.json.JSONStringer;
import org.springframework.stereotype.Component;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.Map;

@Component
public class RabbitMQConsumer {

//    private final Queue myQueue;
//
//    @Autowired
//    public RabbitMQConsumer(@Qualifier(RabbitQueueEnums.temperature_queue) Queue queue){
//        this.myQueue = queue;
//    }

    @PostConstruct
    public void postConstruct(){
        System.out.println("bean of the rabbit listener is initialized");
    }

    @RabbitListener(queues = RabbitQueueEnums.temperature_queue)
    public void handleMessage(String message) throws JsonProcessingException, JSONException {
        System.out.printf("recieved message from queue %s is %s%n", RabbitQueueEnums.temperature_queue, message.toString() );
        System.out.println("---------------");

        // convert to proper hashmap and json that can be reused.
        ObjectMapper objectMapper = new ObjectMapper();
        Map<String, Object> hashMap= objectMapper.readValue(message, Map.class);

        System.out.println("hashMap "+hashMap);

        // convert string to a json object
        JSONObject o = new JSONObject(message);

        System.out.println("jsonObject "+ o.toString());


    }


}

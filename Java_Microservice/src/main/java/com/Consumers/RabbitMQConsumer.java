package com.Consumers;




import com.Service.RapidStore;
import com.enums.RabbitQueueEnums;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.annotation.PostConstruct;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.configurationprocessor.json.JSONException;
import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.stereotype.Component;

import java.util.Map;

@Component
public class RabbitMQConsumer {


    JSONObject latestMessage;
    Map<String, Map<String, Object>> latestWeatherData;

    @Autowired
    RapidStore rapidStore;

    public JSONObject getLatestMessage() {
        return latestMessage;
    }

    public void setLatestMessage(JSONObject message){
        this.latestMessage= message;
    }

    public Map<String, Map<String, Object>> getLatestWeatherData() {
        return latestWeatherData;
    }

    public void setLatestWeatherData(Map<String, Map<String, Object>> latestWeatherData) {
        this.latestWeatherData = latestWeatherData;
    }



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
        TypeReference<Map<String, Map<String, Object>>> typeReference = new TypeReference<Map<String, Map<String, Object>>>() {};
        Map<String, Map<String, Object>> hashMap = objectMapper.readValue(message, new TypeReference<Map<String, Map<String, Object>>>() {});

        System.out.println("hashMap "+hashMap);

        // convert string to a json object

        JSONObject jsonObject = new JSONObject(message.toString());

        System.out.println("jsonObject "+ jsonObject.toString());

        setLatestMessage(jsonObject);
        setLatestWeatherData(hashMap);

        rapidStore.storeInBackgroundWithHashMap( hashMap);






    }


}

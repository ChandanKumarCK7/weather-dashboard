package com.Service;









import com.Consumers.RabbitMQConsumer;
//import com.Repository.TemperatureRepository;
import com.Model.TemperatureData;
import com.Repository.TemperatureRepository;
import org.bson.Document;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Service;


import java.util.Map;

@Qualifier(value = "rapidStore")
@Service
public class RapidStore {


    @Autowired
    TemperatureRepository temperatureRepository;

    @Autowired
    private MongoTemplate mongoTemplate;

    public void storeInBackgroundWithHashMap(Map<String, Map<String, Object>> hashMap) {

        System.out.println("store in background with hashmap" + hashMap);
        // linkedHashMap


        Document cityEntry_document = new Document(hashMap);

        System.out.println("size of the repository just before addition of records" +
                " "+temperatureRepository.count());
        for (Map.Entry<String, Map<String, Object>> entry : hashMap.entrySet()) {
            String cityName = entry.getKey().toString();
            TemperatureData.CityData cityData = new TemperatureData.CityData();

            entry.getValue().forEach((innerKey, innerValue) -> {
                switch (innerKey) {
                    case "temp":
                        cityData.setTemp(((Number) innerValue).doubleValue());
                    case "humidity":
                        cityData.setHumidity(((Number) innerValue).intValue());
                    case "fetched_time":
                        cityData.setFetched_Time(((Number) innerValue).longValue());
                    case "fetched_time_str":
                        cityData.setFetched_time_str(innerValue.toString());
                }
            });


            TemperatureData tempData = new TemperatureData(cityName, cityData);
            try{
                // clean up
//                performCleanupOnCityName(cityName);

                // rapid store
                TemperatureData savedData = temperatureRepository.save(tempData);
                System.out.println("Data saved successfully: " + savedData.toString());

            }
            catch(Exception e){
                System.out.println("exception occurred because of "+e.getMessage());
            }


            System.out.println(tempData.toString());

        }

        System.out.println("size of the repository after addition of records" +
                " "+temperatureRepository.count());
    }

    public void performCleanupOnCityName(String cityName){
        mongoTemplate.remove(Query.query(Criteria.where("city").is(cityName)), TemperatureData.class);
        System.out.println("size of the repository after cleanup of all cities "+temperatureRepository.count());
    }


}

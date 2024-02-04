package com.Model;







import org.bson.codecs.pojo.annotations.BsonCreator;
import org.bson.codecs.pojo.annotations.BsonId;
import org.bson.codecs.pojo.annotations.BsonProperty;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;


@Document(collection = "temperaturedata")
public class TemperatureData {


    @Id
    @BsonProperty("city")
    private String city;

    public TemperatureData(){

    }

    @BsonCreator
    public TemperatureData(String city, CityData cityData) {
        this.city = city;
        this.cityData = cityData;
    }


    private CityData cityData;

    public String getCity() {
        return city;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public CityData getCityData() {
        return cityData;
    }

    public void setCityData(CityData cityData) {
        this.cityData = cityData;
    }

    @Override
    public String toString() {
        return "TemperatureData{" +
                "city='" + city + '\'' +
                ", cityData=" + cityData.toString() +
                '}';
    }


    public static class CityData {
        @BsonProperty(value="fetched_time")
        private long fetched_Time;
        @BsonProperty(value="fetched_time_str")
        private String fetched_time_str;
        @BsonProperty(value="temp")
        private double temp;
        @BsonProperty(value="humidity")
        private int humidity;

        public CityData() {
        }

        @BsonCreator
        public CityData(long fetched_Time, String fetched_time_str, double temp, int humidity) {
            this.fetched_Time = fetched_Time;
            this.fetched_time_str = fetched_time_str;
            this.temp = temp;
            this.humidity = humidity;
        }

        public long getFetched_Time() {
            return fetched_Time;
        }

        public void setFetched_Time(long fetched_Time) {
            this.fetched_Time = fetched_Time;
        }

        public String getFetched_time_str() {
            return fetched_time_str;
        }

        public void setFetched_time_str(String fetched_time_str) {
            this.fetched_time_str = fetched_time_str;
        }

        public double getTemp() {
            return temp;
        }

        public void setTemp(double temp) {
            this.temp = temp;
        }

        public int getHumidity() {
            return humidity;
        }

        public void setHumidity(int humidity) {
            this.humidity = humidity;
        }

        @Override
        public String toString() {
            return "CityData{" +
                    "fetched_Time=" + fetched_Time +
                    ", fetched_time_str='" + fetched_time_str + '\'' +
                    ", temp=" + temp +
                    ", humidity=" + humidity +
                    '}';
        }
    }
}
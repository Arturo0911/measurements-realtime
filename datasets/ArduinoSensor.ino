#include <DHT.h>
#include <MQ135.h>

//definir el tipo de pin que ocupara cada sensor dentro de la placa Arduino
//DHT11
#define DHTPIN 2
#define DHTTYPE DHT11
DHT dht(DHTPIN, DHTTYPE);

//HW-103 o sensor de humedad de suelo
int pinHumedity = 0;

//MQ-135 o sensor de calidad de aire (CO2)
int PIN_MQ135 = 3;
MQ135 sensorMQ(PIN_MQ135);


void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  dht.begin();
}

void loop() {
  // put your main code here, to run repeatedly:
  
  float temp = dht.readTemperature();
  float aire = analogRead(PIN_MQ135);
  float hume = analogRead(pinHumedity);
  float humedity = ((((hume*100)/1023)-100)*-1);
  
  //Serial.print("Temp(C): ");
  Serial.println(temp);
  //Serial.print(" °C");
  //Serial.println(" ");
                  
  //Serial.print("Co2(ppm): ");
  Serial.println(aire);
  //Serial.print(" ppm");
  //Serial.println(" ");
      
  //Serial.print("Humd.Suelo: ");
  Serial.println(humedity);
  //Serial.print(" %");
  //Serial.println(" ");

  delay(2000);

}

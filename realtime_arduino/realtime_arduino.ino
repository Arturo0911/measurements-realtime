
const int sensorHumidity = A0;
const float y = 1.43457;


void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);

}

void loop() {
  // put your main code here, to run repeatedly:
  
  float x;
  float saliny;
  float z ;
  int humidity = analogRead(sensorHumidity);

  x = (humidity *0.5)/1000;
  saliny = (humidity - x) *y;
  z = map(saliny, 0, 1023, 100, 0);
  delay(1000);
  Serial.println(z);

}

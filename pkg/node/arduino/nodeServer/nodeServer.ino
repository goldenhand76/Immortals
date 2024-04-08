#include <ArduinoJson.h>
#include <ESP8266mDNS.h>
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <PubSubClient.h>
#include <EEPROM.h>

// Set WiFi credentials
#define WIFI_SSID "IoT"
#define WIFI_PASS "Thereisnoway97"
#define EEPROM_SIZE 16

WiFiServer server(80);
WiFiClient wclient;
PubSubClient mqttClient(wclient);

const int WHITE = 14;  
const int RED = 12;
const int GREEN = 13;
const int BLUE = 4;

const int   mqttport = 1883;
String prefix;
String brokerIP;
String clientID = "";
String topic = "";
String message = "";
char readChar = ' ';
const char* error;
JsonDocument topics;

const size_t JSON_DOCUMENT_SIZE = 1024; // Adjust the size according to your JSON payload

struct clientData {
  DynamicJsonDocument* data;
  const char* error;
  clientData(){
    data = new DynamicJsonDocument(JSON_DOCUMENT_SIZE);
    error = error;
  }
};

void setup_wifi(){
  // Begin WiFi
  WiFi.begin(WIFI_SSID, WIFI_PASS);
  WiFi.mode(WIFI_STA);
  WiFi.setSleep(false);
  Serial.print("Connecting to ");
  Serial.print(WIFI_SSID);
  
  // Loop continuously while WiFi is not connected
  while (WiFi.status() != WL_CONNECTED){
    delay(1000);
    Serial.print(".");
  }

  // Connected to WiFi
  Serial.println();
  Serial.print("Connected! IP address: ");
  Serial.println(WiFi.localIP());
}

void setup_mqtt(){
  clientID = ESP.getChipId();
  Serial.print("Node Client ID : ");
  Serial.println(clientID);
  prefix = clientID + "/";
  
  brokerIP = readStringFromEEPROM(0);
  Serial.print("MQTT Broker IP : ");
  Serial.println(brokerIP);

  // Setup mqtt client
  if (brokerIP.length() != 0){
    mqttClient.setServer(brokerIP.c_str(), mqttport);
    mqttClient.setCallback(onMessageArrived);
  }else{
    Serial.println("MQTT Setup failed, Invalid IP Address");
  }
}

void setup_http(){
  // Set up mDNS responder
  if (!MDNS.begin(clientID)) {
  Serial.println("Error setting up MDNS responder!");
  }
  else {
  Serial.println("mDNS responder started");
  }
  server.begin();
  Serial.println("TCP server started");
  MDNS.addService("http", "tcp", 80);
}

void setup(){
  Serial.begin(115200);
  EEPROM.begin(EEPROM_SIZE);
  delay(3000);
  setup_mqtt();
  setup_wifi();
  setup_http();
  pinMode(RED, OUTPUT);
  pinMode(GREEN, OUTPUT);
  pinMode(BLUE, OUTPUT);
  pinMode(WHITE, OUTPUT);

  topics["sensor"]["RED"] = prefix + "sensor/Red";
  topics["sensor"]["GREEN"] = prefix + "sensor/Green";
  topics["sensor"]["BLUE"] = prefix + "sensor/Blue";
  topics["actuator"]["RED"] = prefix + "actuator/Red";
  topics["actuator"]["GREEN"] = prefix + "actuator/Green";
  topics["actuator"]["BLUE"] = prefix + "actuator/Blue";
        
}

void loop(){
  MDNS.update();
  // Check if a client has connected
  WiFiClient httpClient = server.accept();
  if (httpClient) {
    Serial.println("New Client.");
    HttpHandler(httpClient);
  }

  if (!mqttClient.connected()){
    Serial.println("Connecting To Server ");
    if (mqttClient.connect(clientID.c_str()))
    {
      Serial.println("Connected To Server");
      mqttClient.subscribe(topics["actuator"]["RED"], 1);
      mqttClient.subscribe(topics["actuator"]["GREEN"], 1);
      mqttClient.subscribe(topics["actuator"]["BLUE"], 1);
    } 
    else 
    {
      Serial.print("failed with state : ");
      Serial.println(mqttClient.state());
      delay(1000);
    }
  }
  mqttClient.loop();
  Rainbow();
  mqttClient.publish(topics["sensor"]["RED"], "This is a message from red");
  mqttClient.publish(topics["sensor"]["GREEN"], "This is a message from green");
  mqttClient.publish(topics["sensor"]["BLUE"], "This is a message from blue");
}
  
void HttpHandler(WiFiClient client)
{
  while (client.connected() && !client.available()) { delay(1); }
  // Read the first line of HTTP request
  String req = client.readStringUntil('\r');
  String path = parseHttpPath(req);

  if (path == "/ip") {
    req = client.readString();
    struct clientData resp = parseHttpData(req);
    if(resp.error){
      JsonDocument doc;
      doc["error"] = resp.error;
      serializeJson(doc, Serial);
      client.print("HTTP/1.1 400 Bad Request");
      client.println(F("Content-Type: application/json"));
      client.println(F("Connection: close"));
      client.print(F("Content-Length: "));
      client.println(measureJsonPretty(doc));
      client.println();
      serializeJsonPretty(doc, client);
      Serial.println("Sending 400");
    
    }else {
        const char* brokerIP = (*resp.data)["brokerIP"];
        Serial.print("Broker IP: ");
        Serial.println(brokerIP);
        writeStringToEEPROM(0, brokerIP);
        
        serializeJson(topics, Serial);
        client.println(F("HTTP/1.0 200 OK"));
        client.println(F("Content-Type: application/json"));
        client.println(F("Connection: close"));
        client.print(F("Content-Length: "));
        client.println(measureJsonPretty(topics));
        client.println();
        serializeJsonPretty(topics, client);
        Serial.println("Sending 200");
        
    }
  } else {
    client.print("HTTP/1.1 404 Not Found\r\n\r\n");
    Serial.println("Sending 404");
  }
  
  client.flush();
}

String parseHttpPath(String request){
  Serial.println(request);
  int addr_start = request.indexOf(' ');
  int addr_end = request.indexOf(' ', addr_start + 1);
  if (addr_start == -1 || addr_end == -1) {
    Serial.print("Invalid request: ");
    Serial.println(request);
    return "";
  }
  return request.substring(addr_start + 1, addr_end);
}

struct clientData parseHttpData(String request) {
  int jsonStart = request.indexOf("{");
  struct clientData resp;
  
  if(jsonStart != -1){
    // Extract the JSON payload
    String jsonString = request.substring(jsonStart);
    DeserializationError error = deserializeJson(*resp.data, jsonString);
    
    // Check for parsing errors
    if (error) {
      Serial.print("Failed to parse JSON: ");
      Serial.println(error.c_str());
      resp.error = error.c_str();
      return resp;
    }

    if ((*resp.data).containsKey("brokerIP") != true){
      const char* error = "brokerIP payload not found";
      resp.error = error;
      return resp;
    }
    
    return resp;
  } else {
    Serial.println("JSON payload not found in HTTP response");
    const char* error = "Json payload not found";
    resp.error = error;
    return resp;
  }
}

void writeStringToEEPROM(int addrOffset, const String &strToWrite)
{
  for (int index = 0; index < strToWrite.length(); index++) {
    EEPROM.write(addrOffset, strToWrite[index]);
    EEPROM.commit();
    delay(10);
    Serial.print("Writing ");
    Serial.print(strToWrite[index]);
    Serial.print(" in address ");
    Serial.println(addrOffset);
    addrOffset++;
  }
}

String readStringFromEEPROM(int addrOffset)
{
  while (readChar != '\0') {
    readChar = EEPROM.read(addrOffset);
    delay(10);
    if (readChar != '\0') {
      brokerIP += readChar;
    }
    addrOffset++;
  }
  return brokerIP;
}

char* toCharArray(String str) {
  return &str[0];
}

void onMessageArrived(char* t, byte* m, unsigned int length) {
  topic = String(t);
  message = String((char*)m);
  message = message.substring(0, length);
  Serial.println(message);
}




void Rainbow(){
  // Red
  analogWrite(RED, 255);
  analogWrite(GREEN, 0);  
  analogWrite(BLUE, 0);
  delay(90);
  // Orange
  for (int i = 0; i <= 127 ; i += 1){
    analogWrite(GREEN, i);
    delay(5);
  }
  delay(90);
  // Yellow
  for (int i = 127; i <= 255 ; i += 1){
    analogWrite(GREEN, i);
    delay(5);
  }
  delay(90);
  //Green
  for (int i = 255; i >= 127; i -= 1){
    analogWrite(RED, i);
    delay(5);
  }
  delay(90);
  //Dark Green
  for (int i = 127; i >= 0; i -= 1){
    analogWrite(RED, i);
    delay(5);
  }
  delay(90);
  //Light Green
  for (int i = 0; i <= 127; i += 1){
    analogWrite(BLUE, i);
    delay(5);
  }
  delay(90);
  //Cyan Blue
  for (int i = 127; i <= 255; i += 1){
    analogWrite(BLUE, i);
    delay(5);
  }
  delay(90);
  //Light Blue
  for (int i = 255; i >= 127; i -= 1){
    analogWrite(GREEN, i);
    delay(5);
  }
  delay(90);
  //Dark Blue
  for (int i = 127; i >= 0; i -= 1){
    analogWrite(GREEN, i);
    delay(5);
  }
  delay(90);
  //Purple
  for (int i = 0; i <= 127; i += 1){
    analogWrite(RED, i);
    delay(5);
  }
  delay(90);
  //Pink
  for (int i = 127; i <= 255; i += 1){
    analogWrite(RED, i);
    delay(5);
  }
  delay(90);
  //Rosy
  for (int i = 255; i >= 127; i -= 1){
    analogWrite(BLUE, i);
    delay(5);
  }
  delay(90);
  //Red
  for (int i = 127; i >= 0; i -= 1){
    analogWrite(BLUE, i);
    delay(5);
  }
}

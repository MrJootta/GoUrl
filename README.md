# GoUrl

Simple too to make url shorter

### Concept

This application allows you to send a full url and get back a shorter version 

### Endpoints

**POST /create**

Endpoint to create new short url's it receives and returns json

_Request_

```
{   
    "url": "example.com"
}
```

_Response_

```
{
    "status": 201,
    "response": {
        "url": "http://localhost:2020/KaXTsk",
        "code": "KaXTsk"
    }
}
```

_Status Codes_

* 500 - When code is not generated or can't save in database
* 405 - When method is not GET
* 400 - When url param is not sent or url can't be parsed
* 201 - When short url is created

**GET /info/{code}**

Endpoint to consult number of visits a code got in the last 24h

_Response_

```
{
    "status": 200,
    "response": {
        "url": "http://localhost:2020/KaXTsk",
        "code": "KaXTsk",
        "number_of_visits": 3
    }
}
```

_Status Codes_

* 405 - When method is not GET
* 404 - When code is not found
* 200 - When returns visits count

**GET /{code}**

Endpoint to be redirected to the saved long url

_Status Codes_

* 405 - When method is not GET
* 404 - When code is not found
* 301 - When redirect happen

### Setup

To setup the application you just need to run the following commands

```
$ -> docker-compose up
$ -> go mod tidy
$ -> go run main.go
```

### Current implementation

##### Services

- Mysql database
- Go Application

### Ideal Implementation

##### Services

- Mysql database
- Memcached
- Go Application
- RabbitMQ

_MySql_

To store the application data, preferably with a master/slave architecture

_Memcached_ 

To store the keys code/url to speed up the access to the URL when redirect call happen

_RabbitMQ_

To deal with process's that can be done using a event driven pattern like when we count the visits, visit count and save is not priority feature for the app and can be done with sime delay

##### Metrics

_OpenTracing_

The new OpenTelemetrics is still in beta so would be using OpenTrancing for the metrics

_Circuit Breaker_

Would use a circuit breaker like [hystrix](https://godoc.org/github.com/afex/hystrix-go/hystrix) to be a layer between application and database/queue

_Monitoring_ 

Besides OpenTracing i'm very found of [Sentry](https://sentry.io), also using jaeger to connect with OpenTracing and grafana for the dashboards





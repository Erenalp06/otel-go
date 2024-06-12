# Monitoring a Fiber Go Application with OpenTelemetry

This guide explains how to run your Fiber Go application integrated with Jaeger, Kibana, and Elasticsearch using Docker.

## Prerequisites

To run this application locally, ensure you have the following tools installed:

- Docker
- Docker Compose
- Jaeger Collector
- ElasticSearch & Kibana

## Configuration and Setup

The application is configured to send tracing data to a Jaeger instance. Jaeger stores the tracing data in Elasticsearch, and Kibana visualizes these data.

### Starting Services with Docker Compose

Use the following `docker-compose.yml` file to start Elasticsearch, Kibana, and Jaeger services.

```yaml
version: '3.7'

services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:7.9.3
    container_name: elasticsearch
    networks:
      - shared_network
    environment:
      - discovery.type=single-node
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
      - network.host=0.0.0.0için
    ports:
      - "9200:9200"
      - "9300:9300"
    healthcheck:
      test: ["CMD-SHELL", "curl --silent --fail localhost:9200/_cluster/health?wait_for_status=yellow&timeout=50s || exit 1"]
      interval: 30s
      timeout: 10s
      retries: 5

  kibana:
    image: docker.elastic.co/kibana/kibana:7.9.3
    container_name: kibana
    networks:
      - shared_network
    ports:
      - "5601:5601"
    depends_on:
      elasticsearch:
        condition: service_healthy
    environment:
      - ELASTICSEARCH_URL=http://elasticsearch:9200

  jaeger:
    image: jaegertracing/all-in-one:latest
    container_name: jaeger-es
    networks:
      - shared_network
    ports:
      - "5775:5775/udp"
      - "6831:6831/udp"
      - "6832:6832/udp"
      - "5778:5778"
      - "16686:16686"
      - "14268:14268"
      - "14250:14250"
      - "9411:9411"
      - "4317:4317"
      - "4318:4318"
    environment:
      - COLLECTOR_ZIPKIN_HTTP_PORT=9411
      - SPAN_STORAGE_TYPE=elasticsearch
      - ES_SERVER_URLS=http://elasticsearch:9200
      - COLLECTOR_OTLP_ENABLED=true
    depends_on:
      elasticsearch:
        condition: service_healthy

networks:
  shared_network:
    external: true

```

Save this configuration as a docker-compose.yml file and run the following command to start all services:

```bash
docker-compose up
```

## Running the Application

To run the application using Docker Compose, follow these steps:

### 1.  Clone the Repository

```bash
git clone <repository_url>
cd <repository_directory>
```

### 3. Start the Application with Docker Compose

Ensure you have `make` installed on your system.

Navigate to the directory where your application is located and run the following command to compile and start your application in the background:


```bash
make run
```

## API Endpoints
The application provides the following API endpoints:

- `GET /api/v1/users` - GetAllUsers
- `GET /api/v1/users/{id}` - GetUserById
- `POST /api/v1/users` - CreateUser
- `PUT /api/v1/users/{id}` - UpdateUser
- `DELETE /api/v1/users/{id}` - - DeleteUser

## Testing the Application

Once the application is running, you can test the API using curl or any API testing tool:

### GET Request
```bash
curl -X GET http://localhost:8085/api/v1/users
```

### POST Request
To create a new user, you can use the following command:

```bash
curl -X POST http://localhost:8085/api/v1/users \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "date": "2024-06-09",
  "city": "New York",
  "country": "USA"
}'
```

##  Monitoring with Jaeger

With the application running, tracing data will be automatically sent to Jaeger. You can access the Jaeger UI at http://localhost:16686 to visualize your tracing data.

## Visualization with Kibana

Access Kibana at http://localhost:5601 to create visualizations and analyze the data stored in Elasticsearch.

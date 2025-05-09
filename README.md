### Test Web App

### Stack
- Golang (Echo)
- Docker

### Run

To build and run the project, use the command below: 
```
docker build -t test-web-app:latest . && docker run -p 127.0.0.1:8080:8080 test-web-app:latest
```

The endpoint should now be accessible from making a POST request to the below url

```
http://127.0.0.1:8080/ticket
```

alternatively you can run this cURL request and you should receive the correct itinerary from the starting airport to the last airport based in the json payload.
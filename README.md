### Test Web App
-  The app can be run using Docker which allows it to have a more consistent runtime environment both locally and on a different machine. Since the app is small the only library used was the Echo web framework which provides a great middleware and router package to the app. To allow for better testing the itinerary ordering function is separated from the http POST request for better testing coverage and separation of concerns during testing.

### Prerequisites
- Go 1.24+ (Echo)
- Docker

### Installation
1. Clone the repo
```
git clone https://github.com/AllanM007/test-web-app
```
2. Navigate to project directory
```
cd test-web-app
```
3. Build and run Docker container
```
docker build -t test-web-app:latest . && docker run -p 127.0.0.1:8080:8080 test-web-app:latest
```

### Tests
- The project has unit tests which can be run using the below command:
```
go test -v
``` 

### Usage

The app has a single endpoint which is accessible from making a POST request to the below url:

```
http://127.0.0.1:8080/tickets/order
```
using the example payload:
```json
[
  ["LAX", "DXB"],
  ["JFK", "LAX"],
  ["SFO", "SJC"],
  ["DXB", "SFO"]
]
```

alternatively one can run this cURL request and should receive the correct itinerary from the starting airport to the last airport based on the json payload

```json
curl --location 'http://127.0.0.1:8080' \
--header 'Content-Type: application/json' \
--data '[
  ["LAX", "DXB"],
  ["JFK", "LAX"],
  ["SFO", "SJC"],
  ["DXB", "SFO"]
]'
```
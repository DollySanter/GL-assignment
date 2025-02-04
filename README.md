# GL-assignment
Weather Api Assignment

# Weather API Server

This project is a simple Go server that fetches weather data based on the user's location and displays it.

## Prerequisites

- Go installed on your machine.
- An API key from OpenWeatherMap (sign up [here](https://openweathermap.org/api)).

## Running the Server

1. Open a terminal and navigate to the directory where `weather_api_server.go` is located.

2. Run the server using the following command:

   ```bash
   $ go run weather_api_server.go
    ```

You should see something like:
starting server at 8080
This indicates that the server is up and running locally on port 8080.

# Calling the Weather API Handler

To get the current weather, you can use the following URL (replacing api_key with your actual API key):

$ http://localhost:8080/weather?apiKey=api_key&part=current
Or 
curl "http://localhost:8080/weather?apiKey=api_key&part=current"

# Running Computations

A console tool for calculating running pace and planning running strategies.

## Description

Running Computations is a Go application that allows:

- Calculating running pace based on distance and time
- Generating split times for a selected distance
- Creating running strategies, including negative split strategy (slower first part, faster second part)
- Supporting multiple time and distance formats
- Using predefined distances (5km, 10km, half marathon, marathon)

## Installation

```bash
git clone https://github.com/piotrowerko/running_computations.git
cd running_computations
go build -o running_cli cmd/cli/main.go
go build -o running_api cmd/api/main.go
```

## CLI Usage

```bash
# Calculate pace for custom distance and time
./running_cli -distance=3.3 -time=1200

# Using HH:MM:SS time format
./running_cli -distance=3.3 -timeformat=00:20:00

# Using predefined distances
./running_cli -preset=10k -timeformat=00:45:00

# Using custom interval for timestamps (every 0.5 km)
./running_cli -preset=5k -timeformat=00:25:00 -interval=0.5

# Negative split strategy (slower start, faster finish)
./running_cli -preset=10k -timeformat=00:45:00 -negativesplit -splitdistance=60 -pacedifference=3
```

### Available CLI Parameters

| Parameter | Description | Default Value |
|-----------|-------------|---------------|
| -distance | Custom distance in kilometers | - |
| -time | Time in seconds | - |
| -timeformat | Time in HH:MM:SS format | - |
| -interval | Interval in kilometers for timestamps | 1.0 |
| -preset | Predefined distance (5k, 10k, half, marathon) | - |
| -negativesplit | Use negative split pacing strategy | false |
| -splitdistance | Percentage of distance for split point | 50 |
| -pacedifference | Percentage difference between slower and faster pace | 5 |

## API Usage

The application also provides a REST API that can be used to perform the same calculations.

```bash
# Start the API server
./running_api
```

The server will start listening on port 8080.

### API Endpoints

#### Health Check
```bash
curl http://localhost:8080/
```

#### Pace Calculation
The API provides a `/pace` endpoint that accepts POST requests with JSON data.

##### Example 1: Basic pace calculation for 10km with a time of 45 minutes
```bash
curl -X POST http://localhost:8080/pace \
  -H "Content-Type: application/json" \
  -d '{
    "distance": 10.0,
    "timeFormat": "00:45:00"
  }'
```

##### Example 2: Using negative split strategy for a half marathon
```bash
curl -X POST http://localhost:8080/pace \
  -H "Content-Type: application/json" \
  -d '{
    "preset": "half",
    "timeFormat": "01:45:00",
    "interval": 1.0,
    "negativeSplit": true,
    "splitDistance": 60,
    "paceDifference": 5
  }'
```

##### Example 3: Using time in seconds and custom distance
```bash
curl -X POST http://localhost:8080/pace \
  -H "Content-Type: application/json" \
  -d '{
    "distance": 5.0,
    "timeInSeconds": 1500,
    "interval": 0.5
  }'
```

## Predefined Distances

- 5k: 5.000 km
- 10k: 10.000 km
- half: 21.097 km (half marathon)
- marathon: 42.195 km (marathon)

## Example Results

```
Using preset distance: 10.000 km
Provided time: 00:45:00 (2700 seconds)
Running pace: 4 min 30 sec / km

Pacing strategy: Even pace strategy
Times at each interval:
1.0 km: 00:04:30
2.0 km: 00:09:00
3.0 km: 00:13:30
4.0 km: 00:18:00
5.0 km: 00:22:30
6.0 km: 00:27:00
7.0 km: 00:31:30
8.0 km: 00:36:00
9.0 km: 00:40:30
10.0 km: 00:45:00
```

## License

This project is available under the terms specified in the [license](LICENSE).

## Author

Piotr Owerko 
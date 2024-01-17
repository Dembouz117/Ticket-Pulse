# Queue Microservice

### Endpoints

#### Randomise Waiting Room and Add to Queue

- Endpoint: /randomise
- Method: GET
- Description: Generates randomise user, adds them to the waiting room, retrieves all users from the waiting room, randomises their order and add them the queue

#### Reset System

- Endpoint: /reset
- Method: GET
- Description:  Resets the system, clears the waiting room, queue, and related data.

## Redis commands

<!-- Get set size -->
SCARD waiting-room

<!-- Get set members -->
SMEMBERS waiting-room

<!-- Get sorted set rank -->
ZRANK queue <milton@gmail.com>

<!-- Get sorted set members -->
ZRANGE queue 0 -1

SET concert-status waiting

<!-- Remove first from queue -->
ZPOPMIN queue
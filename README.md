Web Server: Golang vs NodeJS
================
# Run the web servers

    docker-compose up

Go web server run at `:8000`, Node run at `:3000`.

# Comparing:
## File reading
Quite the same.

1000 requests, 10 concurrent level (sorry, the image has typo, "Golan" should be "Golang").

<img width="1680" alt="1000 requests" src="https://cloud.githubusercontent.com/assets/4214509/19913719/ab3d4efe-a0d7-11e6-843a-69a3b408d862.png">

10000 request, 100 concurrent level

<img width="1680" alt="10000 requests" src="https://cloud.githubusercontent.com/assets/4214509/19913720/ab3fad66-a0d7-11e6-949f-a41567ee870b.png">

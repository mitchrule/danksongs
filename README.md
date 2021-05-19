# [danksongs](https://danksongs.com/)

[![Build Status](https://travis-ci.com/mitchrule/danksongs.svg?branch=main)](https://travis-ci.com/mitchrule/danksongs)

## Setup and Run

First make sure you have [Docker](https://www.docker.com/get-started) installed.
Once you have Docker installed, Danksongs can be built by running `docker-compose up --build`

To run again without having to rebuild, use `docker-compose up`

Shutdown the containers with `docker-compose down`

## Spotify authentication

Inorder for the spotify API querys to work, a ClientID and Secret key is required that can be accessed and created at
(https://developer.spotify.com/dashboard/applications). From there once they are defined in a root level .env file as
SPOTIFY_ID and SPOTIFY_SECRET the API should then function as expected. In the future we'll try to integrate this
into docker.

## Access

Once the Docker container is running, the containers can be accessed as follows:

- UI <http://localhost:3000>
- Database <http://localhost:8081>
- API <http://localhost:8080>

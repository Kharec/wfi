#!/bin/bash
docker run -d --name mariatest -p 3306:3306 --env MARIADB_ROOT_PASSWORD=my-secret-pw  mariadb:latest
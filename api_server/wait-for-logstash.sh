#!/bin/bash

## Or whatever command is used for checking logstash availability
until curl 'http://logstash:9600' > /dev/null; do
  echo "Waiting for logtash..."
  sleep 3;
done

# Start your server
fresh
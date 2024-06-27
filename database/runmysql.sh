#!/bin/bash

while [ "$command" != "stop" ] && [ "$command" != "start" ] && [ "$command" != "purge" ]; do
    echo "Commands available:"
    echo "start: To start Mysql"
    echo "stop: To stop Mysql (remove container too)"
    echo "purge: To purge Mysql (remove container and data too)"
    echo ""
    echo "Enter the name of the command and press [ENTER]: "

    read command
done

if [  "$command" == "purge" ]; then
  docker-compose stop
  docker rm --force cn_mysql
  rm -rf mysql
  exit 0
fi

if [  "$command" == "stop" ]; then
  docker-compose stop
  docker rm --force cn_mysql
  exit 0
fi

if [  "$command" == "start" ]; then
  docker-compose up
fi


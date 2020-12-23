#!/bin/bash
go build -i -o ./bin/notify matthew-cash.com/notification-sender
sudo cp ./bin/notify /usr/local/bin/notify
sudo chmod +x /usr/local/bin/notify
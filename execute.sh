#!/bin/bash

# Put your ngrok auth token here
ngrok config add-authtoken {SET_YOUR_NGROK_AUTH_TOKEN_HERE}

# Application build and run
cd cmd/
go build main.go
chmod +x main
./main
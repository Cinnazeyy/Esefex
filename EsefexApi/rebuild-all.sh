#!/bin/bash

sudo docker build -t jokil/esefexapi:base -f Dockerfile.base .
sudo docker build -t jokil/esefexapi .
sudo docker build -t jokil/esefexapi:pterodactyl -f ./Dockerfile.pterodactyl .
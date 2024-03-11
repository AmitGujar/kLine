#!/bin/bash

sudo mkdir -p -m 755 /etc/apt/keyrings && wget -qO- https://kline.akstest.tech/kline/kline-stable-amd64.gpg | sudo tee /etc/apt/keyrings/kline-stable-amd64.gpg >/dev/null &&
    sudo chmod go+r /etc/apt/keyrings/kline-stable-amd64.gpg &&
    echo "deb [signed-by=/etc/apt/keyrings/kline-stable-amd64.gpg] https://kline.akstest.tech/kline/ ./" | sudo tee /etc/apt/sources.list.d/kline-repo.list >/dev/null &&
    sudo apt update &&
    sudo apt install kline

# kLine

## How to install

```
sudo mkdir -p -m 755 /etc/apt/keyrings && wget -qO- https://kline.akstest.tech/kline/kline-stable-amd64.gpg | sudo tee /etc/apt/keyrings/kline-stable-amd64.gpg >/dev/null &&
sudo chmod go+r /etc/apt/keyrings/kline-stable-amd64.gpg &&
echo "deb [signed-by=/etc/apt/keyrings/kline-stable-amd64.gpg] https://kline.akstest.tech/kline stable main" | sudo tee /etc/apt/sources.list.d/kline-repo.list >/dev/null &&
sudo apt update &&
sudo apt install kline
```

## Requirements

1. Following environment variables
```
AZURE_SUBSCRIPTION_ID
AZURE_RESOURCE_GROUP_NAME
```
2. Connection with the k8s cluster.

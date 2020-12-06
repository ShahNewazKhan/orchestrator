## Dependencies

### Brig cli 

Brig is the Brigade command line client. You can use brig to create/update/delete new brigade Projects, run Builds, etc. To get `brig`, navigate to the [Releases](https://github.com/brigadecore/brigade/releases/) page and then download the appropriate client for your platform. For example, if you’re using Linux or WSL, you can get the 1.4.0 version in this way:

```sh
# Note the k8s client used in brig < 1.4.0 is not compatible with k8s >= 1.18
wget -O brig https://github.com/brigadecore/brigade/releases/download/v1.4.0/brig-linux-amd64
chmod +x brig
sudo mv brig /usr/local/bin/
```

## Create brig project 

```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "key1": "value1",
    "key2": "value2"
}' \
  http://localhost:8081/simpleevents/v1/brigade-11376b7e2acf0c93001f16d10c7ef82ecc88f2292851ba9a5fb147/4cFaM
```
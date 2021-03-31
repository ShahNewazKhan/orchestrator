# Orchestrator 

## Dependencies

[fswatch](https://github.com/emcrisostomo/fswatch) is a file change monitor that receives notifications when the contents of the specified files or directories are modified. fswatch implements several monitors. It is used to do a live reload of the of the `go` codebase during local development.

Follow [install](https://github.com/emcrisostomo/fswatch#installation) instructions for your system. 


## Endpoints

|Method| Endpoint | Description|
|---|---|---|
|`GET`    | `/api/jobs` | lists all job in the database |
|`GET`    | `/api/jobs/:id`| retrieves a `job` based on its ID |
|`POST`   | `/api/jobs` | creates a `job` object and adds it to the database |
|`PATCH`  | `/api/jobs/:id` | changes the `Completed` property of a `job` based on its ID |
|`PATCH`  | `/api/jobs/:id/brigade` | adds the associated brigade `buildId` & `workerId` from the launched process  |
|`PATCH`  | `/api/jobs/:id/status` | changes the job status to one of ['PENDING','STARTED','RUNNING','DONE','ERRORED'] |
|`DELETE` | `/api/jobs/:id` | deletes a `job` based on its ID |


## Envars

|Envar|Default|Description|
|---|---|---|
| `MONGODB_USER` | `root` | Mongo user|
| `MONGODB_PASS` | `rootpassword` | Mongo user password|
| `MONGODB_HOST` | `0.0.0.0` |Mongo host|
| `MONGODB_PORT` | `27017` | Mongo port|

## Launch dev

```sh

go run server.go
```

```sh
curl --request GET '0.0.0.0:3000/api/projects'

curl --request POST '0.0.0.0:3000/api/jobs' \
--header 'Content-Type: application/json' \
--data-raw '{
      "brigadeProject": "brigade-1ac72d272cbb901c97f62f326939588d8eb5fe33e63c1fc467a8d1",
      "name": "coeus/coeus-engine ",
      "brigadeSecret": "genericsecret",
      "videoUrl": "https://www.youtube.com/watch?v=6t_ib-rxKUM"
    }
' --insecure
```

## Initiate a coeus-engine processing job

### GET /api/projects
```sh
curl --request GET 'https://orchestrator.skynet2.shahnewazkhan.usw1.k8g8.com/api/projects' --insecure

# RESPONSE 200 ok

{
    "ok":true,
    "projects":[
        {
            "id":"brigade-1ac72d272cbb901c97f62f326939588d8eb5fe33e63c1fc467a8d1",
            "name":"coeus/coeus-detector",
            "genericGatewaySecret":"genericsecret"
        }
    ]
}
```
### POST /api/jobs 
```sh
# populate brigateProject, name & secret from response object for /api/projects
curl --request POST 'https://orchestrator.skynet2.shahnewazkhan.usw1.k8g8.com/api/jobs' \
--header 'Content-Type: application/json' \
--data-raw '{
      "brigadeProject": "brigade-1ac72d272cbb901c97f62f326939588d8eb5fe33e63c1fc467a8d1",
      "name": "coeus/coeus-detector",
      "brigadeSecret": "genericsecret",
      "videoUrl": "https://www.youtube.com/watch?v=6t_ib-rxKUM"
    }
' --insecure

REPONSE 200 OK: 

{
    "job":
    {
        "_id":"60582285ec6b565ae5d73bc8",
        "created_at":"2021-03-22T04:52:21.998798079Z",
        "updated_at":"2021-03-22T04:52:21.99880005Z",
        "status":"PENDING",
        "name":"coeus/coeus-engine ",
        "complete":false,
        "buildId":"",
        "workerId":""
    },
    "ok":true
}
```


```sh
curl 'https://orchestrator.skynet.shahnewazkhan.usw1.k8g8.com/api/jobs/60583ed0fae81b059f6dd85d' --insecure
```

```sh
curl '0.0.0.0:3000/api/jobs/5fe536acb9c14f2b9de01736/status' \
    --request PATCH \
    --header 'Content-Type: application/json' \
    --data-raw '{"status":"STARTED"}' | jq .
```

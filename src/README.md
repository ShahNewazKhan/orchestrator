# Orchestrator 

## Dependencies

[fswatch](https://github.com/emcrisostomo/fswatch) is a file change monitor that receives notifications when the contents of the specified files or directories are modified. fswatch implements several monitors. It is used to do a live reload of the of the `go` codebase during local development.

Follow [install](https://github.com/emcrisostomo/fswatch#installation) instructions for your system. 


## Endpoints

|Method| Endpoint | Description|
|---|---|---|
|`GET`| `/api/jobs` | lists all job in the database |
|`GET`| `/api/jobs/:id`| retrieves a `job` based on its ID |
|`POST`| `/api/jobs` | creates a `job` object and adds it to the database |
|`PATCH` | `/api/jobs/:id` | changes the `Completed` property of a `job` based on its ID |
|`PATCH` | `/api/jobs/:id/brigade` | adds the associated brigade `buildId` & `workerId` from the launched process  |
|`PATCH` | `/api/jobs/:id/status` | changes the job status to one of ['PENDING','STARTED','RUNNING','DONE','ERRORED'] |
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
curl '0.0.0.0:3000/api/jobs/5fe536acb9c14f2b9de01736/status' \
    --request PATCH \
    --header 'Content-Type: application/json' \
    --data-raw '{"status":"STARTED"}' | jq .
```
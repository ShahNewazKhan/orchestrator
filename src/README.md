# Orchestrator 

## Endpoints

|Method| Endpoint | Description|
|---|---|---|
|`GET`| `/api/jobs` | lists all job in the database |
|`GET`| `/api/jobs/:id`| retrieves a `job` based on its ID |
|`POST`| `/api/jobs` | creates a `job` object and adds it to the database |
|`PATCH` | `/api/jobs/:id` | changes the `Completed` property of a `job` based on its ID |
|`DELETE` | `/api/jobs/:id` | deletes a `job` based on its ID |


## Envars

|Envar|Default|Description|
|---|---|---|

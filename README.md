# stacky
A hosted service to dynamically choose between multiple API instances

## Configuration
You have to define 3 environment variables to configure your Stacky instance:
1. The `STACKY_MONGODB_URI` variable defines the MongoDB connection string
2. The `STACKY_MONGODB_DATABASE` variable defines the MongoDB database name
3. The `STACKY_AUTH_KEYS` variable defines a set of API tokens with the following structure: `KEY:LEVEL,KEY2:LEVEL` (ex. `abc:2,readOnly:1`)
* There are two levels: `1` (read-only) and `2` (read-and-write)

## Structure
Stacky helps you to choose between different API hosts, depending on their latency.
A `stack` represents a set of `host`s and a `host` is just an URL to an API instance.
We may have a stack named `haste` with the following hosts:

    https://hasteb.in
    https://hastebin.com
    https://paste.helpch.at

We can then make a request to the Stacky API and it will return the host with the lowest latency. Just try around a bit.

## Endpoints
You can use the following endpoints:
1. **GET** `/api/v1/stacks/{name}?token=TOKEN`
2. **PUT** `/api/v1/stacks?token=TOKEN&name=stackName&hosts=optional,default,hosts`
3. **DELETE** `/api/v1/stacks/{name}?token=TOKEN`
4. **GET** `/api/v1/stacks/{name}/hosts/best?token=TOKEN`
5. **GET** `/api/v1/stacks/{name}/hosts?token=TOKEN`
6. **PUT** `/api/v1/stacks/{name}/hosts?token=TOKEN&host=myHost`
7. **DELETE** `/api/v1/stacks/{name}/host?token=TOKEB&host=myHost`

## Responses
Responses will have the following structure:
```json
{
    "code": 200, // The status code
    "type": "success", // Either 'success' or 'error'
    "message": "Done", // A short message
    "data": "something" // Some object (doesn't have to be a string)
}
```

## Wrappers
If you wrote an API wrapper, feel free to share it and I will add it to this Readme :)
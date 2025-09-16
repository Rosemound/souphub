<p align="center">
  <h1 align="center">SoupHub</h1>
  <p align="center">A simple, small hub for sharing game servers over the network</p>
</p>

### About
A simple, small hub for sharing your project's game servers over the network.
Allows you to connect your project's game servers to master hubs and share them with others

This is part of the `soup` system that will be released very soon.

### Configuration
Confifure `souph.json` for yourself (remove all comments):
```json
{
  // your hub name 
  "name": "Rosenound CS2 Project Hub",

  // your hub http port
  "port": "50001",

  // your hub description
  "description": "This is very small description",

  // just env: dev or prod
  "environment": "prod",
  
  // access token to your hub API (don't show it to anyone ;)
  "accessToken": "asdasdasdasdasdasdasdas123asdasd",
  
  // your company / project (optional)
  "company": {
    "name": "Rosemound.ru",
    "url": "https://rosemound.ru"
  },

  // servers that u want to share 
  // (ip:port will be visible only for you and master hub)
  // users will only see the name
  // so you don't risk losing your audience when changing the address
  "servers": {
    "127.0.0.1:3333": {
      // unique name
      "name": "hi",
      "category": "test"
    }
  }
}
```


### Endpoints

#### `POST`

Request is sent from the `master hub` to `your hub`: \
`https://<your_domain>/souph/connect?access_token=<?>`\
```json
// Request raw body
{
  "masters": {
    "<your-master-hub-token>": {

      // Master hub uniq name
      "name": "hub.rosemound.ru",
      
      // Sharing servers
      "servers": [
          "173.12.31.144:53111",
          "173.12.31.145:53331"
      ]
    }
  }
}
```

#### `GET`
Request is sent from the `master hub` to `your hub` (periodic): \
`https://<your_domain>/souph/share?access_token=<?>`\
```json
// Request raw body
{
  "master_token": "<your-master-hub-token>"
}
```
```json
// Response (example)
{
  "name": "Rosenound CS2 Project Hub",
  
  "company": {
    "name": "Rosemound",
    "url": "https://rosemound.ru"
  },

  "servers": {
    "173.12.31.144:53111": {
      "name": "pub",
      "category": "test"
    },

    "173.12.31.145:53331": {
      "name": "surf",
      "category": "test"
    }
  }
}

```

# Go Minesweeper API

A simple Minesweeper API in Go using Gorilla and MongoDB
The latest live deploy can be found [hosted on heroku](https://lknminesweeper.herokuapp.com/)

## Development

It's my first time trying out Go and it was really fun. Got really confused with package names, and folder structure at first. 
I still dont fully get it but im beginning to understand why its designed this way.

The process was fairly straightforward, I wanted to avoid writing too much boilerplate so I started designing the API with swagger first, then used the export feature to generate the routes, models and controllers.
After that I configured Heroku , downloaded the sample Go project and started to integrate the code i had generated to the rest of the project.

I chose MongoDB Atlas for the database since it played nicely with my "code-first" approach. I really wanted to use MySQL or PostgreSQL but since I wanted to save time I went for NoSQL.

## Work In Progress
It still needs to be worked on, most of the code inside endpoint controllers needs to be moved to its own service, and the database client wrapped into a Repository to be reused and avoid reauth to Atlas on each request.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

Add your mongo connect settings to your .env file
```
MONGOURI=mongodb+srv://<user>:<pass>@cluster0.tlxza.mongodb.net/<database>?retryWrites=true&w=majority
MONGODBNAME=minesweeperApi
```

```sh
$ git clone https://github.com/leonardoneumann/minesweeperapi.git
$ cd go-getting-started
$ go build -o bin/minesweeperapi -v . # or `go build -o bin/minesweeperapi.exe -v .` in git bash
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).


## Documentation

A swagger generated html doc is displayed when accessing the app's root. The OpenAPI json file is also attached inside static/swagger.

You can also access a the public yaml file on [https://app.swaggerhub.com/apis/leonardoneumann/MinesweeperGoAPI/1.0](https://app.swaggerhub.com/apis/leonardoneumann/MinesweeperGoAPI/1.0)
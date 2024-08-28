package main

type application struct{}

func main() {
    // set up an app config
    app := application{}

    // get app routes
    mux := app.routes()

    // print out a message

    // start the server

}

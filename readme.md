# Safe Blockchain Demo API

API root: https://safe-demo-api.herokuapp.com/

The API is currently being hosted with Heroku. Why not something like AWS lambda? Well... because Heroku didn't ask for a credit card number (no strings attached API hosting I guess).

## Installation
  1. Install [go](https://golang.org/doc/install).
  2. Install [vs code](https://code.visualstudio.com/download) and add the Go extension.
  3. Install [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

## Setup
  1. Clone this repository
  2. cd into the root folder of this directory.
  3. __In main.go__ make sure the `port := "8000"` is being usedand  __comment out__ the "//for production" `port := os.Getenv("PORT")`.
  4. In the terminal, run `go run .`
  5. Open your browser and go to `localhost:8000`


## Notice
  * All blockchains are currently stored on in a map in this script, meaning 
      * Every change to the API will erase all blockchains... 
      * Any spinning up or tearing down of containers that heroku does also resets the map of blockchains...

  Yah, we need to figure out where to actually put the Blockchain.



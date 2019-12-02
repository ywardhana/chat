# Simple Chat App

## How to install
1. Update dependency by running `make dep`
2. Run the app by running `make run`, and the app will be available at `localhost:8001`
3. To use the RESTFull API, basic auth is needed. You can use `username: chat`, `password: chat` for basic auth.

## API List
Here is APIs that can be used to the app:

### [POST] Create chat
RESTFull API to send a chat. Will be available at `http://localhost:8001/chat`

### [POST] Get all message
RESTFull API to get all already sent chat. Will be available at `http://localhost:8001/chat`

### [GET] Chat Websocket
Websocket API to interactively communicate with the app. You can connect to `ws://localhost:8001/websocket`, and send message to the connection which will be read as string.

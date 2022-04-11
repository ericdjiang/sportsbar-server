# Sportsbar Server

Creator: Eric Jiang (edj9)

This is the Go server for the SportsBar app. The widely-used open source [Gorilla Websocket Library](https://github.com/gorilla/websocket) was used to simplify the low level WebSocket abstractions.

For a detailed overview of the system architecture and features, please see the [PJ2 Project Presentation Slides](https://docs.google.com/presentation/d/1nC_UdXGA_aaM4C_ziQbE7LDvCS6khQljyq9nqucc1j4/edit?usp=sharing).

## Backend Design Overview
Go was chosen because it is highly performant and has a powerful concurrency model. The Go server is hosted via Heroku, and the frontend Swift code is accessible via [Gitlab](https://gitlab.oit.duke.edu/ECE564/spring-2022-projects/sportsbar).

The HTTP server exposes two public API resources, which are the websocket url: wss://sportsbar-server.herokuapp.com/ws and the  [Games By Date RESTful GET endpoint](https://sportsbar-server.herokuapp.com/games?days=2).

> Important Note: The server interfaces with the sportsdata.io API to pull game schedules and data. This app is using the free trial for the API, which comes with several limitations. First, the data received from the API is scrambled and not accurate (you will commonly see high-triple digit scoring games, and the schedules/times are often delayed by several hours). Second, there is a rate limit of 1,000 requests per month. In order to stay under this limit, the server has been programmed to fetch game data in 5 hour intervals. If we were paying for the API, the server would instead be fetching data every 10 seconds for near-realtime game data updates.

To support the realtime game data and live chat functionality, the server implements the Websocket protocol to push updates to clients. Websockets were chosen as opposed to normal polling in order to decrease load on the server and allow live stats/chats to be synchronized across all clients in realtime. A new Goroutine (i.e. thread) is allocated for every client connection. Whenever there is a new chat message or a new game statistic update, the server broadcasts the data to all connected clients.

## Backend Flow
- Handle all Websocket connections
    - Create a new thread for every new client
    - Gracefully manage connections which fail or exit
- Continuously poll for live stats
    - Every 10 seconds, poll sportsdata.io API for live updates for today’s games.
    - Broadcast a sports data message to all WS clients
- Relay chat messages
    - Listen for a chat event from a client connection
    - Broadcast the chat message to all other WS clients
- Serve past 7 days’ game data on a RESTful GET endpoint


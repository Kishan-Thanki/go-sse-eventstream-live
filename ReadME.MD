# SSE (SERVER-SIDE EVENTS) 
JavaScript is used to create a Server-Sent Events (SSE) connection between the client (browser) and the server

## What is EventSource?
EventSource is a built-in JavaScript API used to open a persistent connection to the server, allowing the server to send data to the client over time without the need for repeated requests. This is often used for real-time updates, such as notifications, live data feeds, or chat messages.

## How it works:
#### EventSource: 
- When you create an EventSource object, it establishes an open (another event), long-lived HTTP connection to the server, which can continuously send messages to the client.

- new EventSource("/events"): This creates a new instance of EventSource that connects to the /events endpoint on the server. The URL /events is the server endpoint that the client will listen to for incoming messages.

#### Server-Side Behavior:
- The server will send data to the client in a specific format called Server-Sent Events (SSE), which is a standard for delivering data to the browser over a single HTTP connection.

- SSE allows the server to push messages to the client as long as the connection is open. This is different from regular HTTP requests where the client has to repeatedly request new data.

## Explained:
A server-sent event is when a web page automatically gets updates from a server.

This was also possible before, but the web page would have to ask if any updates were available. With server-sent events, the updates come automatically.

### Client-Side:
- Create a new EventSource object, and specify the URL of the page sending the updates (in this example "/events")
- Each time an update is received, the onmessage event occurs
- When an onmessage event occurs, put the received data into the element with id="resp"

The EventSource Events:
`onerror`: This event handler is called when an error occurs during the connection or while receiving messages.

`close()`: This method closes the connection to the server.

### Server-Side:
- Set the "Content-Type" header to "text/event-stream"
- Specify that the page should not cache
- Output the data to send (Always start with "data: ")
- Flush the output data back to the web page
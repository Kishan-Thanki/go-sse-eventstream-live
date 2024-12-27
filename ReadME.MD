# NOTE: This is just a simple demostration of real-time events. It is experimental

# SSE (SERVER-SIDE EVENTS) 
Server-Sent Events (SSE) is a technology that allows real-time communication from the server to the client via a persistent connection. 

The server will send data to the client in a specific format called Server-Sent Events (SSE), which is a standard for delivering data to the browser over a single HTTP connection.

SSE allows the server to push messages to the client as long as the connection is open. This is different from regular HTTP requests where the client has to repeatedly request new data.

## What is EventSource?
EventSource is a built-in JavaScript API, data is pushed to the client without the need for repeated requests.

## How It Works:
### Client-Side:
- An EventSource object is created with a URL of the page sending the updates (e.g., /events), establishing a continuous connection to the server.

- The client listens for incoming updates with the onmessage event, Each time an update is received, the onmessage event occurs

- When an onmessage event occurs, displays the data in an HTML element.

#### The EventSource Events:
- `onerror()`: This event handler is called when an error occurs during the connection or while receiving messages.

- `close()`: This method closes the connection to the server.

### Server-Side:
- The server sends data in the SSE format (data: <message>) over a persistent HTTP connection.

- The server must set the Content-Type to text/event-stream and Specify that the page should not caching to ensure the data is sent in real-time.

- Flush the output data back to the web page

The server sends continuous updates without the client needing to ask repeatedly. Making it real-time.
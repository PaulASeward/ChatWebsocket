var socket = new WebSocket("ws://localhost:8080/ws"); // Connect to ws endpoint

let connect = cb => {
    console.log("Connecting...");

    socket.onopen = () => {console.log("Succesful Connection");};

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = event => {console.log("Socket Closed Connection ", event);};

    socket.onerror = error => {console.log("Socket Error: ", error);};
};

// Send messsages from frontend to backend using socket
let sendMsg = msg => {
    console.log("Sending msg: ", msg);
    socket.send(msg);
};

export {connect, sendMsg};

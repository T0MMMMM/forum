
const messages = document.getElementById('messages');
const input = document.getElementById('input');
let username = document.getElementById("username").getAttribute("data-variable");
const socket = new WebSocket('ws://10.34.4.52:3000/ws');

socket.onmessage = function(event) {
    messages.innerHTML += '<p>' + event.data + '</p>';
    messages.scrollTop = messages.scrollHeight;
};

function sendMessage() {
    const message = input.value.trim();
    if (message && username != "") {
        socket.send(username + " : " + message);
        input.value = '';
        messages.innerHTML += '<p>' + username + " : " + message + '</p>';
        messages.scrollTop = messages.scrollHeight;
    }
}

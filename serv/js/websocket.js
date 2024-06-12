
const messages = document.getElementById('messages');
const input = document.getElementById('input');
const socket = new WebSocket('ws://localhost:3000/ws');

socket.onmessage = function(event) {
    messages.innerHTML += '<p>' + event.data + '</p>';
    messages.scrollTop = messages.scrollHeight;
};

function sendMessage() {
    const message = input.value.trim();
    if (message) {
        socket.send(message);
        input.value = '';
    }
}

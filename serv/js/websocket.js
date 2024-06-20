const messages = document.getElementById('messages');
const input = document.getElementById('input');
const username = document.getElementById("data").getAttribute("data-variable-username");
const id = document.getElementById("data").getAttribute("data-variable-id");
const topicId = document.getElementById("data").getAttribute("data-variable-topic-id");
const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = function(event) {
    let message = event.data.split(":");
    if (parseInt(message[2]) == topicId) {
        messages.innerHTML += '<p>' + message[0] + " : " + message[1] + '</p>';
        messages.scrollTop = messages.scrollHeight;
    }
};

function sendMessage() {
    const message = input.value.trim();
    if (message && username != "") {
        socket.send(id + ":" + username + ":" + message + ":" + topicId);
        input.value = '';
        messages.innerHTML += '<p>' + username + " : " + message + '</p>';
        messages.scrollTop = messages.scrollHeight;
    }
}

const username = document.getElementById("data").getAttribute("data-variable-username");
const id = document.getElementById("data").getAttribute("data-variable-id");
const socket = new WebSocket('ws://10.34.4.52:8080/ws');

socket.onmessage = function(event) {
    
    let message = event.data.split(":");
    if (message[0] == "[TYPEPrivate]" && id == message[4]) {
        const messages = document.getElementById('messages-' + message[1]);
        messages.innerHTML += '<p>' + message[2] + " : " + filterMsg(message[3]) + '</p>';
        messages.scrollTop = messages.scrollHeight
    } else if (message[0] == "[TYPETopic]" && parseInt(message[3]) == document.getElementById("data").getAttribute("data-variable-topic-id")) {
        const messages = document.getElementById('messages');
        messages.innerHTML += '<p>' + message[1] + " : " + filterMsg(message[2]) + '</p>';
        messages.scrollTop = messages.scrollHeight;
    }
};

function sendMessage() {
    const messages = document.getElementById('messages');
    const input = document.getElementById('input');
    
    const topicId = document.getElementById("data").getAttribute("data-variable-topic-id");

    const message = input.value.trim();
    if (message && username != "") {
        socket.send("[TYPETopic]:" + id + ":" + username + ":" + message + ":" + topicId);
        input.value = '';
        messages.innerHTML += '<p>' + username + " : " + reversefilterMsg(message) + '</p>';
        messages.scrollTop = messages.scrollHeight;
    }
}

function sendPrivateMessage(recipientId) {
    const messages = document.getElementById("messages-" + recipientId);
    const input = document.getElementById("input-" + recipientId);
    const message = input.value.trim();
    if (message && username != "") {
        socket.send("[TYPEPrivate]:" + id + ":" + username + ":" + recipientId + ":" + message );
        input.value = '';
        messages.innerHTML += '<p>' + username + " : " + reversefilterMsg(message) + '</p>';
        messages.scrollTop = messages.scrollHeight;
    }
}

function filterMsg(msg) {
    return msg.replaceAll("'", "[[apostroph]]")
}

function reversefilterMsg(msg) {
    return msg.replaceAll("[[apostroph]]", "'")
}

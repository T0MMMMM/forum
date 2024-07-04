














/* --------------------------------------------------------------------------------------------- */

                                                                                
        const IP = "00.00.00.00"


/* --------------------------------------------------------------------------------------------- */



















const username = document.getElementById("data").getAttribute("data-variable-username");
const id = document.getElementById("data").getAttribute("data-variable-id");
var socket;
if (IP != "00.00.00.00") {
    socket = new WebSocket('ws://' + ip + ':8080/ws');
} else {
    socket = new WebSocket('ws://localhost:8080/ws');
}



socket.onmessage = function(event) {
    
    let message = event.data.split(":");
    if (message[0] == "[TYPEPrivate]" && id == message[4]) {
        const messages = document.getElementById('messages-' + message[1]);
        const privateMessageCardMiddleBox = document.getElementById('privateMessagesScroll-' + message[1]);
        messages.innerHTML += '<p>' + message[2] + " : " + filterMsg(message[3]) + '</p>';
        privateMessageCardMiddleBox.scrollTop = privateMessageCardMiddleBox.scrollHeight;
        console.log(messages.scrollHeight);
    } else if (message[0] == "[TYPETopic]" && parseInt(message[3]) == document.getElementById("data").getAttribute("data-variable-topic-id")) {
        const dateTime = new Date().toLocaleString('sv-SE', { timeZoneName: 'short' }).replace(' ', ' ').substring(0, 19);
        const messages = document.getElementById('messages');
        messages.innerHTML += '<div class="topicAnswerDiv"><div class="answer-info"><div><img src="'+ message[4] +'" alt="" class="userPicture"><p>'+ message[1] +',</p><p class="date">'+ dateTime +'</p></div><div><form action="submit_validate_answer" method="post"><button type="submit" name="button" value="'+ message[5] +'"><i class="fa-regular fa-circle-check"></i></button></form></div></div><p>'+ reversefilterMsg(message[2]) +'</p></div>'
        messages.scrollTop = messages.scrollHeight;
    }
};

function sendMessage() {
    const messages = document.getElementById('messages');
    const input = document.getElementById('input');
    
    const topicId = document.getElementById("data").getAttribute("data-variable-topic-id");
    const pp = document.getElementById("data").getAttribute("data-variable-user-pp");

    const dateTime = new Date().toLocaleString('sv-SE', { timeZoneName: 'short' }).replace(' ', ' ').substring(0, 19);
    
    const message = input.value.trim();
    if (message && username != "") {
        socket.send("[TYPETopic]:" + id + ":" + username + ":" + message + ":" + topicId+ ":" + pp);
        input.value = '';
        messages.innerHTML += '<div class="topicAnswerDiv"><div class="answer-info"><div><img src="'+ pp +'" alt="" class="userPicture"><p>'+ username +',</p><p class="date">'+ dateTime +'</p></div><div></div></div><p>'+ reversefilterMsg(message) +'</p></div>'
        messages.scrollTop = messages.scrollHeight;
    }
}

function sendPrivateMessage(recipientId) {
    const messages = document.getElementById("messages-" + recipientId);
    const input = document.getElementById("input-" + recipientId);
    const message = input.value.trim();
    const privateMessageCardMiddleBox = document.getElementById('privateMessagesScroll-' + recipientId);
    if (message && username != "") {
        socket.send("[TYPEPrivate]:" + id + ":" + username + ":" + recipientId + ":" + message );
        input.value = '';
        messages.innerHTML += '<p>' + username + " : " + reversefilterMsg(message) + '</p>';
        privateMessageCardMiddleBox.scrollTop = privateMessageCardMiddleBox.scrollHeight;
    }
}

function filterMsg(msg) {
    return msg.replaceAll("'", "[[apostroph]]")
}

function reversefilterMsg(msg) {
    return msg.replaceAll("[[apostroph]]", "'")
}

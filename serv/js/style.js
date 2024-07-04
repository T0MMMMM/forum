var hider = document.getElementsByClassName('hider').item(0);
var closeHider = document.getElementsByClassName('closeButtonDiv').item(0);
var messageDivs = document.getElementsByClassName('privateMessage');
var messageCards = document.getElementsByClassName('privateMessageCard');
var arrow = document.getElementsByClassName('arrow');

hider.addEventListener("click", hiderr);
var envelope = document.getElementsByClassName('envelope').item(0);
var messageText = document.getElementsByClassName('messageText').item(0);
var topMessage = document.getElementsByClassName('topMessage').item(0);

var privateMessageWrapper = document.getElementsByClassName('privateMessageWrapper').item(0);

closeHider.addEventListener("click", closeHiderFunc);

function hiderr() {
    if (closeHider.classList.length == 1) {
        hider.classList.add("hiderAfter");
        envelope.classList.add('envelopeAfter');
        messageText.classList.add('messageTextAfter');
        topMessage.classList.add('topMessageAfter');
        privateMessageWrapper.classList.add('privateMessageWrapperAfter');
        setTimeout(() => {
            closeHider.classList.add("closeButtonDivAfter");
            envelope.classList.add('envelopeNotDisplay');
        }, 350);
    }
}

function closeHiderFunc() {
    if (hider.classList.length == 2) {
        console.log(2);
        setTimeout(() => {
            closeHider.classList.remove("closeButtonDivAfter");
        }, 1);
        setTimeout(() => {
            privateMessageWrapper.classList.remove('privateMessageWrapperAfter');
            for (let i = 0; i < messageCards.length; i++) {
                messageCards.item(i).classList.remove("privateMessageCardAfter");
            }
            for (let j = 0; j < messageDivs.length; j++) {
                messageDivs.item(j).classList.remove('privateMessageAfter');
            }
        }, 200);
        setTimeout(() => {
            hider.classList.remove("hiderAfter");
        }, 500);
    }
}


for (let i = 0; i < messageDivs.length; i++) {
    messageDivs.item(i).addEventListener("click", function () {
        console.log(messageCards.item(i));
        privateMessageWrapper.classList.remove('privateMessageWrapperAfter');
        messageCards.item(i).classList.add("privateMessageCardAfter");
        for (let i = 0; i < messageDivs.length; i++) {
            messageDivs.item(i).classList.add("privateMessageAfter");
        };
    });
};

for (let k = 0; k < arrow.length; k++) {
    arrow.item(k).addEventListener("click", function () {
        privateMessageWrapper.classList.add('privateMessageWrapperAfter');
        for (let i = 0; i < messageCards.length; i++) {
            messageCards.item(i).classList.remove('privateMessageCardAfter');
        }
        for (let j = 0; j < messageDivs.length; j++) {
            messageDivs.item(j).classList.remove("privateMessageAfter");
        }
    });
}



// Search Bar 

console.log("search");

var searchIcon = document.getElementsByClassName('searchIcon').item(0);
var searchBar = document.getElementsByClassName('searchBar').item(0);

searchIcon.addEventListener('mouseover', function () {
    searchBar.classList.add("searchBarAfter");
    searchIcon.classList.add("searchIconAfter");
})


// topic add comment button 

var writeAnwerInput = document.getElementsByClassName('writeAnwerInput').item(0);

var topicWriteAnswer = document.getElementsByClassName('topicWriteAnswer').item(0);

var sendComment = document.getElementsByClassName('sendComment').item(0);

if (writeAnwerInput != null) {
    writeAnwerInput.addEventListener("focus", function () {
        topicWriteAnswer.classList.add("topicWriteAnswerAfter");
        sendComment.classList.add("sendCommentAfter");
    })
    writeAnwerInput.addEventListener("focusout", function () {
        setTimeout(() => {
            topicWriteAnswer.classList.remove("topicWriteAnswerAfter");
            sendComment.classList.remove("sendCommentAfter");
        }, 100);
    })
}



// var headerUser = document.getElementsByClassName('headerUser').item(0);

// headerUser.addEventListener("click", function () {
//     window.location.href = "/connexion";
// });

// var navButton = document.getElementsByClassName('navButton').item(0);
// var sideMenu = document.getElementsByClassName('sideMenu').item(0);
// navButton.addEventListener("click", function () {
//     navButton.classList.add("navButtonAfter");
//     sideMenu.classList.add("sideMenuAfter");
// })


var categories = document.getElementsByClassName('categories').item(0);

// categories.addEventListener("click", function () {
//     if (categories.classList.length == 1) {
//         categories.classList.add("categoriesAfter");
//     } else {
//         categories.classList.remove("categoriesAfter");
//     }
// })


var like = document.getElementsByClassName('like');
var heart = document.getElementsByClassName('heart');
var likeAmount = document.getElementsByClassName('likeAmount');

for (let i = 0; i < like.length; i++) {
    like.item(i).addEventListener("click", function () {
        if (!heart.item(i).classList.contains("heartAfter")) {
            heart.item(i).classList.add("heartAfter");
            likeAmount.item(i).innerHTML = parseInt(likeAmount.item(i).innerHTML) + 1;
            socket.send("[TYPELike]:" + id + ":" + like.item(i).getAttribute("data-variable-topic-id"))
        } else {
            heart.item(i).classList.remove("heartAfter");
            likeAmount.item(i).innerHTML = parseInt(likeAmount.item(i).innerHTML) - 1;
            socket.send("[TYPERemoveLike]:" + id + ":" +  like.item(i).getAttribute("data-variable-topic-id"))  
        }
    })
}

var dislike = document.getElementsByClassName('dislike');
var thumbDown = document.getElementsByClassName('thumbDown');
var dislikeAmount = document.getElementsByClassName('dislikeAmount');

for (let i = 0; i < dislike.length; i++) {
    dislike.item(i).addEventListener("click", function () {
        if (!thumbDown.item(i).classList.contains("thumbDownAfter")) {
            thumbDown.item(i).classList.add("thumbDownAfter");
            dislikeAmount.item(i).innerHTML = parseInt(dislikeAmount.item(i).innerHTML) + 1;
            socket.send("[TYPEDislike]:" + id + ":" + dislike.item(i).getAttribute("data-variable-topic-id"))
        } else {
            thumbDown.item(i).classList.remove("thumbDownAfter");
            dislikeAmount.item(i).innerHTML = parseInt(dislikeAmount.item(i).innerHTML) - 1;
            socket.send("[TYPERemoveDislike]:" + id + ":" + dislike.item(i).getAttribute("data-variable-topic-id"))
        }
    })
}



// Profile Info

var viewProfilCard = document.getElementsByClassName('view-profil-card').item(0);
var showProfileInfo = document.getElementsByClassName('showProfileInfo').item(0);

showProfileInfo.addEventListener("click", function () {
    viewProfilCard.focus();
})
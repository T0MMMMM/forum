var hider = document.getElementById('hider');
var closeHider = document.getElementById('closeHider');
var messageDivs = document.getElementsByClassName('privateMessage');
var messageCards = document.getElementsByClassName('privateMessageCard');
var arrow = document.getElementsByClassName('arrow').item(0);

hider.addEventListener("click", hiderr);

function hiderr() {
    hider.id = "hiderAfter"
    setTimeout(() => {
        closeHider.id = "closeHiderAfter";
    }, 350);
}

closeHider.addEventListener("click", closeHiderFunc)

function closeHiderFunc() {
    closeHider.id = "closeHider";
    setTimeout(() => {
        for (let i = 0; i < messageCards.length; i++) {
            messageCards.item(i).classList.remove("privateMessageCardAfter");
        }
        for (let j = 0; j < messageDivs.length; j++) {
            messageDivs.item(j).classList.remove('privateMessageAfter');
        }
    }, 200);
    setTimeout(() => {
        hider.id = "hider";
    }, 500);
}


for (let i = 0; i < messageDivs.length; i++) {
    messageDivs.item(i).addEventListener("click", function () {
        console.log(messageCards.item(i));
        messageCards.item(i).classList.add("privateMessageCardAfter");
        for (let i = 0; i < messageDivs.length; i++) {
            messageDivs.item(i).classList.add("privateMessageAfter");
        };
    });
};

arrow.addEventListener("click", function () {
    for (let i = 0; i < messageCards.length; i++) {
        messageCards.item(i).classList.remove('privateMessageCardAfter');
    }
    for (let j = 0; j < messageDivs.length; j++) {
        messageDivs.item(j).classList.remove("privateMessageAfter");
    }
});


var headerUser = document.getElementsByClassName('headerUser').item(0);

headerUser.addEventListener("click", function () {
    window.location.href = "/connexion";
});

var navButton = document.getElementsByClassName('navButton').item(0);
var sideMenu = document.getElementsByClassName('sideMenu').item(0);
navButton.addEventListener("click", function () {
    navButton.classList.add("navButtonAfter");
    sideMenu.classList.add("sideMenuAfter");
})


var categories = document.getElementsByClassName('categories').item(0);

categories.addEventListener("click", function () {
    if (categories.classList.length == 1) {
        categories.classList.add("categoriesAfter");
    } else {
        categories.classList.remove("categoriesAfter");
    }
})


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
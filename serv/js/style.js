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
        hider.id = "hider";
    }, 200);
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
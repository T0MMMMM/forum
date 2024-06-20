/* Connexion */ 

var connexionHiders = document.getElementsByClassName('connexionHider');
var connexionHider = connexionHiders.item(0);
var connexionButtons = document.getElementsByClassName('connexionButton');
var connexionButton = connexionButtons.item(0);
var textButton = document.getElementsByClassName('buttonText').item(0);
var hiderText1 = document.getElementById('hiderText1');
var hiderText2 = document.getElementById('hiderText2');


connexionButton.addEventListener("click", function () {
    if (connexionHider.classList.length == 1) {
        connexionHider.classList.add("connexionHiderAfter");
        textButton.innerHTML = "Log in";
        hiderText1.id = "hiderText1After";
        hiderText2.id = "hiderText2After";
    } else {
        connexionHider.classList.remove("connexionHiderAfter");
        textButton.innerHTML = "Sign up";
        hiderText1.id = "hiderText1";
        hiderText2.id = "hiderText2";
    }
});
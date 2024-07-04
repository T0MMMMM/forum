# Y-NOT

## Étapes pour lancer le programme

1. **Pré-requis :**
   - Assurez-vous d'avoir Go (Golang) version 1.22 installé sur votre machine.
   -  <span style="color:red; font-size: 18px;"> IMPORTANT </span> / Il faut modifier l'adresse ip sous la forme "00.00.00.00" par votre propre adresse ip (que vous pouvez trouver avec la commande `ipconfig` sur Windows et `ip a` sous linux) dans le fichier 'serv/js/websocket.js' pour que les messages soient envoyés en temps réels pour tous les utilisateurs connectés \

2. **Installation et lancement de l'application :**

   **Option 1 : Lancer avec Go**

   - Clonez le dépôt du projet : `git clone https://github.com/T0MMMMM/forum.git
   - Accédez au répertoire du projet : `cd forum`
   - Téléchargez les dépendances nécessaires : 
         `go get github.com/gofiber/fiber/v2`
         `go get github.com/gofiber/websocket/v2`
         `go get github.com/gofiber/template/html/v2`
         `go get modernc.org/sqlite`
   - Exécutez le programme : `go run main.go`
   - Ouvrez votre navigateur et allez à l'adresse : `http://VOTRE_IP:8080`

   **Option 2 : Lancer avec Docker**

   - Clonez le dépôt du projet : `git clone https://github.com/T0MMMMM/forum.git
   - Accédez au répertoire du projet : `cd forum`
   - Construisez l'image Docker : `docker build . -t forum`
   - Exécutez le conteneur Docker : `docker run -p 8080:8080 forum`
   - Ouvrez votre navigateur et allez à l'adresse : `http://VOTRE_IP:8080`

## Toutes les fonctionnalités du forum :

### Système de connexion / enregistrement :

1. **Enregistrement :**
   - Si le nom d'utilisateur n'est pas déjà utilisé, on peut créer un compte.
   - L'utilisateur est redirigé vers l'accueil en étant connecté.

2. **Connexion :**
   - Le nom d'utilisateur et le mot de passe doivent correspondre.

3. **Déconnexion :**
   - Se déconnecte et renvoie sur la page de connexion.

### Personnalisation du profil :

1. En haut à droite, cliquer sur la photo de profil ouvre une petite interface qui affiche les informations de l'utilisateur (pseudo, email, date de création du compte) et lui propose de modifier son profil :
    - Changer de pseudo si le pseudo n'est pas déjà utilisé.
    - Changer de photo de profil parmi différents choix.
    - Les changements sont enregistrés dans la base de données.

### Messages privés :

Chaque utilisateur a accès à une interface en bas à droite de la page d'accueil qui permet d'avoir une liste de tous les utilisateurs enregistrés sur le forum. On peut écrire un message à n'importe qui, le destinataire peut voir les messages qu'il a reçus. Pour faciliter la recherche d'une personne, vous pouvez directement la rechercher grace à une barre de recherche en haut de la liste des utilisateur.

### Voir les profils des autres utilisateurs :

Dans l'interface des messages privés, on peut cliquer sur un bouton pour voir le profil d'un utilisateur et certaines de ses informations (pseudo, email, photo de profil, réponses, topics).

### Topic :

1. **Créer un nouveau topic :**
   - L'utilisateur doit cliquer sur le bouton "create new topic" sur l'interface à gauche.
   - Il est redirigé vers une nouvelle page pour créer un topic. Il peut choisir la catégorie parmi celles proposées, un titre et une description.
   - Le topic est ajouté parmi les autres topics.

2. **Voir les topics :**
   - Sur la page d'accueil, l'utilisateur peut voir tous les topics de tous les utilisateurs.
   - Il peut filtrer les topics par catégorie ou bien en faisant une recherche.

3. **Interactions avec les topics :**
   - Un utilisateur peut liker et/ou disliker un topic.

4. **Accéder au topic via le bouton commentaire :**
   - Sur un topic, un utilisateur peut répondre à la question posé par le topic et participer à la discussion.
   - Le créateur du topic peut valider une seule et unique réponse, et le topic apparaît à l'accueil comme "résolu".

### Compatibilité mobile (Responsive Design) :

Le site est conçu pour être responsif et fonctionnel sur les appareils mobiles. Les interfaces et les fonctionnalités s'adaptent aux différentes tailles d'écran.

arno tom luka
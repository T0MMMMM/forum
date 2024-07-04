# Y-NOT

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

Chaque utilisateur a accès à une interface en bas à gauche de la page d'accueil qui permet d'avoir une liste de tous les utilisateurs enregistrés sur le forum. On peut écrire un message à n'importe qui, le destinataire peut voir les messages qu'il a reçus.

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
   - Un utilisateur peut liker ou disliker un topic.

4. **Accéder au topic via le bouton commentaire :**
   - Sur un topic, un utilisateur peut répondre au topic ou bien aux autres utilisateurs.
   - Le créateur du topic peut valider une seule et unique réponse, et le topic apparaît à l'accueil comme "résolu".
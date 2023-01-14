# Répartition du Travail
## Répartition du travail prévue (après l'arrivé de Matéo)

### Partie Maquette :  

Mathis - Création de la page QuiSommesNous.html et MonCompte.html  

Valentin - Création des autres pages  

### Partie Frontend :  

Mathis - Création des pages Achat.html, MonCompte.html et Erreur.html    
Mathis - Mise en place du style des pages HTML suivant les maquettes  
Mathis - Mise en place de la majorité du responsive sur les 3 pages  
Mathis - Recherche de bug à la fin du projet
Mathis - Création du readme      
 

Valentin - Création de la page Index.html et Location.html  
Valentin - Mise en place du style des pages HTML suivant les maquettes  
Valentin - Finition globale du responsive sur toute les pages  
Valentin - Création de tout les fichiers JavaScript  
Valentin - Correction de bug à la fin du projet  
Valentin - Création du Header    


Matéo - Création de la page Inscription.html et QuiSommesNous.html  
Matéo - Mise en place du style des pages HTML suivant les maquettes  
Matéo - Mise en place d'un version adaptative du site pour mobile sur les 2 pages  
Matéo - Création du Footer  


### Partie Backend :

Matéo - Création de la structure de donnée des voitures et des utilisateurs  
Matéo - Mise en place du serveur web en go  
Matéo - Transformation des pages html en templates 

 
Valentin - Gestions des bases de données (encodage / décodage JSON)  
Valentin - Réorganisation et optimisation du fichier main.go  

# Organisation des dossiers

Le PROJET WEB est trié dans plusieurs dossiers : 

- Dans le dossier "maquette" nous retrouvons toutes les maquettes du site. C'est en quelque sorte les prémisses du site qui nous ont permis de nous mettre d'accord sur le design du site.
  
- Dans le dossier "templates" nous retrouvons toute les pages HTML du site:  
  
    - [Index.html](templates/index.html) qui est la page d'accueil du site et qui permet de naviguer vers les autres pages  
    - [Achat.html](templates/Achat.html) qui est la page qui permet de faire une recherche de véhicule pour un achat  
    - [Location.html](templates/location.html) qui est la page qui permet de faire une recherche de véhicule pour une location  
    - [Resultats.html](templates/Resultats.html) qui est un meilleurs aperçu d'un produit à acheter ou louer  
    - [Inscription.html](templates/Inscription.html) qui est la page qui permet de s'inscrire sur le site  
    - [MonCompte.html](templates/MonCompte.html) qui est la page qui permet de se connecter à son compte  
    - [QuiSommesNous.html](templates/quisommesnous.html) qui est la page qui permet de voir les informations sur l'entreprise  
    - [Erreur.html](templates/Erreur.html) qui est une page qui s'affiche si une erreur est détectée.      

- Dans le dossier [css](css) nous retrouvons tout les fichiers CSS du site qui permettent de mettre en forme les pages HTML.        
Ainsi nous avons décidé de faire un fichier CSS par page HTML pour les grosses pages comme [Index.html](templates/index.html) ou encore [Resultats.html](templates/Resultats.html) et un fichier CSS commun pour les petites pages comme [Achat.html](templates/Achat.html) et [Location.html](templates/location.html).  
D'autant plus que le header et le footer sont les mêmes pour toutes les pages.  

- Dans le dossier [images](images) nous retrouvons toutes les images utilisées sur le site classé dans différents sous-dossier. Notamment l'image du logo de l'entreprise, des images de véhicules, des images de fonds, des images de boutons, des images de logos de réseaux sociaux, etc.  
  
- Dans le dossier [font](font) nous retrouvons toutes les polices utilisées sur le site. Notamment la police "Roboto-Regular.ttf" qui est la police principale du site.  

- Le dossier [json](json) permet de stocker les données comme par exemples les données de connection avec l'adresse mail et le mot de passe.

- Le dossier [script](script) contient les fichiers JavaScript du site. Comme par exemple le fichier [script.js](script/script.js) qui permet de faire fonctionner le slider de la page [Index.html](templates/index.html).  


  
   
  
    







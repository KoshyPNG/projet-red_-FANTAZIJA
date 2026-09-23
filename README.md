# Projet RED - Evil Resident

Jeu de rôle en ligne de commande écrit en Go, inspiré des jeux de survie et
d'horreur. Le joueur combat des monstres, récupère des ressources et améliore son équipement.

## Prérequis

- Go 1.27.1 ou version compatible
- Un terminal capable de lire les entrées clavier

## Installation et lancement

Installer le ficher le projet en ZIP dans le dossier User puis dézipper le .

Depuis la racine du projet :

```bash
PS C:\Users\NumUser\projet-red_Evil_Resident> go run .\main\
```
Le jeu se lançera !!!!

## Déroulement d'une partie

Au début de la partie, le chef des S.T.A.R.S. vous envoie explorer une zone
dangereuse. Choisissez ensuite votre classe. Le menu principal propose :

| Touche | Action |
| --- | --- |
| `1` | Avancer vers une nouvelle zone |
| `2` | Afficher les informations du personnage |
| `3` | Ouvrir l'inventaire |
| `4` | S'entraîner contre M.A.X.I.M.E. |
| `5` | Quitter la partie |

En avançant, un événement aléatoire peut se produire : combat, campement,
marchand ou découverte de pièces.

## Classes

Chaque classe commence avec trois actions différentes :

- **CRS S.T.A.R.S.** : vie élevée, avec `Coup de matraque`, `Cri de guerre`
	et `Bloquer`.
- **Unité Tactique S.T.A.R.S.** : profil dégat critique, avec `Glock 26`,
	`Fléchette de poison` et `Dodge`.
- **Unité d'extermination S.T.A.R.S.** : grande réserve d'essence et gros dégat sans critique, avec `Lance flamme`, `Rechargement` et `Coup de crosse`.

## Combats

Pendant un combat, choisissez le numéro d'une action. 

Les monstres peuvent être vaincus en réduisant leur vie à zéro. Les ennemis
sont le Zombie, le Claqueur, Solar et le Goliath. L'entraînement permet
d'affronter M.A.X.I.M.E.

Pendant un combat :

- `1` à `3` : utiliser une attaque du personnage ;
- `4` : ouvrir l'inventaire et utiliser un objet ;
- `5` : quitter le combat contre M.A.X.I.M.E.

## Campement

Dans un campement, le joueur peut :

- se reposer pour récupérer de la vie et de l'essence ;
- fabriquer des objets à partir des ressources de l'inventaire ;
- quitter le campement.

Recettes disponibles :

| Objet fabriqué | Ressources nécessaires |
| --- | --- |
| Potion de soin | Herbe |
| Baril d'essence | Peau fermentée |
| Potion de poison | Champignon |
| Casquette Gucci Fraise | Tissus + Corde |
| Gilet pare-balles | Plaque en fer + Kevlar |
| Timberland | Tissus + Caoutchouc |

## Marchand

Le marchand vend des potions, de l'essence et des sacoches qui augmentent la
capacité de l'inventaire. Il propose aussi une ressource de craft aléatoire
pour 10 pièces. Les ressources possibles sont `Herbe`, `Champignon`,
`Peau fermenté`, `Tissus`, `Corde`, `Plaque en fer`, `Kevlar` et `Caoutchouc`.

Solar peut donner le `PP7 silencieux`. Cet objet s'utilise pendant un combat
et vainc immédiatement le monstre ciblé.

## Structure du projet

- `main/` : point d'entrée du programme.
- `character.go` : rcéation et statistiques du personnage.
- `combat.go` : logique des combats et des tours.
- `monstre.go` : définition et initialisation des monstres.
- `camp.go` : repos et fabrication d'objets.
- `marchand.go` : boutique et achats.
- `accessinventory.go` et `display.go` : inventaire et objets utilisables.
- le reste sont des fichier utilisé par les fichier principaux. 

## État du projet

Projet d'étude réaliser en 1 semaine avec 2.0000000001 personnes . L'interface est le terminal.
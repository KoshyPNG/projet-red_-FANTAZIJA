# Projet RED - Evil Resident

Jeu de rôle en ligne de commande écrit en Go, inspiré des jeux de survie et
d'horreur. Explorez des zones dangereuses, combattez des monstres, récupérez
des ressources et améliorez votre équipement.

## Prérequis

- Go 1.27.1 ou une version compatible ;
- un terminal capable de lire les entrées clavier.

## Installation et lancement

Clonez le dépôt ou téléchargez-le, puis exécutez le programme depuis la racine
du projet :

```bash
PS C:\Users\NumUser\projet-red_Evil_Resident> go run .\main\
```
Le jeu se lançera !!!!

## Déroulement d'une partie

Le chef des S.T.A.R.S. vous envoie explorer une zone dangereuse. Le menu
principal propose les actions suivantes :

| Touche | Action |
| --- | --- |
| `1` | Avancer vers une nouvelle zone |
| `2` | Afficher les informations du personnage |
| `3` | Ouvrir l'inventaire |
| `4` | S'entraîner contre M.A.X.I.M.E. |
| `5` | Quitter la partie |

Chaque déplacement déclenche un événement aléatoire : combat, campement,
marchand ou découverte de pièces.

## Classes

Chaque classe possède ses propres statistiques et trois actions de départ :

| Classe | Profil | Actions |
| --- | --- | --- |
| **CRS S.T.A.R.S.** | 100 PV, 50 d'essence | `Coup de matraque`, `Cri de guerre`, `Bloquer` |
| **Unité Tactique S.T.A.R.S.** | 80 PV, 65 d'essence | `Glock 26`, `Fléchette de poison`, `Dodge` |
| **Unité d'extermination S.T.A.R.S.** | 60 PV, 100 d'essence | `Lance flamme`, `Rechargement`, `Coup de crosse` |

## Combats

Les ennemis rencontrés sont le Zombie, le Claqueur, Solar et le Goliath.
L'entraînement permet d'affronter M.A.X.I.M.E. Réduisez les PV du monstre à
zéro pour gagner le combat.

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
| Baril d'essence | Peau fermenté |
| Potion de poison | Champignon |
| Casquette Gucci Fraise | Tissus + Corde |
| Gilet par balles | Plaque en fer + Kevlar |
| Timberland | Tissus + Caoutchouc |

## Marchand

Le marchand vend des objets de soin, de l'essence, des poisons et des
sacoches. Une sacoche augmente la capacité de l'inventaire de 10 places ; elle
peut être achetée au maximum trois fois.

| Objet ou service | Prix |
| --- | ---: |
| Potion de soin | 15 pièces |
| Grande potion de soin | 30 pièces |
| Baril d'essence | 25 pièces |
| Potion de poison | 20 pièces |
| Sacoche | 50 pièces |
| Ressource de craft aléatoire | 10 pièces |

Les ressources disponibles sont `Herbe`, `Champignon`, `Peau fermenté`,
`Tissus`, `Corde`, `Plaque en fer`, `Kevlar` et `Caoutchouc`.

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
# 🧠 Hackaton Golang - Piscine Zone01
Ce dépôt contient les exercices réalisés lors du **Hackaton Go** organisé pendant la 3ᵉ semaine de la [piscine Golang](https://github.com/hayatmz/piscine-golang) sur le campus Zone01.

## 📅 Contexte
- 🕘 **Début** : 9h du matin
- 🔁 **Rythme** : un nouvel exercice toutes les **2 heures**
- 🌙 **Durée** : 24h non-stop, jusqu'au lendemain 9h
- 🧑‍💻 **Lieu** : en présentiel au campus Saint-Marc, Rouen
- 🌌 **Ambiance** : nuit blanche, code, boissons énergisantes... et beaucoup d'entraide ! 💪

## Objectifs
- Mettre en pratique l'ensemble des notions acquises durant les premières semaines de la piscine golang.
- Résoudre des problèmes variés sous contraites de temps (et fatigue).
- Apprendre à travailler en autonomie et en entraide dans un contexte intensif.

## Contenu
[rot14.go](./rot14.go) : Chiffrement [ROT14](https://fr.wikipedia.org/wiki/Chiffrement_par_d%C3%A9calage) d'une chaîne de caractères.<br>
[abort.go](./abort.go) : Retourne la [médiane](https://fr.wikipedia.org/wiki/M%C3%A9diane_(statistiques)) de 5 entiers.<br>
[collatzcountdown.go](./collatzcountdown.go) : Calcule le nombre d'étapes de la [conjecture de Collatz](https://fr.wikipedia.org/wiki/Conjecture_de_Syracuse).<br>
[comcheck](./comcheck/main.go) : Détecte la présence de mots sentibles dans les arguments.<br>
[enigma.go](./enigma.go) : Manipulation avancée de pointeurs multiples.<br>
[pilot](./pilot/main.go) : Déclaration et affichage d'une structure ```Pilot```.<br>
[fixthemain](./fixthemain/main.go) : Corriger un programme de gestion d'état d'une porte.<br>
[compact.go](./compact.go) : Nettoie une clice de chaînes en supprimant les valeurs 'falsy'.<br>
[activebits.go](./activebits.go) : Compte le nombre de bits à 1 dans la représentation binaire d'un entier.<br>
[max.go](./max.go) : Retounrne la plus grande valeur d'une slice d'entiers.<br>
[unmatch.go](./unmatch.go) : Trouve la veleur sans paire dans une slice.<br>
[join.go](./join.go) : Concatène une slice de chaînes avec un séparateur.<br>

## Ce que j'ai appris
- L'importance de l'organisation sous contrainte de temps.
- Comment réfléchir rapidement à une solution efficace en Go.
- Travailler efficacement avec les autres sous pression.
- Manipuler en profondeur les slices, structures, pointesurs, fonctions, etc.

## Installation
Pour explorer ou tester les projets localement :

1. Assure toi dd'avoir Go installé sur ta machine. Tu peux vérifier avec :<br>
```
go version
```

Si besoin, [installe Golang](https://go.dev/doc/install)

2. Cloner le dépôt :<br>

```
git clone https://github.com/hayatmz/hackaton-golang
```
```
cd hackaton-golang
```

Si ton dossier n'a pas encore de module Go, commence par :<br>
```
go mod init piscine
```

Pour installer la bibliothèque pédagogique de Zone01 (utilisée pour **fixthemain/main.go**) :<br>
```
go get github.com/01-edu/z01
```

3. Pour lancer un fichier unique (hors dossier spécifique) :<br>
```
go run nom_du_fichier.go
```

**Et** pour lancer les fichiers situés dans un dossier :<br>
```
cd nom_du_dossier
```
```
go run main.go
```

4. Pour compiler puis exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```

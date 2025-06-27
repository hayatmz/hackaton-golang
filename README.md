# 🧠 Hackaton Golang - Piscine Zone01
Ce dépôt contient les exercices réalisés lors du **Hackaton Go** organisé pendant la 3ᵉ semaine de la [piscine Golang](./https://github.com/hayatmz/piscine-golang) sur le campus Zone01.

## 📅 Contexte
- 🕘 **Début** : 9h du matin
- 🔁 **Rythme** : un nouvel exercice toutes les **2 heures**
- 🌙 **Durée** : 24h non-stop, jusqu'au lendemain 9h
- 🧑‍💻 **Lieu** : en présentiel au campus Saint-Marc, Rouen
- 🌌 **A**mbiance** : nuit blanche, code, boissons énergisantes... et beaucoup d'entraide ! 💪

## Objectifs
- Mettre en pratique l'ensemble des notions acquises durant les premières semaines de la piscine golang.
- Résoudre des problèmes variés sous contraites de temps (et fatigue).
- Apprendre à travailler en autonomie et en entraide dans un contexte intensif.

## Contenu
[rot14.go](./rot14.go) : .<br>
[abort.go](./abort.go) : .<br>
[collatzcountdown.go](./collatzcountdown.go) : .<br>
[comcheck](./comcheck/main.go) : .<br>
[enigma.go](./enigma.go) : .<br>
[pilot](./pilot/main.go) : .<br>
[fixthemain](./fixthemain/main.go) : .<br>
[compact.go](./compact.go) : .<br>
[activebits.go](./activebits.go) : .<br>
[max.go](./max.go) : .<br>
[unmatch.go](./unmatch.go) : .<br>
[join.go](./join.go) : .<br>

## Ce que j'ai appris
- L'importance de l'organisation sous contrainte de temps.
- Comment réfléchir rapidement à une solution efficace en Go.
- Travailler efficacement avec les autres sous pression.
- Manipuler en profondeur les slices, structures, pointesurs, fonctions, etc.

## Installation
Pour explorer ou tester les projets localement :

1. Assure toi dd'avoir Go installé sur ta machine. Tu peux vérifier avec :<br>
```go version```<br>

Si besoin, [installe Golang](./golang.org/dl)

2. Cloner le dépôt :<br>

```git clone https://github.com/hayatmz/hackaton-golang```<br>
```cd hackaton-golang```<br>

Si ton dossier n'a pas encore de module Go, commence par :<br>
```go mod init piscine```<br>

Pour installer la bibliothèque pédagogique de Zone01 (utilisée pour **fixthemain/main.go**) :<br>
```go get github.com/01-edu/z01```<br>

3. Pour lancer un fichier unique (hors dossier spécifique) :<br>
```go run nom_du_fichier.go```<br>

**Et** pour lancer les fichiers situés dans un dossier :<br>
```cd nom_du_dossier```<br>
```go run main.go```<br>

4. Pour compiler puis exécuter :<br>
```go build nom_du_fichier.go```<br>
```./nom_du_fichier```
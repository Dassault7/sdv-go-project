# get-joke

`get-joke` est une application CLI écrite en Go qui permet de récupérer des blagues depuis [JokeAPI](https://v2.jokeapi.dev/). L'utilisateur peut personnaliser les blagues récupérées en fonction de plusieurs paramètres comme la langue, la catégorie, le nombre de blagues, et bien plus.

Le résultat est affiché dans la console.

## Installation

Pour installer `get-joke`, il suffit de cloner le dépôt et de lancer la commande suivante :

```bash
go build -o build/get-joke
```

Pour utiliser `get-joke` depuis n'importe où, il faut ajouter le chemin du dossier `build` dans la variable d'environnement `PATH`.


## Fonctionnalités

- Récupération de blagues depuis JokeAPI
- Personnalisation des blagues récupérées
- Affichage des blagues récupérées

### Commandes

Pour avoir une blague aléatoire, on utilise la commande suivante :

```bash
get-joke
```

Voici les options disponibles (aucune option n'est obligatoire) :

- `--amount | -a` : Nombre de blagues à récupérer (par défaut : `1`)
- `--blacklist | -b` : Liste des catégories à exclure (par défaut : `none`). Possible : `nsfw`, `religious`, `political`, `racist`, `sexist`, `explicit`
- `--category | -c` : Catégorie de la blague (par défaut : `any`). Possible : `any`, `misc`, `programming`, `pun`, `spooky`, `christmas`
- `--help | -h` : Affiche l'aide
- `--lang | -l` : Langue de la blague (par défaut : `en`). Possible : `en`, `fr`, `de`, `es`, `cs`, `pt`
- `--output-url | -o` : Affiche l'URL de la requête
- `--type | -t` : Type de blague. (par défaut `none`). Possible : `single`, `twopart`

Exemple :

```bash
get-joke -o -a 2 -l de -c programming -b racist
```

On peux voir les options disponibles avec la commande suivante :

```bash
get-joke --help
```

Les valeurs pour les options sont sensibles à la casse. On peut les obtenir en utilisant la commande:

```bash
get-joke list <command>
```

Exemple :

```bash
get-joke list language  # Alias: lang, l
get-joke list categories  # Alias: category, cat, c
get-joke list blacklist  # Alias: black, bl, b
get joke list types  # Alias: type, t
```

## Explications

Projet réalisé dans le cadre de l'évaluation du module de programmation en Go.
J'ai utilisé toutes les bases vues en cours pour réaliser ce projet.

- Gestion des erreurs
- Tests unitaires
- Requêtes HTTP
- Goroutines

J'ai pris la liberté d'ajouter une couche CLI pour rendre l'application plus interactive, 
en utilisant la librairie [Cobra](https://cobra.dev/).  
Cette librairie permet de créer des applications CLI de manière simple et rapide, avec la gestion des flags et des arguments.



### Sources

- [JokeAPI](https://v2.jokeapi.dev/)
- [Chanel](https://www.youtube.com/watch?v=nNXhePi3xwE)
- [Cobra](https://umarcor.github.io/cobra/)

### Remarques

Utilisation d'IA pour comprendre plus facilement des concepts de Go. (Goroutines, Channels, etc.)  
Mais _ce projet n'a pas de code source venant d'IA_.

## Auteurs

Développé par [Jarod Guichard](https://github.com/Dassault7).

# Projet d'API de ToDoList

Ce projet est une API REST de gestion de ToDoList multi-utilisateur, construite avec **Go** et les modules suivants :

- **Fiber** : framework web rapide et minimaliste
- **Gorm** : ORM pour interagir avec PostgreSQL
- **Swagger** : documentation interactive de l'API
- **Logrus** : gestion avancée des logs
- **PostgreSQL** : base de données relationnelle

## Fonctionnalités principales
- **Gestion des utilisateurs** :
    - Création d'un utilisateur avec un mot de passe.
    - Connexion pour récupérer un token JWT.
- **Gestion des ToDoLists** :
    - Ajout, modification et suppression de tâches.
    - Un utilisateur peut uniquement accéder à ses propres tâches.

---

## Prérequis
- Docker et Docker Compose

### Variables d'environnement
Créez un fichier `.env` à la racine du projet avec les variables suivantes :

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=todolist_db
JWT_SECRET=your_jwt_secret
```

---

## Installation
1. Clonez le dépôt :
   ```bash
   git ...
   cd ...
   ```

2. Lancez l'application avec Docker Compose :
   ```bash
   docker-compose up --build
   ```

3. Accédez à la documentation Swagger :
   [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)

---

## Routes de l'API

### Authentification

#### 1. Créer un utilisateur
**POST** `/api/v1/auth/register`

Body JSON :
```json
{
  "username": "example_user",
  "password": "example_password"
}
```

Réponse :
```json
{
  "message": "User registered successfully"
}
```

#### 2. Se connecter
**POST** `/api/v1/auth/login`

Body JSON :
```json
{
  "username": "example_user",
  "password": "example_password"
}
```

Réponse :
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### Gestion des ToDo
**Toutes les routes suivantes nécessitent un header d'authentification :**
```http
Authorization: Bearer <token>
```

#### 1. Récupérer toutes les tâches
**GET** `/api/v1/todos`

Réponse :
```json
[
  {
    "id": 1,
    "title": "Acheter du pain",
    "description": "Aller à la boulangerie",
    "completed": false
  }
]
```

#### 2. Créer une tâche
**POST** `/api/v1/todos`

Body JSON :
```json
{
  "title": "Acheter du lait",
  "description": "Passer au supermarché",
  "completed": false
}
```

Réponse :
```json
{
  "message": "Task created successfully",
  "task": {
    "id": 2,
    "title": "Acheter du lait",
    "description": "Passer au supermarché",
    "completed": false
  }
}
```

#### 3. Modifier une tâche
**PUT** `/api/v1/todos/:id`

Body JSON :
```json
{
  "title": "Acheter des oeufs",
  "description": "Pour faire une omelette",
  "completed": true
}
```

Réponse :
```json
{
  "message": "Task updated successfully"
}
```

#### 4. Supprimer une tâche
**DELETE** `/api/v1/todos/:id`

Réponse :
```json
{
  "message": "Task deleted successfully"
}
```

---

## Authentification JWT
L'API utilise des tokens JWT pour authentifier les utilisateurs. Lorsqu'un utilisateur se connecte, il reçoit un token qu'il doit inclure dans le header `Authorization` pour accéder aux routes protégées.

### Exemple d'usage
Header HTTP :
```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## Journalisation (Logging)
**Logrus** est utilisé pour gérer les logs de l'application. Les différents niveaux de logs (info, warning, error) permettent de suivre les activités et débugger efficacement.

# Dev Web Conso API

[![CI](https://github.com/M-PROGRAMMATION/dev-web-conso-api/actions/workflows/ci.yml/badge.svg)](https://github.com/M-PROGRAMMATION/dev-web-conso-api/actions/workflows/ci.yml)

API REST en Go développée dans le cadre du cours **Développement Web / Consommation d'API** (YNOV).

## 🛠️ Stack technique

- **Langage :** [Go](https://go.dev/) 1.27
- **Librairies :** uniquement la bibliothèque standard (`net/http`, `encoding/json`)

## 📁 Structure du projet

```
dev-web-conso-apii/
├── go.mod              # Déclaration du module et version de Go
├── README.md
└── src/
    ├── main.go          # Point d'entrée, déclaration des routes HTTP
    ├── banner/
    │   └── banner.go    # Affichage d'une bannière stylisée au démarrage du serveur
    ├── routes/
    │   ├── default.go   # Route racine "/"
    │   ├── health.go    # Route "/health"
    │   ├── artists.go   # Routes "/artists"
    │   └── relation.go  # Routes "/relation"
    ├── objects/
    │   ├── artists_object.go
    │   └── relation_object.go
    └── utils/
        └── json.go      # Helpers d'écriture JSON
```

## 🚀 Lancer le projet

```bash
go run ./src
```

Le serveur démarre par défaut sur `http://localhost:8080`.

## 📡 Routes disponibles

| Méthode | Route             | Description                                      |
|---------|-------------------|---------------------------------------------------|
| GET     | `/`               | Healthcheck simple, retourne `OK`                  |
| GET     | `/health`         | Statut du serveur, heure de démarrage et uptime    |
| GET     | `/artists`        | Retourne les artistes au format JSON               |
| GET     | `/relation`       | Retourne toutes les relations (dates/lieux)        |
| GET     | `/relation/{id}`  | Retourne une relation par son id                   |

### Exemple de réponse — `GET /artists`

```json
[
  {
    "id": 1,
    "name": "Aurore Nomade",
    "creationDate": 2017,
    "members": ["Lina", "Nora", "Sami"],
    "firstAlbum": "12-05-2019"
  }
]
```

### Exemple de réponse — `GET /relation/1`

```json
{
  "id": 1,
  "datesLocations": {
    "toulouse-france": ["14-03-2025", "15-03-2025"],
    "lyon-france": ["02-04-2025"],
    "paris-france": ["20-05-2025"]
  }
}
```

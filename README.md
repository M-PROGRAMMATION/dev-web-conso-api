# Dev Web Conso API

[![CI](https://github.com/M-PROGRAMMATION/dev-web-conso-api/actions/workflows/ci.yml/badge.svg)](https://github.com/M-PROGRAMMATION/dev-web-conso-api/actions/workflows/ci.yml)

API REST en Go développée dans le cadre du cours **Développement Web / Consommation d'API** (YNOV).

## 🛠️ Stack technique

- **Langage :** [Go](https://go.dev/) 1.27
- **Librairies :** la bibliothèque standard (`net/http`, `encoding/json`) pour l'API, et [`swaggo/swag`](https://github.com/swaggo/swag) + [`swaggo/http-swagger`](https://github.com/swaggo/http-swagger) pour la doc Swagger

## 📁 Structure du projet

```
dev-web-conso-apii/
├── go.mod              # Déclaration du module et version de Go
├── README.md
├── data/
│   ├── artists.json     # Persistance JSON des artistes
│   └── relations.json   # Persistance JSON des relations
└── src/
    ├── main.go          # Point d'entrée, déclaration des routes HTTP
    ├── orm/
    │   └── orm.go        # Mini ORM générique (store JSON : All/Find/Create/Update/Delete/Where)
    ├── routes/
    │   ├── init.go       # Initialisation des stores + enregistrement des routes
    │   ├── default.go    # Route racine "/"
    │   ├── health.go     # Route "/health"
    │   ├── artists.go    # Routes "/artists"
    │   └── relation.go   # Routes "/relation"
    ├── objects/
    │   ├── artists_object.go
    │   └── relation_object.go
    └── utils/
        ├── json.go       # Helpers d'écriture JSON
        ├── banner.go     # Affichage d'une bannière stylisée au démarrage du serveur
        └── store.go      # Déclaration et seed des stores (ArtistStore, RelationStore)
```

## 🚀 Lancer le projet

```bash
go run ./src
```

Le serveur démarre par défaut sur `http://localhost:8080`.

## 📖 Documentation Swagger

La doc interactive est générée automatiquement depuis des commentaires (`swag`) placés au-dessus de chaque handler, puis servie via `swaggo/http-swagger`.

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g src/main.go -o docs
```

À relancer après chaque ajout/modif de route (les fichiers `docs/` sont générés, à ne pas éditer à la main). Une fois le serveur lancé :

👉 http://localhost:8080/swagger/index.html

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

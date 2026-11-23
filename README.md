# Dev Web Conso API

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
    └── banner/
        └── banner.go    # Affichage d'une bannière stylisée au démarrage du serveur
```

## 🚀 Lancer le projet

```bash
go run ./src
```

Le serveur démarre par défaut sur `http://localhost:8080`.

## 📡 Routes disponibles

| Méthode | Route       | Description                                      |
|---------|-------------|---------------------------------------------------|
| GET     | `/`         | Healthcheck, retourne `OK`                         |
| GET     | `/artists`  | Retourne un artiste au format JSON                 |

### Exemple de réponse — `GET /artists`

```json
{
  "id": 1,
  "name": "Aurore Nomade",
  "creationDate": 2017,
  "members": ["Lina", "Nora", "Sami"],
  "firstAlbum": "12-05-2019"
}
```

# Anime Repository

This repository hosts two applications: a REST API backend written in Golang and a web application frontend written in React. The backend provides endpoints for fetching animes and episodes, as well as scraping episode sources from a third-party website to embed as iframes in the web application. The frontend consists of two pages: one for listing animes and another for listing and playing anime episodes.

## Table of Contents

- [Backend](#backend)
  - [Libraries Used](#libraries-used)
  - [Endpoints](#endpoints)
- [Frontend](#frontend)
  - [Technologies Used](#technologies-used)
  - [Pages](#pages)
- [Setup](#setup)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Contributing](#contributing)
- [License](#license)

## Backend

### Libraries Used

- **Migrations**: [go-migrations](https://github.com/ShkrutDenis/go-migrations)
- **JWT**: [golang-jwt](https://github.com/golang-jwt/jwt/v5)
- **Swagger**: [echo-swagger](https://github.com/swaggo/echo-swagger)
- **Mocking DB**: [go-mocket](https://github.com/selvatico/go-mocket)
- **ORM**: [gorm](https://gorm.io/gorm)

### Endpoints

- **Fetch Animes**
  - **GET** `/api/animes`
  - Fetch a list of all available animes.

- **Fetch Episodes**
  - **GET** `/api/animes/:id`
  - Fetch a list of episodes for a specific anime.

- **Scrape Episode Source**
  - **GET** `/episodes/:id/src`
  - Scrape the source of an episode from a third-party website and embed it as an iframe.

## Frontend

### Technologies Used

- **TailwindCSS**: for styling
- **TypeScript**: for type-safe JavaScript
- **Vite**: for fast development and build tooling

### Pages

- **Anime List**
  - Displays a list of all available animes.

- **Episode List and Player**
  - Displays a list of episodes for a selected anime.
  - Allows playing an episode by embedding it as an iframe.


## Setup

### Backend setup

1. Install dependencies
```sh
make install-api
```

2. Run Migrations
```sh
make migrate
```

3. Run
```sh
make dev-api
```
 or run both
```sh
make -j2
```

### Frontend setup

1. Install dependencies
```sh
make install-client
```

2. Run
```sh
make dev-client
```

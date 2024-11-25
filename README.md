# kita-go-auth

Yet another authentication application built in GO! This project is built _purely_ for learning GO, it is **not** intended for production use case.
You _can_ use this for self-hosted applications that are **not** exposed to the wider internet.

Thank you in advance for any feedback, issues reports and contributions!

## Features

- Register
- Login
- Logout
- Refresh token
- Activate account
- Deactivate account
- Delete account

## Libraries

**TODO:** Link to library

- Gin
- Gorm
- JWT
- Bcrypt
- Rate
- GoDotEnv

## Development

Run the following command to build the Docker image:

```
docker build -t kita-go-auth .
```

Run the container with the built image:

```
docker run -p 3001:3001 kita-go-auth
```

Run using `docker-compose` or `podman-compose` (Recommended)

```
docker-compose up --build
```

```
podman-compose up --build
```

```
docker stop $(docker ps -aq) && docker rm $(docker ps -aq); podman rm kita-go-auth; podman build -t kita-go-auth .; podman-compose up -d
```

## Goals

Including the list below, the goal is to create a microservice focused on authentication and user management as a domain.

- [x] Basic CRUD operations for user
- [ ] Basic role system: Admin, Basic, Guest
- [ ] Admin CRUD operations to control users (User CRUD)
- [ ] Something else?

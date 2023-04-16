# :books: Todo Rest Api GoLang

This project it's to understand the mainly tools used for language GoLang to API services, ORM, Migration and another stuffs to fixed the knowledge adquired.

## :wrench: Setup and Tools used

- GoLang 1.20
- Chi router
- GoOrm
- MigrationDB
- GoDotEnv
- Docker

This list does not complete yet, will increase with the time.

## :dart: Goals

- Understading the tools most commom used with GoLang

- Produce a front to consume the API (In discussion if Angular or VueJs)

- Try adopt the best pratices for GoLang Projects

- Understand Api Rest with GoLang

## :rocket: Using the project

**With Docker**
- Open a terminal in the directory root this project and run:
```console
user@machine:~$ docker build -t todo-go-rest
```
- When docker process finished, execute:
```console
user@machine:~$ docker run --rm -it  todo-go-rest:latest
```
- You will be able to see **a item TODO in Json**

**With Docker compose**
- Open a terminal in the directory root this project and run:
```console
user@machine:~$ docker compose up
```

The system will start the application and you can to access the: **localhost:3000/todo**

This documentation it's not complete, I'm still working on this resource to automatization better
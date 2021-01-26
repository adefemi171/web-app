# Web App



# App Details
.

# Proposed Stack

Golang, JS (NextJS)


# Top-level directory layout

    📦web-app
        📦docs
            ┣ 📜
        📦frontend
            ┣ 📦components
                ┗ 📜link.js
            ┣ 📦pages
                ┣ 📦api
                    ┗ 📜hello.js
                ┗ 📜_app.js
                ┗ 📜index.js
                ┗ 📜table.js
            ┣ 📦public
                ┗ 📜favicon.ico
                ┗ 📜vercel.svg
            ┣ 📦styles
                ┗ 📜globals.css
                ┗ 📜Home.module.css
                ┗ 📜Table.module.css
            ┣ 📜.env.example
            ┣ 📜Dockerfile
            ┣ 📜package.json
            ┣ 📜package-lock.json

            📦src
            ┣ 📦pkg
                ┣ 📦controller
                    ┗ 📜controller.go
                ┣ 📦route
                    ┗ 📜route.go
            ┗ 📜.env.example
            ┗ 📜Dockerfile
            ┗ 📜go.mod
            ┗ 📜go.sum
            ┗ 📜main.go
        ┣ 📜.gitignore
        ┣ 📜docker-compose.yml
        ┣ 📜go.mod
        ┣ 📜go.sum
        ┣ 📜README.md

## The task can be found [here](https://github.com/adefemi171/web-app/blob/main/Docs/Calipsa%20-%20Junior%20Devops%20Technical%20Assessment.1611157693.pdf)

# How to setup project and run locally

### Clone the repository 

```
git clone https://github.com/adefemi171/web-app.git
```
### Checking Out
The App is built on the ``` master ``` branch you will need to checkout to the app branch using:

```
git checkout master
```

### Open Two different terminal/console

#### CLI 1

1. Change Directory to api
2. Create a .env file with the values below:
```
APP_PORT=:8080
PROMETHEUS_PORT=:9110
```
3. Run the go src
```
go run main.go
```
4. In your browser, navigate to

```
http://localhost:8080/
```
for the web application and navigate to

```
http://localhost:9110/metrics
```
for the Prometheus metrics

#### CLI 2

1. Change Directory to frontend
2. Install all dependencies using
```
npm install
```
3. Run the web application
```
npm run dev
```
4. In your browser, navigate to

```
http://localhost:3000/
```

# How to setup project using docker
The [backend](https://hub.docker.com/repository/docker/adefemi171/web-api) and [frontend](https://hub.docker.com/repository/docker/adefemi171/web-app) image is published on dockerhub

1. Run the docker-compose.yml file using
```
docker-compose up
```
And visit the endpoint in your browser for both front end and backend respectively as specified above

# How to setup project using helm
1. I used [Kind]() to create my cluster
2. In the root folder run:
```
kind create cluster --config kind.yaml
```
3. Change Directory to helm folder and run
```
helm install <application_name> .
```
4. Then watch the pod create with
```
kubectl get pod -w
```
5. Once the pod has been created, you need to ge the service name and port forward to be able to access the container, run
```
kubectl get service
kubectl port-forward svc/<service_name> 8080:8080 3000:3000 9110:9110
```

And visit the endpoint in your browser for both front end and backend respectively as specified above


NOTE: [Helm](https://helm.sh/docs/intro/install/), [Docker](https://docs.docker.com/get-docker/), [Golang](https://golang.org/doc/install), [Node](https://nodejs.org/en/download/), and [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/) needs to be installed to run this application





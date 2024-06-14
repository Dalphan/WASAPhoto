# WASAPhoto
Web And Software Architecture exam project

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white">
  <img src="https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white">
  <img src="https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white">
  <img src="https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=F7DF1E">
  <img src="https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D">
  <img src="https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white">
  <img src="https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white">
  <img src="https://img.shields.io/badge/GIT-E44C30?style=for-the-badge&logo=git&logoColor=white">
</div>

## Description
_Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can
upload your photos directly from your PC, and they will be visible to everyone following you._

## Design Specification
Each user will be presented with a stream of photos (images) in reverse chronological order, with
information about when each photo was uploaded (date and time) and how many likes and comments
it has. The stream is composed by photos from “following” (other users that the user follows). 

Users can place (and later remove) a “like” to photos from other users. Also, users can add comments to any
image (even those uploaded by themself). Only authors can remove their comments.

Users can ban other users. If user Alice bans user Eve, Eve won’t be able to see any information about
Alice. Alice can decide to remove the ban at any moment.
Users will have their profiles. The personal profile page for the user shows: the user’s photos (in reverse
chronological order), how many photos have been uploaded, and the user’s followers and following.
Users can change their usernames, upload photos, remove photos, and follow/unfollow other users.
Removal of an image will also remove likes and comments.

A user can search other user profiles via username.

A user can log in just by specifying the username.

## Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an the API server
    * `service/database` contains database related operations
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
    * `service/utils` contains reusable utility functions for various tasks in the project
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is the web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding
	
## OpenAPI
### How to view the OpenAPI documentation 
Visit the [Swagger editor](https://editor.swagger.io/) website and paste the API documentation inside `doc/api.yaml`.

## Backend
### How to build 
If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:
```
go build ./cmd/webapi/
```
If you're using the WebUI and you want to embed it into the final executable:
```
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## Frontend
### How to run (in development mode)
You can launch the backend only using:
```
go run ./cmd/webapi/
```
If you want to launch the WebUI, open a new tab and launch:
```
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

## Deployment
### How to build the images 
Backend
```
$ docker build -f Dockerfile.backend -t wasaphoto-backend:latest .
```
Frontend 
```
$ docker build -f Dockerfile.frontend -t wasaphoto-frontend:latest .
```
### How to run the container images
Backend
```
$ docker run -it --rm -p 3000:3000 wasaphoto-backend:latest
```
Frontend
```
$ docker run -it --rm -p 5137:80 wasaphoto-frontend:latest
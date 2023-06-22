## docker tasks

As a Devops you are asked to make use of Docker in production.
This is a working example with the Dockerfile, but it has some issues.

### Questions
1. If any of the files in the folder changes, yarn reinstalls all dependencies from scratch.
   Would you be able to modify the Dockerfile to leverage Docker layer caching?
2. The application runs as a root. Is it possible to force it to run under an unprivileged user?
3. A developer wants to use the same Dockerfile for development on their local machine, but it
   requires rebuilding the image every time. Is there a way to improve it?

### Answers
1. Completed. Answer in `Dockerfile`
2. Completed. Answer in `Dockerfile`
3. Semi-completed. `docker-compose.yml` updated so that changes being made in the local directory are being reflected in the running container. Will first need to install the dependencies of the application, and any subsequent time dependencies are added, the image will have to be rebuilt. e.g. `docker build -t node-app .`. Run `docker compose up` to run the app using docker-compose. And if you exec into the container, e.g. `docker exec -it <container-id> sh`, you will see the `index.js` being updated inside the container as you save your work locally. However, this does not get reflected in the browser. My guess is that there was a bug introduced that affected the hot reload of the browser when changes are applied. See GitHub issue here: https://github.com/facebook/create-react-app/issues/9984 
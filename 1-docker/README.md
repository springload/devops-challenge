## docker tasks

NOTE: the answers in here are in addition to the comments I have made on the actual dockerfile

As a Devops you are asked to make use of Docker in production.
This is a working example with the Dockerfile, but it has some issues.

NOTE: the answers in here are in addition to the comments I have made on the actual dockerfile

1. If any of the files in the folder changes, yarn reinstalls all dependencies from scratch.
   Would you be able to modify the Dockerfile to leverage Docker layer caching?
A:Package.js and Yarn.lock both should be copied first and cached as they are required by yarn to install dependancies, index.js should be copied after YARN install as it will be changed often and copying it before would mean YARN would re install everytime

2. The application runs as a root. Is it possible to force it to run under an unprivileged user?
A: Simply create a new usergroup and user and then switch to the user at the end of the file(so that everything can be built as root). Alternatively you can run the docker daemon in rootless mode which would allow you to build the container as a non root user (https://docs.docker.com/engine/security/rootless/).

3. A developer wants to use the same Dockerfile for development on their local machine, but it
   requires rebuilding the image every time. Is there a way to improve it?
A: I would mount the the local directory as a volume in the docker container by passing the -v flag when running the container. 
Such as--> docker run -v </hostdir:/containterdir> 
Alternatively this can also be done via docker-compose (i have changed the compose.yml file)
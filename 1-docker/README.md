## docker tasks

As a Devops you are asked to make use of Docker in production.
This is a working example with the Dockerfile, but it has some issues.

1. If any of the files in the folder changes, yarn reinstalls all dependencies from scratch.
   Would you be able to modify the Dockerfile to leverage Docker layer caching?
   
   Elina: Yes, it can be done by moving the 'yarn install' layer before the 'COPY' files layer. So that docker cache keeps what has been done before the changes in the folder. It allows that rerunning the changed files and all subsequent steps only. Please see dockerfile for details.
   
2. The application runs as a root. Is it possible to force it to run under an unprivileged user?

	Elina: Yes, you can do this by using docker privilege mode. However, there're security concerns when usuing privilege capabilities. An attacher who is inside the container can take the advantage of these capabilities to attack the host. Please see docker-compose.yml for using docker privilege mode. 
	
3. A developer wants to use the same Dockerfile for development on their local machine, but it
   requires rebuilding the image every time. Is there a way to improve it?
   
   Elina: The development environment can use the code when starting a container by mounting local directory instead of creating a new image on each time. Please see docker-compose.yml for mounting source code directory. When you execute the docker-compose command, the volumes directive in docker-compose.yml file mounts source directories or volumes from your computer at target paths inside the container. If a matching target path exists already as part of the container image, it will be overwritten by the mounted path.

## docker tasks

As a Devops you are asked to make use of Docker in production.
This is a working example with the Dockerfile, but it has some issues.

1. If any of the files in the folder changes, yarn reinstalls all dependencies from scratch.
   Would you be able to modify the Dockerfile to leverage Docker layer caching?
2. The application runs as a root. Is it possible to force it to run under an unprivileged user?
3. A developer wants to use the same Dockerfile for development on their local machine, but it
   requires rebuilding the image every time. Is there a way to improve it?

### Answers

1. COPY adds files from your Docker client’s current directory, this command should be run toward the end of the Dockerfile, in order to take advantage of Docker Layer Caching. I have moved this command to run after yarn install, so that if any of the files in the folder do change, `RUN yarn install` layer cache will not be invalidated.

2. It is possible to force the application to run as an unprivileged user. This can be forced by adding the `--user` flag when you start the docker container, i.e. `docker run --user nobody image:tag`. Docker then executes the program with the user specified. This particular user, an existing user on the host, can be located within the User Database (`/etc/passwd`).

3. The developer can share their project directory with a started container, using bind mounts, that way they only need to build the image once, not every time code is changed. If docker compose is used, the change can be made in the docker-compose.yml file to mount a local directory i.e.

```
  volumes:
      - ./source_dir:/app/target_dir
```

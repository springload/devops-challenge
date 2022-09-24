## docker tasks

As a Devops you are asked to make use of Docker in production.
This is a working example with the Dockerfile, but it has some issues.

1. If any of the files in the folder changes, yarn reinstalls all dependencies from scratch.
   Would you be able to modify the Dockerfile to leverage Docker layer caching?
   
Solution: Copying only the package.json and yarn.lock files before running `yarn install` allows this to be cached until these files change rather than any file.

2. The application runs as a root. Is it possible to force it to run under an unprivileged user?

Solution: Add a command in the Dockerfile to add a non-root user and set this user

3. A developer wants to use the same Dockerfile for development on their local machine, but it
   requires rebuilding the image every time. Is there a way to improve it?

Solution: Mount the local development directory as a volume in the docker container so that local changes are immediately reflected without the need to rebuild the image
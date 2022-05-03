## docker tasks

As a Devops you are asked to make use of Docker in production.
This is a working example with the Dockerfile, but it has some issues.

1. If any of the files in the folder changes, yarn reinstalls all dependencies from scratch.
   Would you be able to modify the Dockerfile to leverage Docker layer caching?
2. The application runs as a root. Is it possible to force it to run under an unprivileged user?
3. A developer wants to use the same Dockerfile for development on their local machine, but it
   requires rebuilding the image every time. Is there a way to improve it?

## Candidate notes:
- Hope the changes in this commit satisfy at least some if not all of the requirements.
    - The Dockerfile have been split into two separate files:
        - **dependencies.Dockerfile**:
            - Create a docker image containing only the yarn prerequisites for running our yarn app.
            - This forms the base image for our application container.
        - **app.Dockerfile**:
            - Create a docker image containing our yarn app built up on top of the dependencies docker base image.
    - All yarn commands are now running under unprivileged user **node**.
    - The dependencies base image (which is tagged as `companyrepo/app_yarn_dependencies:latest`) can be made available in the company docker registry ahead of time.
        - In doing so, developers would pull the dependencies docker image instead of building it locally in their workstations (using `docker-compose -f docker-compose.yml up --build application`).  And they can then concentrate in making development code changes in the **app** directory only.


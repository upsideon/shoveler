#!/bin/sh

# Logging into DockerHub account.
docker login -u upsideon

# Building shoveler-backend and pushing to DockerHub.
cd backend
docker build -t ${SHOVELER_DOCKERHUB_REPO}:shoveler-backend .
docker push ${SHOVELER_DOCKERHUB_REPO}:shoveler-backend
cd -

# Building shoveler-frontend and pushing to DockerHub.
cd frontend
docker build -t ${SHOVELER_DOCKERHUB_REPO}:shoveler-frontend .
docker push ${SHOVELER_DOCKERHUB_REPO}:shoveler-frontend
cd -


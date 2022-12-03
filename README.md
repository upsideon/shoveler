## Shoveler

When it snows, there are people who have difficulty or are unable to shovel the walkway at their place of residence. Some have family members, friends, or contractors to help them with this task. For those who don't have these resources, Shoveler seeks to be a solution.

The goals of this project, which is still in its infancy, are as follows:

- If you are in need of snow shoveling assistance, there is a dedicated venue in which to make that need known and fulfilled.
- If you have a desire to help others in your community, providing snow shoveling assistance is a clear and present option.

There is no business model here. If this becomes something more than code on a Git repository, the intention is to run the service as a non-profit via donations.

### Development

All services are containerized using Docker. If you haven't installed Docker already, please follow the instructions [here](https://docs.docker.com/engine/install/) before proceeding.

If it is your first time running Shoveler, execute the following commands from the top-level directory to bring it up:

```
docker-compose up -d cassandra
docker exec -i shoveler_cassandra_1 cqlsh < database/migrations.cql
docker-compose up -d shoveler-frontend shoveler-backend
```

On subsequent runs, as long as you don't remove the Cassandra container, you can execute the following commands to bring up Shoveler:

```
docker-compose up -d cassandra && sleep 60 && docker-compose up --build
```

The sleep command gives the Cassandra database enough time to come up before the Shoveler backend attempts to access it. In the future, either intelligent retries by the backend or health checks to support `depends_on` in `docker-compose.yml` will be added to address this problem.

The web page can be accessed by going to your browser and navigating to [http://localhost:3000](http://localhost:3000). To access the backend service directly, make HTTP requests to [http://localhost:8080](http://localhost:8080). If these ports are already in use on your system, feel free to update the port mapping in `docker-compose.yml` to meet your needs.

When you're ready to stop Shoveler, execute the following command from the top-level directory:

```
docker-compose stop
```

### Deploying Container Images

To deploy container images to DockerHub for the frontend and backend services, first add the following environment variable export to the configuration script for your favorite shell. An example for Bash is provided below:

```
echo "export SHOVELER_DOCKERHUB_REPO=<dockerhub-repo>" >> ~/.bashrc
```

Then execute `dockerhub-push.sh` build the images and push them up to DockerHub.

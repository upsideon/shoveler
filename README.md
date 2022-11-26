## Shoveler

When it snows, there are people who have difficulty or are unable to shovel the walkway at their place of residence. Some have family members, friends, or contractors to help them with this task. For those who don't have these resources, Shoveler seeks to be a solution.

The goals of this project, which is still in its infancy, are as follows:

- If you are in need of snow shoveling assistance, there is a dedicated venue in which to make that need known and fulfilled.
- If you have a desire to help others in your community, providing snow shoveling assistance is a clear and present option.

There is no business model here. If this becomes something more than code on a Git repository, the intention is to run the service as a non-profit via donations.

### Development

All services are containerized using Docker. If you haven't installed Docker already, please follow the instructions [here](https://docs.docker.com/engine/install/) before proceeding.

To bring up Shoveler, simply execute the following command from the top-level directory:

```
docker-compose up -d
```

The web page can be accessed by going to your browser and navigating to [http://localhost:3001](http://localhost:3001). To access the backend service directly, make HTTP requests to [http://localhost:8081](http://localhost:8081). If these ports are already in use on your system, feel free to update the port mapping in `docker-compose.yml` to meet your needs.

When you're ready to bring Shoveler down, execute the following command from the top-level directory:

```
docker-compose down
```

### Deployment

At the time of writing, Shoveler has not been deployed to the public Internet.
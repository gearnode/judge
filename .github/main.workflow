workflow "feature-branch" {
  on = "push"
  resolves = [
    "Run unit test",
    "Build docker image",
    "Login to Docker Hub",
    "Tag image for Docker Hub",
    "Push image to Docker Hub",
  ]
}

action "Run unit test" {
  uses = "./.github/actions/go-test"
}

action "Build docker image" {
  needs = [
    "Run unit test",
    "Login to Docker Hub",
  ]
  uses = "actions/docker/cli@master"
  args = "build . -t judge"
}

action "Login to Docker Hub" {
  uses = "actions/docker/login@76ff57a"
  secrets = ["DOCKER_PASSWORD"]
  env = {
    DOCKER_USERNAME = "gearnode"
  }
}

action "Tag image for Docker Hub" {
  needs = ["Build docker image"]
  uses = "actions/docker/tag@master"
  args = "judge gearnode/judge --no-latest --no-sha"
}

action "Push image to Docker Hub" {
  needs = ["Login to Docker Hub", "Tag image for Docker Hub"]
  uses = "actions/docker/cli@master"
  env = {
    CONTAINER_REGISTRY_PATH = "gearnode"
    IMAGE_NAME = "judge"
  }
  args = ["push", "$CONTAINER_REGISTRY_PATH/$IMAGE_NAME"]
}

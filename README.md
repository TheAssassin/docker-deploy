# Docker-Deploy

This little application can be used to keep a git repository inside a Docker
in sync with a remote repository.

The application spawns a little HTTP server exposing an API endpoint that can
be used as a URL for web hooks like GitHub, Gitea, GitLab etc. offer them.

Although it is a bit against the Docker philosophy (where a container should
rather be replaced instead of updated), this kind of deploy system is easier
to set up than having a companion tool/container which has to re-create the
other containers.

More info to follow soon!

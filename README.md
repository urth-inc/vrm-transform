# vrm-transform

A Go library for transforming GLB and VRM files.

## Overview

`vrm-transform` is a Go library designed to facilitate the transformation and manipulation of GLB/VRM files.

## Features

- Resize textures in GLB or VRM files.
- Convert png/jpeg textures to KTX2 format.

## Contributing

We welcome contributions! Please follow these steps to contribute to the project:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes.
4. Push to the branch.
5. Open a pull request.
6. Ensure that your code adheres to the coding standards and includes appropriate tests.

## Development

You can use Docker to set up a consistent development environment. Follow the steps below to get started.

1. Clone this repository:

```
git clone git@github.com:urth-inc/vrm-transform.git
```

Add `test.glb` to `/assets` directory. You can use any GLB/VRM file for testing.

2. Build `vrm-transform` image using Docker:

```
docker build -t vrm-transform .
```

Verify that the `vrm-transform` image was created:

```
docker image ls | grep vrm-transform
# Expected output
vrm-transform                              latest               34e8a2a5f0a5   17 seconds ago   2.16GB
```

3. Run the Docker container and open a Bash shell:

```
docker run --name vrm-transform -it vrm-transform:latest
root@5927e498153f:/app#
```

If you encounter following error, you need to remove the existing container with the same name.

```
docker: Error response from daemon: Conflict. The container name "/vrm-transform" is already in use by container "5927e498153fe6d26d5f2f12415a37909ec1a8a27084a104a6acd61ddae721ed". You have to remove (or rename) that container to be able to reuse that name.
See 'docker run --help'.

# Stop and Remove the existing container
docker stop vrm-transform && docker rm vrm-transform

# Run the Docker container and open a Bash shell again
docker run --name vrm-transform -it vrm-transform:latest
```

You can now run vrm-transform commands inside the container.

```
# main.go is the entry point of the test application.
./main
```

4. Check the output file:

If you want to see the output file in a 3D viewer, move the output file from the container to the host and then open it in the viewer. Here is an example command to move output.glb from the container to the host:

```
docker container cp vrm-transform:/app/output.glb .
```

We recommend using the `glTF Report` to check your Model.

- glTF Report: https://gltf.report/

## License

This project is licensed under the MIT License.

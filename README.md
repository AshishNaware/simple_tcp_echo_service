<h1 align="center">Welcome to simple_tcp_echo_service 👋</h1>
<p>
  <img alt="GitHub Release Version" src="https://img.shields.io/github/v/tag/AshishNaware/simple_tcp_echo_service?label=GitHub%20Version" />
  <img alt="Docker Image Version" src="https://img.shields.io/docker/v/ashishnaware/simple_tcp_echo_service?label=Docker%20Version&sort=semver" />
  <a href="https://choosealicense.com/licenses/mit/" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> A minimalist, Go-based TCP echo server for learning and testing networking, socket inspection, and containerization workflows.

## Install

```sh
go run main.go
```

## Usage

1. Run the tcp server locally or using the docker image
```sh
go run src/main.go
```
OR using docker
```sh
docker run -p 50000:50000 ashishnaware/simple_tcp_echo_service:latest
```

2. On new terminal, run
```sh
nc localhost 50000
```
And then write any message and press enter.

You should be able to see the echoed message back from the server.

## Author

👤 **Ashish Naware**

* Website: https://www.ashishnaware.me
* Github: [@AshishNaware](https://github.com/AshishNaware)
* LinkedIn: [@ashishnaware](https://linkedin.com/in/ashishnaware)

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2025 [Ashish Naware](https://github.com/https:\/\/github.com\/AshishNaware).<br />
This project is [MIT](https://choosealicense.com/licenses/mit/) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
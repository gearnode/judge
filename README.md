# Judge

> Simple distributed access management

Judge is an "access management service" inspired by the AWS IAM Service.

## Table of content

- [Abstract](#abstract)
- [Concepts](#concepts)
- [Roadmap](#roadmap)
- [Contributing](#contributing)

## Abstract

Judge attempts dealing with _access management_ in a distributed system. Judge is inspired by the _AWS IAM Service_.

Judge implementation is focus on these values:
- [Simple API](#simple-api)
- [Distributable](#distributable)
- [Auditable](#auditable)
- [Extendable](#extendable)

### Simple API

Jude provide a _human readable_ syntax to execute an _access request_. An _access request_ can be expressed like this:

> **Who** is **able** to do **what** on **something** given some **context**?

- **Who**: The subject of the action, for example your End-User.
- **Able**: The effect which can be either "allow" or "deny".
- **What**: An action the subject attemps to perform.
- **Something**: An object the subject attemps to target.
- **Context**: The context containing information about the environment such as the IP Address, request date, the resource owner name or any other information you want to pass along.

### Distributable

Jude expose an API through [Google gRPC](https://grpc.io/).

Google gRPC is built on top of **HTTP/2**. **HTTP/2** come with many performance improvements:
- binary instead of textual with **HTTP/1**
- fully multiplexed, instead of ordered and blocking with **HTTP/1**
- header compression to reduce overhead

HTTP/2 tends to be the new default standard in the web's ecosystem.

Google gRPC implements official support of 11 programming languages (less maintenance for SDKs). Easy extendable (authentication, load balancing, etc.) and distribuable (low latency, etc.). By using Google gRPC Judge API is by design built in for polyglot distributed system.

TODO: explain storage solution
  - many database alrady solve the distributable issue.
    - judge interface optimised storage type (now only k/v database).
  - inject the storage you want and be free about query optimization.
  - maybe provide a dedicated database!

### Auditable

- auditable by design
- open catalog
- judge use judge to protect this catalog (example by design)
  - may explain dog fooding

### Extendable


## Concepts

### Object Resource Name (ORN)

Object Resource Name (ORN) uniquely identify resources in your system. The
Object Resource Name contains many information used by Judge to evaluate the
authorization.

### Policy
TODO

### End-User
TODO

### Organization
TODO

### Authorization Request
TODO


## Usage

- mkcert 127.0.0.1

:warning: this project can only be used in your gopath directory. Generate protos are painful and I can't lose time again.

TODO explain make


Start judge server:
```
bin/judgeserver -tls-crt priv/server.crt -tls-key priv/server.key -tls
```

## Roadmap

### v1alpha2

- apiserver
    - Reflection gRPC API support
    - TLS support
    - Simple CLI interface
    - Policy Management (list,show,create,update,delete)

- judgectl
    - Define CLI interface
    - Policy support

- subject management
    - add / update / remove subject
    - link policy to subject

- groups management
    - add / update remove group
    - link subject with group
    - link policy to group


### v1alpha1

- [x] gRPC basic setup
- [ ] Policy management via gRPC (list,create,delete,update,show)
- [x] basic storage interface
- [x] gRPC server configurable
- [x] alpha authorize rules define
- [x] alpha orn format define

## Contributing

1. Fork it (<https://github.com/gearnode/judge/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

<!-- Markdown link & img dfn's -->
[travis-image]: https://travis-ci.com/gearnode/judge.svg?branch=master
[travis-url]: https://travis-ci.org/gearnode/judge

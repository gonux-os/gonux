# God

*Golang daemon*


## About

God is a daemon responsible for providing IPC to a linux system. Its goal is to allow communication between various applications over multiple protocols, providing a translation layer between them.

*e.g. A process should be able to send a message using the GraphQL protocol to another process that receives this data in the gRPC protocol.*

### Protocols:

* [ 💥 ] gRPC *(over Protobuf over TCP or UNIX socket)*
* [ 💭 ] GraphQL *(over HTTP)*
* [ 💭 ] RESTful API *(over HTTP)*
* [ 💭 ] DBus *(over UNIX socket)*

**Legend:**
```
💭 - Still in planning phase
💥 - Entirely experimental
❗ - Complete but unstable
✅ - Complete and stable
⛔ - Abandoned / Deprecated
```


## Philosophy

God doesn't aim to be a DBus substitute in philosophy, but being DBus compatible, in practice perhaps it may be used as such?

# GoNUX Init

*It's an init, innit?*


## About

GoNUX Init is an init system meant to be run by the kernel as the first process (PID 1) after boot.

Its main goal is to bootstrap the system, manage processes, services, users, shells, and essencially make sure everything is ready and in the right place for the rest of the system to run.

*In a perfect world, GoNUX Init would be a drop-in replacement to systemd, providing not only its own opinionated interface, but also a compatibility layer with systemd's interface, file formats, and commands.*


## Architecture

GoNUX Init is in its early stages of development, so its architecture is subject to sudden and drastic changes, however the main goal is to be configurable (also ideally pluggable).

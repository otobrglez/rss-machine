# rss-machine

[rss-machine][rss-machine] is web service for parsing RSS channels written in Go.

[![Build Status][circle-ci-badge]][circle-ci]
[![ImageLayers][imagelayers-badge]][imagelayers]
[![Docker Pulls][docker-pulls-badge]]()
[![Docker Stars][docker-stars-badge]]()
	

Supported types feed types `RSS 1.0`, `RSS 2.0` and `ATOM 1.0`.

Service uses [SlyMarbo/rss](https://github.com/SlyMarbo/rss) for parsing RSS feeds.

## Running from Docker

```
docker run -ti -p 5000:5000 otobrglez/rss-machine

curl <docker_host>:5000/parse?url=http://rss.cnn.com/rss/edition.rss
```

## Author

- [Oto Brglez][me]

[circle-ci]: https://travis-ci.org/otobrglez/rss-machine
[circle-ci-badge]: https://travis-ci.org/otobrglez/rss-machine.svg?branch=master
[imagelayers-badge]: https://badge.imagelayers.io/otobrglez/rss-machine:latest.svg 
[imagelayers]: https://imagelayers.io/?images=otobrglez/rss-machine:latest 
[docker-pulls-badge]: https://img.shields.io/docker/pulls/otobrglez/rss-machine.svg
[docker-stars-badge]: https://img.shields.io/docker/stars/otobrglez/rss-machine.svg
[rss-machine]: http://github.com/otobrglez/rss-machine
[me]: https://github.com/otobrglez

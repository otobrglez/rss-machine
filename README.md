# rss-machine

[rss-machine][rss-machine] is web service for parsing RSS channels written in Go.

[![Build Status][circle-ci-badge]][circle-ci]
[![ImageLayers][imagelayers-badge]][imagelayers]
[![Docker Pulls][docker-pulls-badge]][docker-hub]
[![Docker Stars][docker-stars-badge]][docker-hub]
	

Supported types feed types `RSS 1.0`, `RSS 2.0` and `ATOM 1.0`.

Service uses [SlyMarbo/rss](https://github.com/SlyMarbo/rss) for parsing RSS feeds.

## Running from Docker

```
docker run -ti -p 5000:5000 otobrglez/rss-machine

curl <docker_host>:5000/parse?url=http://rss.cnn.com/rss/edition.rss
```

JSON response

```json
{
    "description": "CNN.com delivers up-to....",
    "image": {
        "height": 0,
        "title": "CNN.com - Top Stories",
        "url": "http://i.cdn.turner.com/cnn/.e/img/1.0/logo/cnn.logo.rss.gif",
        "width": 0
    },
    "items": [
            {
                "content": "A Florida civil j....",
                "date": "2016-03-21T12:26:43Z",
                "enclosures": null,
                "id": "http://edition.cnn.com/2016/03/20/opinions/hulk-hogan-verdict-warning-shot-media-opinion-callan/index.html",
                "link": "http://rss.cnn.com/c/35494/f/676993/s/4e65a3c7/sc/13/l/0Ledition0Bcnn0N0C20A160C0A30C20A0Copinions0Chulk0Ehogan0Everdict0Ewarning0Eshot0Emedia0Eopinion0Ecallan0Cindex0Bhtml0Deref0Fedition/story01.htm",
                "read": false,
                "summary": "",
                "title": "Hulk Hogan verdict body-slams Gawker",
            }
            # ... 
    ],
    "link": "http://edition.cnn.com/index.html?eref=edition",
    "nickname": "",
    "refresh": "2016-03-21T12:41:28.162216767Z",
    "title": "CNN.com - Top Stories",
    "unread": 25,
    "updateurl": "http://rss.cnn.com/rss/edition.rss"
}     
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
[docker-hub]: https://hub.docker.com/r/otobrglez/rss-machine/

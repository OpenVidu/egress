# Changelog

Notable changes of the OpenVidu fork of `egress`, which is versioned in lockstep with OpenVidu.
The changes of the upstream project are listed in its GitHub releases.

## 3.10.0

### Added

- A stream url can carry a tag, `#ov-stream-tag=<tag>`, which the egress reports after the redacted url
  of the stream and never sends to the RTMP server.

### Improved

- The default template asks for each camera at the size it draws it, so recording no longer makes every participant publish their 720p layer.

### Fixed

- Room Composite and Web Egress no longer fails with "websocket url timeout reached" when chrome takes more than 20 s to start, as its first launch can on a slow or busy disk: it now gets 60 s.
- The first Room Composite / Web Egress after the egress starts no longer has to read chrome from disk: the container reads it into the page cache in the background shortly after startup.

## 3.9.0

### Fixed

- Chrome no longer contacts Google services in the background.

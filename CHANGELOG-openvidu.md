# Changelog

Notable changes of the OpenVidu fork of `egress`, which is versioned in lockstep with OpenVidu.
The changes of the upstream project are listed in its GitHub releases.

## 3.10.0

### Added

- A stream url can carry a tag, `#ov-stream-tag=<tag>`, which the egress reports after the redacted url
  of the stream and never sends to the RTMP server.

## 3.9.0

### Fixed

- Chrome no longer contacts Google services in the background.

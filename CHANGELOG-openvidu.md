# Changelog

Notable changes of the OpenVidu fork of `egress`, which is versioned in lockstep with OpenVidu.
The changes of the upstream project are listed in its GitHub releases.

## [Unreleased]

### Fixed

- **Chrome no longer contacts Google services in the background**: the Chrome that renders web and room composite egresses reached Safe Browsing, the component updater, the optimization guide, the network time service, the default search engine, the Google account listing and the push messaging service at every launch. It now starts with those features off and a managed policy without a default search provider, and the account and push messaging hosts never resolve. Recorded pages still load their own resources; only a "Sign in with Google" button cannot load.

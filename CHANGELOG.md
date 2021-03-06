# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]
### Added
- game names are now checked case insensitive
- if no comment set, use mumble name

### Changed
### Deprecated
### Removed
### Fixed
### Security


## [0.3.0]
### Added
- Notifications im mumble if in game player name and mumble user name don't match.

### Fixed
- Player leaving mumble no longer causes the bot to go into a panic.
- A player leaving or disconnecting from the game no longer causes the bot to do an update.

## [0.2.2]
### Fixed
- Fixed moving player to dead when voted out for real this time.
- New Nil pointer dereference bug fixed.

## [0.2.1]
### Fixed
- Players always get moved to the dead channel when voted out.
- Nil pointer dereference bug fixed.
- Fixed certificate commands in README.

## [0.2.0]
### Added
- Allow passing botname, cert, key, server as parameters from config file.

### Changed
- Updated README to include setup instructions.

## [0.1.0]
### Added
- Initial Release
# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2017-08-11
### Information
**The API has been largely changed, so code compiled against latest version will not compile without changes!**
### Added
- Time frame filters for fixtures.
- Competition - will replace SoccerSeason in compliance with the documentation of Version 1 of the Football-Data API.
- Introduce FixtureStatus_{Postponed,Scheduled,Timed}.
- Set methods for HttpClient and Token.
- Fluent-style configuration methods as alternatives to the Set{HttpClient,Token} methods.
### Changed
- Fix invalid return value of LeagueTable methods.
- Fix typo in naming of "LastUpdated" field in SoccerSeason (now Competition) struct.
- Fix #6 where some fields are sometimes not decoded properly or at all.
- Fix `TeamsOfCompetition` using the wrong URL to fetch data.
### Known issues
- There is no way to configure the wanted data format (`minified`/`full`/`compressed`). This will be provided in some way in an upcoming version.

[Unreleased]: https://github.com/icedream/go-footballdata/compare/v0.1.0...develop
[0.1.0]: https://github.com/icedream/go-footballdata/compare/055b2a4227d16de49751e27660ba4b172d334987...v0.1.0
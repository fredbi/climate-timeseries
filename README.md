# climate-timeseries `[en]`

[fr](README.fr.md)` | `[en](README.md)

## Primer

> _Disclaimer_
>
> As I am scrambling to achieve a first MVP, most of the writing here is
> pretty much unpolished and rushed. Please bear with typo and lack of clarity
> in this early drafting phase.

This project is all about the construction of a database and an API
to expose the data series used by [TheShiftProject](www.theshiftproject.org).

In particular, the ambition is to bring easy access to all the pieces of data
that contributed to the [Plan de Transformation de l'Economie Française](https://ilnousfautunplan.fr/).

As an ancillary goal, this project aims to substitute a richer data-sharing platform to 
[the Shift Project's data portal](https://www.theshiftdataportal.org).

Eventually, we'd like to store and distribute data from Europe used by the book
[Décabornons! 9 propositions pour que l'Europe change d'ère](https://www.amazon.com/dp/2738138802).

## How to contribute
TODO

## API

This project exposes APIs through REST endpoints.

There are 2 APIs:
* a public API to contribute and consume time series
* an admin API to tend to nomenclatures (units, geographical areas, sectors...)

### Available API clients

* [go](clients/go)
* [R](clients/R)
* [python](clients/python)
* [typescript-node](clients/typescript-node)
* ...

> Prioritize R to be used primarily for data visualization.
>
> Investigate API integration for Jupyter.

TODO: figure out how to publish these package on each language repo (R, pypi, npm)

### Authentication

* Want to authenticate admins and public contributors from oauth2 identity provider (e.g. google, github, facebook...)
* Want to manage API keys for access from apps (e.g. R/python data visualization & analysis clients)

TODO

## Project

### Desirable features

* search timeseries
  * by theme
  * 
### API concepts

* series: time series about anything, mostly yearly series
* versioning: time series exist in different versions and are published with a semantic version tag
* contributors: any contributor can ingest a new time series or may fork an existing one to create its own version
* publication control: versioned time series are subject to some validation/review workflow before being made available to the public
* tags: a few things can be freely tagged with keywords are become searchable
* unit conversion: the API comes with a universal unit converter to store and render measurements with virtually any unit system
* nomenclatures: we use a lot of nomenclatures to classify and search time series
   * constant: mathematical and physical constants
   * mdimension: base measured dimensions
   * mdomain: domains that pertain to measurements
   * measurement: physical and economic measurements
   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)
   * munit: measurement units
   * musystem: systems of measurement units
   * ostatus: owner statuses
   * owner: series owner
   * period: time series periods (e.g. monthly, yearly...)
   * role: series owner role
   * source: data sources
   * status: series and versions statuses
   * theme: climate themes
   * zone: geographical zones
   * ztype: zone types


### Stack

_draft/scratchpad notes & reminders_

* API flavor: REST, OpenAPI 2.0
* API implementation: golang
* API boiler-plate: goswagger (with its quirks and flexibility, yes)
* DB: Postgres
* hosting & deployment: Heroku? (TBD)
* CI: TBD
* UI: TBD (react?)
* documentation 
  * API documentation: redoc
  * project documentation: gihub wiki
* social: TBD
* non-go clients: swagger.io? Let's see how this plays out for R...

### Micro stack
* Internal pkg communication
  * repo interface
  * pivot data structure is the one generated for OpenAPI.
    Note: goswagger don't play well with struct tags.
  * future APIs: stand ready to open to GraphQL & grpc, but still with OpenAPI models a

* Micro-servicing. A lot of TBD
  * Ivan's kit. I am still a fan. Let's see however what's new around, e.g. go-kit
  * tracing: try out the transition opentracing (check: transition from opencenus?). Work out how this works on heroku & al.
  * releasing: goreleaser. Want to continue experimenting with changie (gone diktat of semantic commits).
  * pg driver: jackc/sqlxi & al. are the best ones
  * SQL: don't know yet. Maybe some db-first ORM? 
    Don't want GORM whatsoever (code-first). 
    Maybe will just go with squirrel and jpmoiron/sqlx, the barebones way.
  * testing: bxcodec/faker, brianvoe/gofakeit. I am familiar with their codebase, so it is easy to fix or enhance them.

### Misc
* [DB creation and upgrade scripts](db/migrations)

### Foray into the future

Gathering useful data and distributing them to a wider audience is a first step. It is already not that easy, so let's do that first.

Producing an enticing UI on top of the API is something that could be contemplated, but decidedly not my chore business. In the "why not" category for now.

The ultimate goal, thanks to the contribution mechanism, is to introduce lineage tracking between series and transformations. That story shall also be told...

This feature would introduce the additional concept of a series transformation model: model Y produces series Y1 and Y2 from series X1,X2,X3, which in turn can be trace back to originating
from model Y... The objective is to bring clarity about how all these hypothesis and computation depend on each other, and what
would be the impact of altering some of the input data.

## TODOs

This project is still nascent. Everything is TODO...

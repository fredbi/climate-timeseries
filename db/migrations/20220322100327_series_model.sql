-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgis;
COMMENT ON EXTENSION postgis IS 'Add support for geometry and geospatial indexing';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
COMMENT ON EXTENSION "uuid-ossp" IS 'Add support for uuid extra type';

-- check if container enables this
-- CREATE EXTENSION IF NOT EXISTS semver;

CREATE TABLE nomenclatures(
    id serial NOT NULL PRIMARY KEY,
    short_key text NOT NULL UNIQUE,
    title json NOT NULL,
    description_short json NOT NULL DEFAULT '{}',
    description_long json NOT NULL DEFAULT '{}',
    audit_trail json NOT NULL DEFAULT '{}'
);
COMMENT ON TABLE nomenclatures IS 'A general purpose parent nomenclature to capture static parameterization, with some multi-language description';

CREATE VIEW all_nomenclatures AS
  SELECT 
    (tableoid::regclass)::text AS class,
    id, short_key,description_short,description_long,audit_trail
  FROM nomenclatures
;
COMMENT ON VIEW all_nomenclatures IS 'A view over all nomenclature tables, with the class exposed';

CREATE VIEW nomenclature_classes AS
  SELECT DISTINCT 
    (tableoid::regclass)::text AS class, 
    obj_description(tableoid::regclass)::text AS title
  FROM nomenclatures
;
COMMENT ON VIEW nomenclature_classes IS 'All available nomenclature classes';
  
-- TODO: link external users
CREATE TABLE owners(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    audit_trail json
);
INSERT INTO owners(name, email) VALUES
('fredbi', 'fredbi@yahoo.com')
;

CREATE TABLE roles(LIKE nomenclatures INCLUDING ALL);
COMMENT ON TABLE roles IS 'Contributors to timeseries and other objects contribute with some role. A role describes the activity of a contributor';

INSERT INTO roles(short_key, title) VALUES
('ADMIN', '{"fr": "administrateur", "en": "administrator"}'),
('CONTRIBUTOR', '{"fr": "contributeur", "en": "contributor"}'),
('AUTHOR', '{"fr": "auteur", "en": "author"}'),
('VALIDATOR', '{"fr": "valideur", "en": "validator"}'),
('NONE', '{"fr": "aucun rôle en particulier", "en": "no particular role"}')
;

CREATE TABLE time_periods(LIKE nomenclatures INCLUDING ALL);

INSERT INTO time_periods(short_key, title) VALUES
('YEARLY', '{"fr": "annuelle", "en": "yearly"}'),
('QUARTERLY', '{"fr": "trimestrielle", "en": "quarterly"}')
;

CREATE TABLE data_sources(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE data_sources
  ADD COLUMN
    ratings integer NOT NULL DEFAULT 0,
  ADD COLUMN
    tags jsonb NOT NULL DEFAULT '{}'
;

CREATE INDEX data_source_by_tag ON data_sources USING gin(tags);

CREATE TABLE zone_types(LIKE nomenclatures INCLUDING ALL);

INSERT INTO zone_types(short_key, title) VALUES
('WORLD', '{"fr": "monde", "en": "world"}'),
('COUNTRY', '{"fr": "pays", "en": "country"}'),
('REGION', '{"fr": "région du monde", "en": "region"}'),
('UNION', '{"fr": "union économique", "en": "economic union"}'),
('CITY', '{"fr": "ville", "en": "city"}'),
('AREA', '{"fr": "zone", "en": "area"}'),
('OTHER', '{"fr": "autres", "en": "others"}'),
('NONE', '{"fr": "sans type", "en": "without type"}')
;

CREATE TABLE zones(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE zones
  ADD COLUMN
    zone_type_id integer REFERENCES zone_types(id),
  ADD COLUMN
    zone_geometry geometry(Geometry,4326)
 ;

INSERT INTO zones(short_key, title,zone_type_id) VALUES
('WORLD' , '{"fr": "Monde", "en": "World"}', (SELECT id FROM zone_types WHERE short_key='WORLD')),
('FR'    , '{"fr": "France", "en": "France"}', (SELECT id FROM zone_types WHERE short_key='COUNTRY')),
('UK'    , '{"fr": "Grande-Bretagne", "en": "United Kingdom"}', (SELECT id FROM zone_types WHERE short_key='COUNTRY')),
('EUROPE', '{"fr": "Europe (incl. the UK)", "en": "Europe (incl. the UK)"}', (SELECT id FROM zone_types WHERE short_key='REGION')),
('EU'    , '{"fr": "European Union (incl. the UK before 2021)", "en": "European Union (incl. the UK before 2021)"}', (SELECT id FROM zone_types WHERE short_key='UNION')),
('NONE'  , '{"fr": "sans zone", "en": "without zone"}', (SELECT id FROM zone_types WHERE short_key='NONE'))
;

CREATE TABLE measurements(LIKE nomenclatures INCLUDING ALL);

INSERT INTO measurements(short_key, title) VALUES
('ANGLE', '{"fr": "angle", "en":"angle"}'),
('AREA', '{"fr": "surface", "en":"area"}'),
('CURRENCY', '{"fr": "devise", "en":"currency"}'),
('DISTANCE', '{"fr": "distance", "en":"distance"}'),
('ELECTRIC_CHARGE', '{"fr": "charge électrique", "en":"electric charge"}'),
('ENERGY', '{"fr": "énergie", "en":"energy"}'),
('MASS', '{"fr": "masse", "en":"mass"}'),
('POWER', '{"fr": "puissance", "en":"puissance"}'),
('PRESSURE', '{"fr": "pression", "en":"pressure"}'),
('TEMPERATURE', '{"fr": "temperature", "en":"temperature"}'),
('TIME', '{"fr": "temps", "en":"time"}'),
('VOLTAGE', '{"fr": "tension électrique", "en":"electric voltage"}'),
('VOLUME', '{"fr": "volume", "en":"volume"}'),
('WEIGHT', '{"fr": "poids", "en":"weight"}')
;

CREATE TABLE measurement_domains(LIKE nomenclatures INCLUDING ALL);

INSERT INTO measurement_domains(short_key, title) VALUES
('AERONAUTICS', '{"fr": "aviation", "en":"aeronautics"}'),
('AUTOMOTIVE', '{"fr": "automobile", "en":"automotive"}'),
('CHEMISTRY', '{"fr": "chimie", "en":"chemistry"}'),
('CLIMATE', '{"fr": "climat", "en":"climate"}'),
('COMMON', '{"fr": "communs", "en":"common"}'),
('CONSTRUCTION', '{"fr": "construction", "en":"construction"}'),
('ELECTRICITY', '{"fr": "électricité", "en":"electricity"}'),
('ENERGY', '{"fr": "énergie", "en":"energy"}'),
('ENGINEERING', '{"fr": "ingéniérie", "en":"engineering"}'),
('FLUIDS', '{"fr": "fluides", "en":"fluids"}'),
('LIGHT', '{"fr": "lumière", "en":"light"}'),
('MAGNETISM', '{"fr": "magnétisme", "en":"magnetism"}'),
('MATHS', '{"fr": "mathématiques", "en":"mathematics"}'),
('MECHANICS', '{"fr": "mécanique", "en":"mechanics"}'),
('NUTRITION', '{"fr": "nutrition", "en":"nutrition"}'),
('RADIOLOGY', '{"fr": "radiologie", "en":"radiology"}'),
('THERMAL', '{"fr": "thermique", "en":"thermal"}')
;

CREATE TABLE measurement_has_measurement_domains (
    measurement_id integer NOT NULL REFERENCES measurements(id),
    measurement_domain_id integer NOT NULL REFERENCES measurement_domains(id)
);

INSERT INTO measurement_has_measurement_domains(measurement_id, measurement_domain_id) VALUES
((SELECT id FROM measurements WHERE short_key = 'DISTANCE'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON')),
((SELECT id FROM measurements WHERE short_key = 'ANGLE'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON')),
((SELECT id FROM measurements WHERE short_key = 'MASS'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON')),
((SELECT id FROM measurements WHERE short_key = 'AREA'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON')),
((SELECT id FROM measurements WHERE short_key = 'VOLUME'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON')),
((SELECT id FROM measurements WHERE short_key = 'TIME'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON')),
((SELECT id FROM measurements WHERE short_key = 'TEMPERATURE'), (SELECT id FROM measurement_domains WHERE short_key = 'COMMON'))
;

CREATE TABLE measurement_unit_systems(LIKE nomenclatures INCLUDING ALL);

INSERT INTO measurement_unit_systems(short_key, title) VALUES
('METRIC', '{"fr": "système métrique", "en":"metric system"}'),
('IMPERIAL', '{"fr": "système impérial (US)", "en":"impérial system (US)"}'),
('BSU', '{"fr": "système britanique (UK)", "en":"British standard system (US)"}'),
('NONE', '{"fr": "aucun système", "en":"no system"}')
;

CREATE TABLE measurement_unit_multipliers (LIKE nomenclatures INCLUDING ALL);
ALTER TABLE measurement_unit_multipliers
  ADD COLUMN
    factor numeric NOT NULL DEFAULT 1
;

INSERT INTO measurement_unit_multipliers(short_key, title,factor) VALUES
('f', '{"fr": "femto", "en":"femto"}', 1E-12),
('n', '{"fr": "nano", "en":"nano"}', 1E-9),
('μ', '{"fr": "micro", "en":"micro"}', 1E-6),
('m', '{"fr": "milli", "en":"milli"}', 1E-3),
('c', '{"fr": "centi", "en":"centi"}', 1E-2),
('d', '{"fr": "deci", "en":"deci"}', 1E-1),
('', '{"fr": "", "en":""}', 1),
('da', '{"fr": "deca", "en":"deca"}', 1E1),
('h', '{"fr": "hecto", "en":"hecto"}', 1E2),
('k', '{"fr": "kilo", "en":"kilo"}', 1E3),
('M', '{"fr": "mega", "en":"mega"}', 1E6),
('G', '{"fr": "giga", "en":"giga"}', 1E9),
('P', '{"fr": "peta", "en":"peta"}', 1E12)
;

CREATE TABLE measurement_units(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE measurement_units
  ADD COLUMN
    measurement_id integer NOT NULL REFERENCES measurements(id),
  ADD COLUMN
    measurement_unit_systems_id integer REFERENCES measurement_unit_systems(id)
;

-- TODO: unit aliases
-- TODO: unit compounds e.g. m.s-1
INSERT INTO measurement_units(short_key, title, measurement_id) VALUES
-- distance
('m', '{"fr": "mètre", "en": "meter"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
('yd', '{"fr": "yard", "en": "yard"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
('mi', '{"fr": "mile (US)", "en": "mile (US)"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
('in', '{"fr": "pouce (US)", "en": "inch (US)"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
('ft', '{"fr": "pied (US)", "en": "foot (US)"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
('ly', '{"fr": "année-lumière", "en": "light year"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
-- volume
('m3', '{"fr": "mètre cube", "en": "cubic meter"}', (SELECT id FROM measurements WHERE short_key='VOLUME')),
('l', '{"fr": "litre", "en": "liter"}', (SELECT id FROM measurements WHERE short_key='VOLUME')),
('gal', '{"fr": "gallon (US)", "en": "gallon (US)"}', (SELECT id FROM measurements WHERE short_key='VOLUME')),
('bbl', '{"fr": "baril de pétrole", "en": "oil barrel"}', (SELECT id FROM measurements WHERE short_key='VOLUME')),
('bbl-us', '{"fr": "baril (US)", "en": "barrel (US)"}', (SELECT id FROM measurements WHERE short_key='VOLUME')),
('bbl-uk', '{"fr": "baril (UK)", "en": "barrel (UK)"}', (SELECT id FROM measurements WHERE short_key='VOLUME')),
-- energy & heat
('J', '{"fr": "joule", "en": "joule"}', (SELECT id FROM measurements WHERE short_key='ENERGY')),
('W.h', '{"fr": "watt.heure", "en": "watt.hour"}', (SELECT id FROM measurements WHERE short_key='ENERGY')),
('cal', '{"fr": "calories", "en": "calories"}', (SELECT id FROM measurements WHERE short_key='ENERGY')),
('btu', '{"fr": "British thermal unit", "en": "British thermal unit"}', (SELECT id FROM measurements WHERE short_key='ENERGY')),
('TeP', '{"fr": "tonne équivalent pétrole", "en": "ton of oil equivalent"}', (SELECT id FROM measurements WHERE short_key='ENERGY')),
-- power
('W', '{"fr": "watt", "en": "watt"}', (SELECT id FROM measurements WHERE short_key='POWER')),
('hp', '{"fr": "cheval", "en": "horse power"}', (SELECT id FROM measurements WHERE short_key='POWER')),
('ch', '{"fr": "cheval", "en": "horse power"}', (SELECT id FROM measurements WHERE short_key='POWER')),
-- mass
('g', '{"fr": "gramme", "en": "gram"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('q', '{"fr": "quintal", "en": "hundredweight (metric)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('cwt-uk', '{"fr": "quintal (UK)", "en": "hundredweight (UK)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('cwt-us', '{"fr": "quintal (US)", "en": "hundredweight (US)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('t', '{"fr": "tonne (métrique)", "en": "ton (metric)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('t-uk', '{"fr": "tonne (UK)", "en": "ton (UK)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('t-us', '{"fr": "tonne (US)", "en": "ton (US)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('lbs', '{"fr": "livre (US)", "en": "pound (US)"}', (SELECT id FROM measurements WHERE short_key='MASS')),
('u', '{"fr": "masse atomique", "en": "atomic mass"}', (SELECT id FROM measurements WHERE short_key='MASS')),
-- weight
('N', '{"fr": "newton", "en": "newton"}', (SELECT id FROM measurements WHERE short_key='WEIGHT')),
-- electricity
('C', '{"fr": "coulomb", "en": "coulomb"}', (SELECT id FROM measurements WHERE short_key='ELECTRIC_CHARGE'))
;

UPDATE measurement_units SET measurement_unit_systems_id = (SELECT id FROM measurement_unit_systems WHERE short_key = 'NONE');
ALTER TABLE measurement_units ALTER COLUMN measurement_unit_systems_id SET NOT NULL;


CREATE TABLE measurement_unit_has_conversions (
    from_unit_id integer NOT NULL REFERENCES measurement_units(id),
    to_unit_id integer NOT NULL REFERENCES measurement_units(id),
    factor numeric NOT NULL DEFAULT 1,
    intercept numeric NOT NULL DEFAULT 0,
    formula text,
    PRIMARY KEY(from_unit_id, to_unit_id)
);

CREATE TABLE themes(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE themes
  ADD COLUMN
    tags jsonb,  
  ADD COLUMN
    linked_documents json
;

CREATE TABLE object_statuses(LIKE nomenclatures INCLUDING ALL);

CREATE TABLE series (
    id serial PRIMARY KEY,
    title json NOT NULL,
    description_short json NOT NULL DEFAULT '{}',
    description_long json NOT NULL DEFAULT '{}',
    time_period_id integer NOT NULL REFERENCES time_periods(id),
    measurement_unit_id integer NOT NULL REFERENCES measurement_units(id),
    status_id integer NOT NULL REFERENCES object_statuses(id),
    zone_id integer NOT NULL REFERENCES zones(id),
    data_source_id integer NOT NULL REFERENCES data_sources(id),
    tags jsonb NOT NULL DEFAULT '{}',
    linked_documents json,
    audit_trail jsonb NOT NULL DEFAULT '{}'
);

CREATE TABLE series_has_themes (
    series_id integer NOT NULL REFERENCES series(id),
    theme_id integer NOT NULL REFERENCES themes(id),
    context_short json,
    context_long json,
    PRIMARY KEY(series_id, theme_id)
);

CREATE TABLE series_has_owners (
    series_id integer NOT NULL REFERENCES series(id),
    owner_id uuid NOT NULL REFERENCES owners(id),
    role_id integer NOT NULL REFERENCES roles(id),
    context_short json,
    context_long json,
    PRIMARY KEY(series_id, owner_id)
);

CREATE TABLE constants(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE constants
  ADD COLUMN
    symbol text NOT NULL,
  ADD COLUMN
    value numeric NOT NULL,
  ADD COLUMN
    measurement_unit_id integer REFERENCES measurement_units(id)
;

INSERT INTO constants(short_key, symbol, value, title, measurement_unit_id) VALUES
('PI',        'π', PI()                , '{"fr":"le nombre pi", "en":"the pi number"}'                                        , NULL),
('EULER',     'e', EXP(1)              , '{"fr":"la constante d''Euler", "en":"the Euler constant"}'                          , NULL),
('ELECTRON',  'e', -1.602176565 * 1E-19, '{"fr": "charge électrique de l''électron", "en": "eletrical charge of an electron"}', (SELECT id FROM measurement_units WHERE short_key='C'))
;

CREATE TABLE constant_has_measurement_domains (
    constant_id integer NOT NULL REFERENCES constants(id),
    measurement_domain_id integer REFERENCES measurement_domains(id),
    PRIMARY KEY(constant_id, measurement_domain_id)
);

CREATE TABLE series_produces_versions (
    series_id integer NOT NULL REFERENCES series(id),
    version text NOT NULL DEFAULT 'v0.0.0',
    -- version semver NOT NULL DEFAULT 'v0.0.0',
    versioned_series_id uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    version_owner_id uuid NOT NULL REFERENCES owners(id),
    version_status_id integer NOT NULL REFERENCES object_statuses(id),
    version_data_source_id integer NOT NULL REFERENCES data_sources(id),
    formula text,
    version_description_short json,
    version_description_long json,
    audit_trail json,
    PRIMARY KEY(series_id, version)
);

CREATE TABLE versioned_timeseries (
    versioned_series_id uuid NOT NULL REFERENCES series_produces_versions(versioned_series_id),
    started_at date NOT NULL,
    value numeric,
    values json,
    PRIMARY KEY(versioned_series_id, started_at)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

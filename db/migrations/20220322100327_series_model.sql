-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgis;
COMMENT ON EXTENSION postgis IS 'Add support for geometry and geospatial indexing';

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

CREATE VIEW nomenclature_classes AS
  SELECT DISTINCT 
    (tableoid::regclass)::text AS class, 
    obj_description(tableoid::regclass)::text AS title
  FROM nomenclatures
;
  
-- TODO: link external users
CREATE TABLE owners(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    audit_trail json
);

CREATE TABLE roles INHERITS nomenclatures;

INSERT INTO roles(short_key, title) VALUES
('CONTRIBUTOR', '{"fr": "contributeur", "en": "contributor"}'),
('AUTHOR', '{"fr": "auteur", "en": "author"}'),
('VALIDATOR', '{"fr": "valideur", "en": "validator"}'),
('NONE', '{"fr": "aucun rôle en particulier", "en": "no particular role"}'),
;

CREATE TABLE time_periods INHERITS nomenclatures;

CREATE TABLE data_sources(
    ratings integeri NOT NULL DEFAULT 0,
    tags jsonb NOT NULL DEFAULT '{}'
) INHERITS nomenclatures;

CREATE INDEX data_source_by_tag 
CREATE TABLE zone_types INHERITS nomenclatures;

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

CREATE TABLE zones(
    zone_type_id integer NOT NULL REFERENCES zone_types.id,
    zone_geometry geometry(Geometry,4326)
) INHERITS nomenclatures;

INSERT INTO zone(short_key, title,zone_type_id) VALUES
('WORLD', '{"fr": "Monde", "en": "World"}', (SELECT id FROM zone_types WHERE short_key='WORLD')),
('FR', '{"fr": "France", "en": "France"}', (SELECT id FROM zone_types WHERE short_key='COUNTRY')),
('UK', '{"fr": "Grande-Bretagne", "en": "United Kingdom"}', (SELECT id FROM zone_types WHERE short_key='COUNTRY')),
('EUROPE', '{"fr": "Europe (incl. the UK)", "en": "Europe (incl. the UK)"}', (SELECT id FROM zone_types WHERE short_key='REGION')),
('EU', '{"fr": "European Union (incl. the UK before 2021)", "en": "European Union (incl. the UK before 2021)"}', (SELECT id FROM zone_types WHERE short_key='UNION')),
('NONE', '{"fr": "sans zone", "en": "without zone"}', (SELECT id FROM zone_types WHERE short_key='NONE')),
;

CREATE TABLE measurements INHERITS nomenclatures;

INSERT INTO measurements(short_key, title) VALUES
('DISTANCE', {"fr": "distance", "en":"distance"}),
('ANGLE', {"fr": "angle", "en":"angle"}),
('MASS', {"fr": "masse", "en":"mass"}),
('VOLUME', {"fr": "volume", "en":"volume"}),
('AREA', {"fr": "surface", "en":"area"}),
('WEIGHT', {"fr": "poids", "en":"weight"}),
('ENERGY', {"fr": "énergie", "en":"energy"}),
('CURRENCY', {"fr": "devise", "en":"currency"}),
('TIME', {"fr": "temps", "en":"time"}),
('TEMPERATURE', {"fr": "temperature", "en":"temperature"}),
('POWER', {"fr": "puissance", "en":"puissance"}),
('PRESSURE', {"fr": "pression", "en":"pressure"})
;

CREATE TABLE measurement_domains INHERITS nomenclatures;

INSERT INTO measurement_domains(short_key, title) VALUES
('MATH', {"fr": "mathématiques", "en":"mathematics"}),
('COMMON', {"fr": "communs", "en":"common"}),
('ENERGY', {"fr": "énergie", "en":"energy"}),
('ENGINEERING', {"fr": "ingéniérie", "en":"engineering"}),
('ELECTRICITY', {"fr": "électricité", "en":"electricity"}),
('MAGNETISM', {"fr": "magnétisme", "en":"magnetism"}),
('CONSTRUCTION', {"fr": "construction", "en":"construction"}),
('NUTRITION', {"fr": "nutrition", "en":"nutrition"}),
('RADIOLOGY', {"fr": "radiologie", "en":"radiology"}),
('LIGHT', {"fr": "lumière", "en":"light"}),
('FLUIDS', {"fr": "fluides", "en":"fluids"}),
('MECANICS', {"fr": "mécanique", "en":"mechanics"}),
('THERMAL', {"fr": "thermique", "en":"thermal"}),
;

CREATE TABLE measurement_has_measurement_domains (
    measurement_id integer NOT NULL REFERENCES measurements.id,
    measurement_domain_id integer NOT NULL REFERENCES measurement_domains.id
);

CREATE TABLE unit_systems INHERITS nomenclatures;

INSERT INTO unit_systems(short_key, title) VALUES
('METRIC', {"fr": "système métrique", "en":"metric system"}),
('IMPERIAL', {"fr": "système impérial (US)", "en":"impérial system (US)"}),
('BSU', {"fr": "système brittanique (UK)", "en":"British standard system (US)"}),
('NONE', {"fr": "aucun système", "en":"no system"})
;

CREATE TABLE unit_multipliers (
    factor numeric NOT NULL DEFAULT 1
)INHERITS nomenclatures;

INSERT INTO unit_multipliers(short_key, title,factor) VALUES
('f', {"fr": "femto", "en":"femto"}, 1E-12),
('n', {"fr": "nano", "en":"nano"}, 1E-9),
('μ', {"fr": "micro", "en":"micro"}, 1E-6),
('m', {"fr": "milli", "en":"milli"}, 1E-3),
('c', {"fr": "centi", "en":"centi"}, 1E-2),
('d', {"fr": "deci", "en":"deci"}, 1E-1),
('', {"fr": "", "en":""}, 1),
('da', {"fr": "deca", "en":"deca"}, 1E1),
('h', {"fr": "hecto", "en":"hecto"}, 1E2),
('k', {"fr": "kilo", "en":"kilo"}, 1E3),
('M', {"fr": "mega", "en":"mega"}, 1E6),
('G', {"fr": "giga", "en":"giga"}, 1E9),
('P', {"fr": "peta", "en":"peta"}, 1E12),
;

CREATE TABLE units (
    measurement_id integer NOT NULL REFERENCES measurements.id,
    unit_systems_id integer NOT NULL REFERENCES unit_systems.id
) INHERITS nomenclatures;

INSERT INTO units(short_key, title, measurement_id) VALUES
-- distance
('m', '{"fr": "mètres", "en": "meters"}', (SELECT id FROM measurements WHERE short_key='DISTANCE')),
('yd', '{"fr": "yards", "en": "yards"}'),
('mi', '{"fr": "miles (US)", "en": "miles (US)"}'),
('in', '{"fr": "pouce (US)", "en": "inches (US)"}'),
('ft', '{"fr": "pieds (US)", "en": "feet (US)"}'),
('ly', '{"fr": "années-lumières", "en": "light years"}'),

-- volume
('m3', '{"fr": "mètres cubes", "en": "cubic meters"}'),
('l', '{"fr": "litres", "en": "liters"}'),
('gal', '{"fr": "gallons (US)", "en": "gallons (US)"}'),
('bbl', '{"fr": "barrils de pétrole", "en": "oil barrels"}'),
('bbl-us', '{"fr": "barils (US)", "en": "barrels (US)"}'),
('bbl-uk', '{"fr": "barils (UK)", "en": "barrels (UK)"}'),

-- energy & heat
('J', '{"fr": "joules", "en": "joules"}'),
('W', '{"fr": "watts", "en": "watts"}'),
('cal', '{"fr": "calories", "en": "calories"}'),
('btu', '{"fr": "British Thermal Units", "en": "British Thermal Units"}'),
('hp', '{"fr": "chevaux", "en": "horse power"}'),
('ch', '{"fr": "chevaux", "en": "horse power"}'),
('TeP', '{"fr": "tonnes équivalent pétrole", "en": "..."}'),

-- weight and mass
('g', '{"fr": "grammes", "en": "grams"}'),
('t', '{"fr": "tonnes", "en": "tons"}'),
('t', '{"fr": "tonnes (UK)", "en": "tons (UK)"}'),
('t', '{"fr": "tonnes (US)", "en": "tons (US)"}'),
('lbs', '{"fr": "livres (US)", "en": "pounds (US)"}'),
('u', '{"fr": "masses atomiques", "en": "atomic masses"}'),
('N', '{"fr": "newtons", "en": "newtons"}'),
;


CREATE TABLE measurement_unit_has_conversions (
    from_unit_id integer NOT NULL REFERENCES measurement_units.id,
    to_unit_id integer NOT NULL REFERENCES measurement_units.id,
    factor numeric NOT NULL DEFAULT 1,
    intercept numeric NOT NULL DEFAULT 0,
    formula text,
    PRIMARY KEY(from_unit_id, to_unit_id)
);

CREATE TABLE themes (
  tags jsonb,  
  linked_documents json
) INHERITS nomenclatures;

CREATE TABLE object_statuses INHERITS nomenclatures;

CREATE TABLE series (
    id serial PRIMARY KEY,
    title json NOT NULL,
    description_short json NOT NULL DEFAULT '{}',
    description_long json NOT NULL DEFAULT '{}',
    time_period_id integer NOT NULL REFERENCES time_periods.id,
    measurement_unit_id integer NOT NULL REFERENCES measurement_units.id,
    status_id integer NOT NULL REFERENCES object_statuses.id,
    zone_id integer NOT NULL REFERENCES zones.id,
    data_source_id integer NOT NULL REFERENCES data_sources.id,
    tags jsonb NOT NULL DEFAULT '{}',
    linked_documents json,
    audit_trail jsonb NOT NULL DEFAULT '{}'
);

CREATE TABLE series_has_themes (
    series_id integer NOT NULL REFERENCES series.id,
    theme_id integer NOT NULL REFERENCES themes.id,
    context_short json,
    context_long json,
    PRIMARY KEY(series_id, theme_id)
);

CREATE TABLE series_has_owners (
    series_id integer NOT NULL REFERENCES series.id,
    owner_id uuid NOT NULL REFERENCES owners.id,
    role_id integer NOT NULL REFERENCES roles.id,
    context_short json,
    context_long json,
    PRIMARY KEY(series_id, owner_id)
);

CREATE TABLE constants (
    symbol text NOT NULL PRIMARY KEY,
    value numeric,
) INHERITS nomenclatures;

INSERT INTO constants(short_key, symbol, value, description_short) VALUES
('PI', 'π', PI(), '{"fr":"le nombre pi", "en":"the pi number"}'),
('EULER', 'e', EXP(1), '{"fr":"la constante d''Euler", "en":"the Euler constant"}'),
('ELECTRON', 'e', , '{"fr": "charge électrique de l''électron", "en": "eletrical charge of an electron"}')
;

CREATE TABLE constant_has_measurement_domains (
    constant_id integer NOT NULL REFERENCES constants.id,
    measurement_domain_id integer REFERENCES measurement_domains.id,
    PRIMARY KEY(constant_id, measurement_domain_id)
);

CREATE TABLE series_produces_versions (
    series_id integer NOT NULL REFERENCES series.id,
    version text NOT NULL DEFAULT '{v0.0.0}',
    -- version semver NOT NULL DEFAULT '{v0.0.0}',
    versioned_series_id uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    version_owner_id uuid NOT NULL REFERENCES owners.id,
    version_status_id integer NOT NULL REFERENCES object_statuses.id,
    version_data_source_id integer NOT NULL REFERENCES data_sources.id,
    formula text,
    version_description_short json,
    version_description_long json,
    audit_trail json,
    PRIMARY KEY(series_id, version)
);

CREATE TABLE versioned_timeseries (
    versioned_series_id uuid NOT NULL REFERENCES series_has_versions.versioned_series_id,
    started_at datetime NOT NULL,
    value numeric,
    values json,
    PRIMARY KEY(versioned_series_id, started_at)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

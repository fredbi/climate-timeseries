-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgis;
COMMENT ON EXTENSION postgis IS 'Add support for geometry and geospatial indexing';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
COMMENT ON EXTENSION "uuid-ossp" IS 'Add support for uuid extra type';

CREATE EXTENSION IF NOT EXISTS "pg_trgm";
COMMENT ON EXTENSION "pg_trgm" IS 'Add support for simple text search, by text similarity';

-- TODO: enable full text search on desriptions
-- check if container enables this
-- extension not available on heroku
-- CREATE EXTENSION IF NOT EXISTS semver;

CREATE TABLE nomenclatures(
    id serial NOT NULL PRIMARY KEY,
    short_code text NOT NULL UNIQUE,
    title json NOT NULL,
    description_short json NOT NULL DEFAULT '{}',
    description_long json NOT NULL DEFAULT '{}',
    audit_trail json NOT NULL DEFAULT '{}'
);

COMMENT ON TABLE nomenclatures IS 'A template nomenclature table to capture static parameterization, with some multi-language description';
COMMENT ON COLUMN nomenclatures.short_code IS 'A unique mnemonic key for a value, normalized as: uppercased, plain ASCII-only';
COMMENT ON COLUMN nomenclatures.title IS 'A title object, lower-cased, available in different languages';
COMMENT ON COLUMN nomenclatures.audit_trail IS 'A modification tracking object (with UTC timings), formed as: {"created_at":"{date}","updated_at":"{date}","created_by":"{owner name}","last_updated_by":"{owner name}"}';

CREATE TABLE all_nomenclatures(
    class text NOT NULL PRIMARY KEY,
    table_name text NOT NULL,
    title json NOT NULL,
    description_short json NOT NULL DEFAULT '{}',
    description_long json NOT NULL DEFAULT '{}',
    metadata json NOT NULL DEFAULT '{}'
);

COMMENT ON TABLE nomenclatures IS 'A list of all nomenclature classes, with some UI metadata consumed by admin interfaces';

INSERT INTO all_nomenclatures(class, table_name, title, metadata) VALUES
  ('OWNER', 'owners', '{"fr": "propriétaires de données", "en" :"data owners"}',
     '{"from_template":false}'),
  ('ROLE', 'roles', '{"fr": "rôles utilisateur", "en" :"user roles"}',
     '{"from_template":true}'),
  ('PERIOD', 'time_periods', '{"fr": "périodes temporelles", "en" :"time period"}',
     '{"from_template":true}'),
  ('SOURCE', 'data_sources', '{"fr": "sources de données", "en" :"data sources"}',
     '{"from_template":true, "extra_fields": ["rating", "tags"], "tag_search": true}'),
  ('ZONE', 'zones', '{"fr": "zonages", "en" :"zoning"}',
     '{"from_template":true, "extra_fields": ["zone_type_id", "zone_geometry"], "has_one_class": {"ZTYPE": "zone_type_id"}}'),
  ('ZTYPE', 'zone_types', '{"fr": "types de zones", "en" :"zone types"}',
     '{"from_template":true}'),
  ('MEASUREMENT', 'measurements', '{"fr": "grandeurs physiques", "en" :"physical measurements"}',
     '{"from_template":true, "has_zero_many_class": {"MDOMAIN": "measurement_has_measurement_domains"}}'),
  ('MDOMAIN', 'measurement_domains', '{"fr": "domaines de mesure", "en" :"measurement domains"}',
     '{"from_template":true}'),
  ('MDIMENSIONS', 'measurement_dimensions', '{"fr": "dimensions de mesures physiques", "en" :"physical measurements dimensions"}',
     '{"from_template":true}'),
  ('MUSYSTEM', 'measurement_unit_systems', '{"fr": "systèmes d''unités de mesure", "en" :"measurement unit systems"}',
     '{"from_template":true}'),
  ('MUNIT', 'measurement_units', '{"fr": "unités de mesure", "en" :"measurement units"}',
     '{"from_template":true, "extra_fields": ["measurement_id", "measurement_unit_system_id"], "has_one_class":{"MEASUREMENT": "measurement_id", "MUSYSTEM": "measurement_unit_system_id"}}'),
  ('THEME', 'themes', '{"fr": "thématiques", "en" :"themes"}',
     '{"from_template":true, "extra_fields": ["tags", "linked_documents"], "has_one_class":{"MEASUREMENT": "measurement_id", "MUSYSTEM": "measurement_unit_system_id"}, "tag_search": true}'),
  ('STATUS', 'object_statuses', '{"fr": "statuts d''un object", "en" :"object statuses"}',
     '{"from_template":true}'),
  ('OSTATUS', 'owner_statuses', '{"fr": "statuts utilisateurs", "en" :"user statuses"}',
     '{"from_template":true}'),
  ('CONSTANT', 'constants', '{"fr": "constantes physiques et mathématiques", "en" :"physical and mathematical constants"}',
     '{"from_template":true, "extra_fields": ["symbol", "value", "metadata", "measurement_unit_id"], "has_zero_one_class": {"MUNIT": "measurement_unit_id"}}')
;

CREATE TABLE owner_statuses(LIKE nomenclatures INCLUDING ALL);

INSERT INTO owner_statuses(short_code, title) VALUES
  ('ACTIVE', '{"fr": "actif", "en": "active"}'),
  ('DISABLED', '{"fr": "désactivé", "en": "disabled"}'),
  ('LOCKED', '{"fr": "verrouillé", "en": "locked"}')
;

-- TODO: link external users
CREATE TABLE owners(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    owner_status_id integer NOT NULL REFERENCES owner_statuses(id),
    audit_trail json NOT NULL DEFAULT '{}'
);

COMMENT ON TABLE owners IS 'Registered owners of the timeseries data';

INSERT INTO owners(name, email, owner_status_id) VALUES
('admin', 'fredbi@yahoo.com', (SELECT id FROM owner_statuses WHERE short_code = 'ACTIVE'))
;

UPDATE owners SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE roles(LIKE nomenclatures INCLUDING ALL);

COMMENT ON TABLE roles IS 'Contributors to timeseries and other objects contribute with some role. A role describes the activity of a contributor';

INSERT INTO roles(short_code, title) VALUES
('ADMIN', '{"fr": "administrateur", "en": "administrator"}'),
('CONTRIBUTOR', '{"fr": "contributeur", "en": "contributor"}'),
('AUTHOR', '{"fr": "auteur", "en": "author"}'),
('VALIDATOR', '{"fr": "valideur", "en": "validator"}'),
('NONE', '{"fr": "aucun rôle en particulier", "en": "no particular role"}')
;

UPDATE roles SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE time_periods(LIKE nomenclatures INCLUDING ALL);
COMMENT ON TABLE owners IS 'Reference of legit periods for time series';

INSERT INTO time_periods(short_code, title) VALUES
('YEARLY', '{"fr": "annuelle", "en": "yearly"}'),
('QUARTERLY', '{"fr": "trimestrielle", "en": "quarterly"}'),
('MONTHLY', '{"fr": "mensuelle", "en": "monthly"}')
;

UPDATE time_periods SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE data_sources(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE data_sources
  ADD COLUMN
    ratings integer NOT NULL DEFAULT 0,
  ADD COLUMN
    tags json NOT NULL DEFAULT '{}'
;

INSERT INTO data_sources(short_code, title) VALUES
  ('GIEC', '{"fr": "GIEC", "en": "GIEC"}'),
  ('SHIFT', '{"fr": "TheShiftProject", "en": "TheShiftProject"}'),
  ('SHIFTERS', '{"fr": "Shifters (bénévoles)", "en": "Shifters (volunteers)"}'),
  ('RTE', '{"fr": "RTE (Réseau de Transport d''Electricité, France)", "en": "RTE (France''s energy transmission system operator)"}'),
  ('EDF', '{"fr": "EDF (Electricité de France)", "en": "EDF (France''s electricity generation utility)"}'),
  ('EGEDA', '{"fr": "Expert Group on Energy Data and Analysis (EGEDA)", "en": "Expert Group on Energy Data and Analysis (EGEDA)"}'),
  ('WORLDBANK', '{"fr": "Banque Mondiale", "en": "World Bank"}'),
  ('IMF', '{"fr": "Fond Monétaire International (FMI)", "en": "International Monetary Fund (IMF)"}'),
  ('IEA', '{"fr": "Agence Internationale de l''Energie (IEA)", en": "International Energy Agency (IEA)"}'),
  ('USEIA', '{"fr": "US Energy Information Administration (EIA)", en": "US Energy Information Administration (EIA)"}'),
  ('CAIT', '{"fr": "CAIT Climate Data Explorer (World Resources Institute)", en": "CAIT Climate Data Explorer (World Resources Institute)"}'),
  ('EDGAR', '{"fr": "EDGAR: Emissions Database for Global Atmospheric Research", "en": "EDGAR: Emissions Database for Global Atmospheric Research"}'),
  ('EORA', '{"fr": "Global Supply Chain Database (EORA)", "en": "Global Supply Chain Database (EORA)"}')
;

UPDATE data_sources SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

-- make tags searchable
CREATE INDEX data_sources_by_tag ON data_sources USING gist((tags::text) gist_trgm_ops);

CREATE TABLE zone_types(LIKE nomenclatures INCLUDING ALL);

INSERT INTO zone_types(short_code, title) VALUES
('WORLD', '{"fr": "monde", "en": "world"}'),
('COUNTRY', '{"fr": "pays", "en": "country"}'),
('REGION', '{"fr": "macro région", "en": "macro area"}'),
('UNION', '{"fr": "union économique ou polique", "en": "economic or political union"}'),
('CITY', '{"fr": "ville", "en": "city"}'),
('AREA', '{"fr": "région", "en": "area"}'),
('OTHER', '{"fr": "autres", "en": "others"}'),
('NONE', '{"fr": "sans type", "en": "without type"}')
;

UPDATE zone_types SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE zones(LIKE nomenclatures INCLUDING ALL);

ALTER TABLE zones
  ADD COLUMN
    zone_type_id integer REFERENCES zone_types(id),
  ADD COLUMN
    zone_geometry geometry(Geometry,4326)
;

INSERT INTO zones(short_code, title,zone_type_id) VALUES
('WORLD' , '{"fr": "Monde", "en": "World"}', (SELECT id FROM zone_types WHERE short_code='WORLD')),
--
('AT'    , '{"fr": "Autriche", "en": "Austria"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('BE'    , '{"fr": "Belgique", "en": "Belgium"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('BG'    , '{"fr": "Bulgarie", "en": "Bulgaria"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('CH'    , '{"fr": "Suisse", "en": "Switzerland"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('CZ'    , '{"fr": "République Tchèque", "en": "Czech Republic"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('DE'    , '{"fr": "Allemagne", "en": "Germany"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('DK'    , '{"fr": "Danemark", "en": "Denmark"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('ES'    , '{"fr": "Espagne", "en": "Spain"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('FI'    , '{"fr": "Finlande", "en": "Finland"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('FR'    , '{"fr": "France", "en": "France"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('GB'    , '{"fr": "Grande-Bretagne", "en": "United Kingdom"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('GR'    , '{"fr": "Grèce", "en": "Greece"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('HU'    , '{"fr": "Hongrie", "en": "Hungary"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('IE'    , '{"fr": "Irlande", "en": "Irland"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('IT'    , '{"fr": "Italy", "en": "Italy"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('NO'    , '{"fr": "Norvège", "en": "Norway"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('PL'    , '{"fr": "Pologne", "en": "Poland"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('PT'    , '{"fr": "Portugal", "en": "Portugal"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('RO'    , '{"fr": "Roumanie", "en": "Romania"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('SE'    , '{"fr": "Suède", "en": "Sweden"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('UA'    , '{"fr": "Ukraine", "en": "Ukraine"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('RU'    , '{"fr": "Russie", "en": "Russia"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('TR'    , '{"fr": "Turquie", "en": "Turkey"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('US'    , '{"fr": "Etats Unis", "en": "United States"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('MX'    , '{"fr": "Mexique", "en": "Mexico"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('CA'    , '{"fr": "Canada", "en": "Canada"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('CN'    , '{"fr": "Chine", "en": "China"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('IN'    , '{"fr": "Inde", "en": "India"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('AU'    , '{"fr": "Australie", "en": "Australia"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
('BR'    , '{"fr": "Brésil", "en": "Brazil"}', (SELECT id FROM zone_types WHERE short_code='COUNTRY')),
--
('EUROPE', '{"fr": "Europe (incl. the UK)", "en": "Europe (incl. the UK)"}', (SELECT id FROM zone_types WHERE short_code='REGION')),
('EU'    , '{"fr": "European Union (incl. the UK before 2021)", "en": "European Union (incl. the UK before 2021)"}', (SELECT id FROM zone_types WHERE short_code='UNION')),
('AFRICA', '{"fr": "Afrique", "en": "Africa"}', (SELECT id FROM zone_types WHERE short_code='REGION')),
('ASIA'  , '{"fr": "Asie", "en": "Asia"}', (SELECT id FROM zone_types WHERE short_code='REGION')),
('ASPA'  , '{"fr": "Asie-Pacifique", "en": "Asia-Pacific"}', (SELECT id FROM zone_types WHERE short_code='REGION')),
('MEAST' , '{"fr": "Moyen-Orient", "en": "Middle-East"}', (SELECT id FROM zone_types WHERE short_code='REGION')),
('NA'    , '{"fr": "Amerique du Nord", "en": "North America"}', (SELECT id FROM zone_types WHERE short_code='REGION')),
--
('NONE'  , '{"fr": "sans zone", "en": "without zone"}', (SELECT id FROM zone_types WHERE short_code='NONE'))
;

UPDATE zones SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurement_dimensions(LIKE nomenclatures INCLUDING ALL);

COMMENT ON TABLE measurement_dimensions IS 'Universal dimensions for physical and economic measurements';

INSERT INTO measurement_dimensions(short_code, title) VALUES
('L', '{"fr": "longueur", "en": "length"}'),
('M', '{"fr": "masse", "en": "mass"}'),
('T', '{"fr": "temps", "en": "time"}'),
('N', '{"fr": "quantité de matière", "en": "amount of substance"}'),
('I', '{"fr": "intensité électrique", "en": "electric intensity"}'),
('J', '{"fr": "intensité lumineuse", "en": "light intensity"}'),
('θ', '{"fr": "température", "en": "temperature"}'),
-- not a physical measurement, but it matters too
('$', '{"fr": "monnaie", "en": "money"}')
;

UPDATE measurement_dimensions SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurements(LIKE nomenclatures INCLUDING ALL);

ALTER TABLE measurements
  ADD COLUMN
    dimensions text
;

COMMENT ON TABLE measurements IS 'Physical and economic measurements';
COMMENT ON COLUMN measurements.dimensions IS 'Formula describing the dimension of the measurement, expressed in terms of universal dimensions (e.g. length, mass, time...)';

INSERT INTO measurements(short_code, title, dimensions) VALUES
-- dimension-less measurements
('QUANTITY', '{"fr": "quantité (unités)", "en":"amount (units)"}',NULL),
('ANGLE', '{"fr": "angle", "en":"angle"}',NULL),
-- non-physical: money
('CURRENCY', '{"fr": "devise", "en":"currency"}', '$'),
-- the 7 base measurements recognized by the international system
('LENGTH', '{"fr": "longueur (distance)", "en":"length (distance)"}', 'L'),
('TIME', '{"fr": "temps", "en":"time"}', 'T'),
('MASS', '{"fr": "masse", "en":"mass"}', 'M'),
('ELECTRIC_INTENSITY', '{"fr": "intensité électrique", "en":"electric intensity"}', 'I'),
('TEMPERATURE', '{"fr": "temperature", "en":"temperature"}', 'θ'),
('LIGHT_INTENSITY', '{"fr": "intensité lumineuse", "en":"light intensity"}', 'J'),
('SUBSTANCE', '{"fr": "quantité de matière", "en":"amount of substance"}', 'N'),
-- derived measurements
('AREA', '{"fr": "surface", "en":"area"}','L^2'),
('VOLUME', '{"fr": "volume", "en":"volume"}', 'L^3'),
('FREQUENCY', '{"fr": "fréquence", "en":"fréquence"}', 'T^-1'),
('VOLMASS', '{"fr": "masse volumique", "en":"volumic mass"}', 'M.L^-3'),
('ACCELERATION', '{"fr": "accélération", "en":"acceleration"}', 'L.T^-2'),
('ELECTRIC_CHARGE', '{"fr": "charge électrique", "en":"electric charge"}', 'T.I'),
('ENERGY', '{"fr": "énergie", "en":"energy"}', 'M.L^2.T^-2'),
('ENTHALPY', '{"fr": "enthalpie", "en":"enthalpy"}', 'M.L^2.T^-2'),
('ENTROPY', '{"fr": "entropie", "en":"entropy"}', 'M.L^2.T^-2.θ^-1'),
('POWER', '{"fr": "puissance", "en":"power"}', 'M.L^2.T^-3'),
('PRESSURE', '{"fr": "pression", "en":"pressure"}', 'M.L^-1.T^-2'),
('SPEED', '{"fr": "vitesse", "en":"speed"}', 'L.T^-1'),
('VOLTAGE', '{"fr": "tension électrique", "en":"electric voltage"}','M.L^2.T^-3.I^-1'),
('INDUCTION', '{"fr": "induction magnétique", "en":"magnetic induction"}','M.T^-2.I^-1'),
('WEIGHT', '{"fr": "poids", "en":"weight"}', 'M.L.T^-2'),
('FORCE', '{"fr": "force", "en":"force"}', 'M.L.T^-2'),
('TORQUE', '{"fr": "couple", "en":"torque"}', 'M.L^2.T^-2')
;

UPDATE measurements SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurement_domains(LIKE nomenclatures INCLUDING ALL);

INSERT INTO measurement_domains(short_code, title) VALUES
('AERONAUTICS', '{"fr": "aviation", "en":"aeronautics"}'),
('AUTOMOTIVE', '{"fr": "automobile", "en":"automotive"}'),
('CHEMISTRY', '{"fr": "chimie", "en":"chemistry"}'),
('CLIMATE', '{"fr": "climat", "en":"climate"}'),
('COMMON', '{"fr": "domaine commun", "en":"common domain"}'),
('CONSTRUCTION', '{"fr": "construction", "en":"construction"}'),
('ELECTRICITY', '{"fr": "électricité", "en":"electricity"}'),
('ENERGY', '{"fr": "énergie", "en":"energy"}'),
('ENGINEERING', '{"fr": "ingéniérie", "en":"engineering"}'),
('FINANCE', '{"fr": "finance", "en":"finance"}'),
('FLUIDS', '{"fr": "fluides", "en":"fluids"}'),
('LIGHT', '{"fr": "lumière", "en":"light"}'),
('MAGNETISM', '{"fr": "magnétisme", "en":"magnetism"}'),
('MATHS', '{"fr": "mathématiques", "en":"mathematics"}'),
('MECHANICS', '{"fr": "mécanique", "en":"mechanics"}'),
('NONE', '{"fr": "aucun domaine en particulier", "en":"no particular domain"}'),
('NUTRITION', '{"fr": "nutrition", "en":"nutrition"}'),
('RADIOLOGY', '{"fr": "radiologie", "en":"radiology"}'),
('THERMAL', '{"fr": "thermique", "en":"thermal"}')
;

UPDATE measurement_domains SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurement_has_measurement_domains (
    measurement_id integer NOT NULL REFERENCES measurements(id),
    measurement_domain_id integer NOT NULL REFERENCES measurement_domains(id),
    audit_trail jsonb NOT NULL DEFAULT '{}'
);

INSERT INTO measurement_has_measurement_domains(measurement_id, measurement_domain_id) VALUES
((SELECT id FROM measurements WHERE short_code = 'LENGTH'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
((SELECT id FROM measurements WHERE short_code = 'ANGLE'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
((SELECT id FROM measurements WHERE short_code = 'MASS'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
((SELECT id FROM measurements WHERE short_code = 'AREA'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
((SELECT id FROM measurements WHERE short_code = 'VOLUME'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
((SELECT id FROM measurements WHERE short_code = 'TIME'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
((SELECT id FROM measurements WHERE short_code = 'TEMPERATURE'), (SELECT id FROM measurement_domains WHERE short_code = 'COMMON')),
--
((SELECT id FROM measurements WHERE short_code = 'LENGTH'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'ANGLE'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'TIME'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'MASS'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'PRESSURE'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'ENERGY'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'FORCE'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS')),
((SELECT id FROM measurements WHERE short_code = 'TORQUE'), (SELECT id FROM measurement_domains WHERE short_code = 'MECHANICS'))
--
;
UPDATE measurement_has_measurement_domains SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurement_unit_systems(LIKE nomenclatures INCLUDING ALL);

INSERT INTO measurement_unit_systems(short_code, title) VALUES
('METRIC', '{"fr": "système métrique", "en":"metric system"}'),
('IMPERIAL', '{"fr": "système impérial (US)", "en":"impérial system (US)"}'),
('BSU', '{"fr": "système britanique (UK)", "en":"British standard system (US)"}'),
('DOMAIN', '{"fr": "unité usuelle", "en":"customary unit"}'),
('NONE', '{"fr": "aucun système", "en":"no system"}')
;

UPDATE measurement_unit_systems SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurement_unit_multipliers (LIKE nomenclatures INCLUDING ALL);
ALTER TABLE measurement_unit_multipliers
  ADD COLUMN
    factor numeric NOT NULL DEFAULT 1
;

INSERT INTO measurement_unit_multipliers(short_code, title,factor) VALUES
('', '{"fr": "", "en":""}', 1),
('a', '{"fr": "atto", "en":"atto"}', 1E-18),
('f', '{"fr": "femto", "en":"femto"}', 1E-15),
('p', '{"fr": "pico", "en":"pico"}', 1E-12),
('n', '{"fr": "nano", "en":"nano"}', 1E-9),
('μ', '{"fr": "micro", "en":"micro"}', 1E-6),
('m', '{"fr": "milli", "en":"milli"}', 1E-3),
('c', '{"fr": "centi", "en":"centi"}', 1E-2),
('d', '{"fr": "deci", "en":"deci"}', 1E-1),
('da', '{"fr": "deca", "en":"deca"}', 1E1),
('h', '{"fr": "hecto", "en":"hecto"}', 1E2),
('k', '{"fr": "kilo", "en":"kilo"}', 1E3),
('M', '{"fr": "mega", "en":"mega"}', 1E6),
('G', '{"fr": "giga", "en":"giga"}', 1E9),
('T', '{"fr": "tera", "en":"tera"}', 1E12),
('P', '{"fr": "peta", "en":"peta"}', 1E15),
('E', '{"fr": "hexa", "en":"hexa"}', 1E18)
;

UPDATE measurement_unit_multipliers SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

CREATE TABLE measurement_units(LIKE nomenclatures INCLUDING ALL);

COMMENT ON TABLE measurement_units IS 'Units for all quantities, amounts and measurements';

ALTER TABLE measurement_units
  ADD COLUMN
    measurement_id integer NOT NULL REFERENCES measurements(id),
  ADD COLUMN
    measurement_unit_system_id integer REFERENCES measurement_unit_systems(id),
  ADD COLUMN
    included_multiplier_id integer REFERENCES measurement_unit_multipliers(id),
  ADD COLUMN
    is_standard bool NOT NULL DEFAULT false,
  ADD COLUMN
    metadata json NOT NULL DEFAULT '{}'
;

COMMENT ON COLUMN measurement_units.included_multiplier_id IS 'This unit already factors in a multiplier of the base unit, e.g. km';
COMMENT ON COLUMN measurement_units.metadata IS 'Additional data for rendering, such as aliases as UTF-8 glyphs';

-- TODO: unit aliases
-- TODO: unit compounds e.g. m.s-1
INSERT INTO measurement_units(short_code, title, measurement_id) VALUES
-- distance
('m', '{"fr": "mètre", "en": "meter"}', (SELECT id FROM measurements WHERE short_code='LENGTH')),
('yd', '{"fr": "yard (US)", "en": "yard (US)"}', (SELECT id FROM measurements WHERE short_code='LENGTH')),
('mi', '{"fr": "mile (US)", "en": "mile (US)"}', (SELECT id FROM measurements WHERE short_code='LENGTH')),
('in', '{"fr": "pouce (US)", "en": "inch (US)"}', (SELECT id FROM measurements WHERE short_code='LENGTH')),
('ft', '{"fr": "pied (US)", "en": "foot (US)"}', (SELECT id FROM measurements WHERE short_code='LENGTH')),
('ly', '{"fr": "année-lumière", "en": "light year"}', (SELECT id FROM measurements WHERE short_code='LENGTH')),
-- angle
('rad', '{"fr": "radian", "en": "radian"}', (SELECT id FROM measurements WHERE short_code='ANGLE')),
('deg', '{"fr": "degré", "en": "degree"}', (SELECT id FROM measurements WHERE short_code='ANGLE')),
('minute', '{"fr": "minute d''angle", "en": "minute"}', (SELECT id FROM measurements WHERE short_code='ANGLE')),
('ste', '{"fr": "stéradian", "en": "steradian"}', (SELECT id FROM measurements WHERE short_code='ANGLE')),
-- area
('m2', '{"fr": "mètre carré", "en": "square meter"}', (SELECT id FROM measurements WHERE short_code='AREA')),
-- volume
('m3', '{"fr": "mètre cube", "en": "cubic meter"}', (SELECT id FROM measurements WHERE short_code='VOLUME')),
('l', '{"fr": "litre", "en": "liter"}', (SELECT id FROM measurements WHERE short_code='VOLUME')),
('gal', '{"fr": "gallon (US)", "en": "gallon (US)"}', (SELECT id FROM measurements WHERE short_code='VOLUME')),
('bbl', '{"fr": "baril de pétrole", "en": "oil barrel"}', (SELECT id FROM measurements WHERE short_code='VOLUME')),
('bbl-us', '{"fr": "baril (US)", "en": "barrel (US)"}', (SELECT id FROM measurements WHERE short_code='VOLUME')),
('bbl-uk', '{"fr": "baril (UK)", "en": "barrel (UK)"}', (SELECT id FROM measurements WHERE short_code='VOLUME')),
-- energy & heat
('J', '{"fr": "joule", "en": "joule"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('kJ', '{"fr": "kilojoule", "en": "kilojoule"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('MJ', '{"fr": "megajoule", "en": "megajoule"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('GJ', '{"fr": "gigajoule", "en": "gigajoule"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('eV', '{"fr": "électron-volt", "en": "electron-volt"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('erg', '{"fr": "erg", "en": "erg"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('W.h', '{"fr": "watt-heure", "en": "watt.hour"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('kW.h', '{"fr": "kilowatt-heure", "en": "kilowatt-hour"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('MW.h', '{"fr": "megawatt-heure", "en": "megawatt-hour"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('GW.h', '{"fr": "gigawatt-heure", "en": "gigawatt-hour"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('TW.h', '{"fr": "terawatt-heure", "en": "terawatt-hour"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('cal', '{"fr": "calories", "en": "calories"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('kcal', '{"fr": "kilocalories", "en": "kilocalories"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('Btu', '{"fr": "British thermal unit", "en": "British thermal unit"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('tep', '{"fr": "tonne équivalent pétrole", "en": "ton of oil equivalent"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
('tec', '{"fr": "tonne équivalent charbon", "en": "ton of coal equivalent"}', (SELECT id FROM measurements WHERE short_code='ENERGY')),
-- power
('W', '{"fr": "watt", "en": "watt"}', (SELECT id FROM measurements WHERE short_code='POWER')),
('J/s', '{"fr": "joule/seconde", "en": "joule/second"}', (SELECT id FROM measurements WHERE short_code='POWER')),
('VA', '{"fr": "volt-ampère", "en": "volt-ampere"}', (SELECT id FROM measurements WHERE short_code='POWER')),
('hp', '{"fr": "cheval-vapeur", "en": "horse power"}', (SELECT id FROM measurements WHERE short_code='POWER')),
('ch', '{"fr": "cheval", "en": "horse power"}', (SELECT id FROM measurements WHERE short_code='POWER')),
-- time
('s', '{"fr": "seconde", "en": "second"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('min', '{"fr": "minute", "en": "minute"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('h', '{"fr": "heure", "en": "hour"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('day', '{"fr": "jour", "en": "hour"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('wk', '{"fr": "semaine", "en": "week"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('mth', '{"fr": "mois", "en": "month"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('yr', '{"fr": "année", "en": "year"}', (SELECT id FROM measurements WHERE short_code='TIME')),
('cent', '{"fr": "siècle", "en": "century"}', (SELECT id FROM measurements WHERE short_code='TIME')),
-- mass
('kg', '{"fr": "kilogramme", "en": "kilogram"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('g', '{"fr": "gramme", "en": "gram"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('q', '{"fr": "quintal", "en": "hundredweight (metric)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('cwt-uk', '{"fr": "quintal (UK)", "en": "hundredweight (UK)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('cwt-us', '{"fr": "quintal (US)", "en": "hundredweight (US)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('t', '{"fr": "tonne (métrique)", "en": "ton (metric)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('t-uk', '{"fr": "tonne (UK)", "en": "ton (UK)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('t-us', '{"fr": "tonne (US)", "en": "ton (US)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('lbs', '{"fr": "livre (US)", "en": "pound (US)"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('u', '{"fr": "masse atomique", "en": "atomic mass"}', (SELECT id FROM measurements WHERE short_code='MASS')),
('TeCO2', '{"fr": "tonne équivalent CO2", "en": "ton of CO2 equivalent"}', (SELECT id FROM measurements WHERE short_code='MASS')),
-- substance
('mol', '{"fr": "mole", "en": "mole"}', (SELECT id FROM measurements WHERE short_code='SUBSTANCE')),
-- weight
('N', '{"fr": "newton", "en": "newton"}', (SELECT id FROM measurements WHERE short_code='WEIGHT')),
('daN', '{"fr": "deca newton", "en": "deca newton"}', (SELECT id FROM measurements WHERE short_code='WEIGHT')),
-- temperature
('K', '{"fr": "°Kelvin", "en": "°Kelvin"}', (SELECT id FROM measurements WHERE short_code='TEMPERATURE')),
('Cel', '{"fr": "°Celsius", "en": "°Celsius"}', (SELECT id FROM measurements WHERE short_code='TEMPERATURE')),
('Far', '{"fr": "°Farenheit", "en": "°Farenheit"}', (SELECT id FROM measurements WHERE short_code='TEMPERATURE')),
-- electricity
('A', '{"fr": "ampère", "en": "ampere"}', (SELECT id FROM measurements WHERE short_code='ELECTRIC_INTENSITY')),
('V', '{"fr": "volt", "en": "volt"}', (SELECT id FROM measurements WHERE short_code='VOLTAGE')),
('C', '{"fr": "coulomb", "en": "coulomb"}', (SELECT id FROM measurements WHERE short_code='ELECTRIC_CHARGE')),
-- speed
('m/s', '{"fr": "mètre/seconde", "en": "meter/second"}', (SELECT id FROM measurements WHERE short_code='SPEED')),
('rad/s', '{"fr": "radian/seconde", "en": "radian/second"}', (SELECT id FROM measurements WHERE short_code='SPEED')),
('km/h', '{"fr": "kilomètre/heure", "en": "kilometer/hour"}', (SELECT id FROM measurements WHERE short_code='SPEED')),
('mph', '{"fr": "mile/heure", "en": "mile/hour"}', (SELECT id FROM measurements WHERE short_code='SPEED')),
-- pressure
('Pa', '{"fr": "pascal", "en": "pascal"}', (SELECT id FROM measurements WHERE short_code='PRESSURE')),
('hPa', '{"fr": "hectopascal", "en": "hectopascal"}', (SELECT id FROM measurements WHERE short_code='PRESSURE')),
('bar', '{"fr": "bar", "en": "bar"}', (SELECT id FROM measurements WHERE short_code='PRESSURE')),
('millibar', '{"fr": "millibar", "en": "millibar"}', (SELECT id FROM measurements WHERE short_code='PRESSURE')),
-- light
('cd', '{"fr": "candela", "en": "candela"}', (SELECT id FROM measurements WHERE short_code='LIGHT_INTENSITY'))
;

UPDATE measurement_units SET measurement_unit_system_id = (SELECT id FROM measurement_unit_systems WHERE short_code = 'NONE');
UPDATE measurement_units SET included_multiplier_id = (SELECT id FROM measurement_unit_multipliers WHERE short_code = '');

ALTER TABLE measurement_units ALTER COLUMN measurement_unit_system_id SET NOT NULL;
ALTER TABLE measurement_units ALTER COLUMN included_multiplier_id SET NOT NULL;

UPDATE measurement_units SET is_standard = true
WHERE
short_code IN (
    'm', 's','kg','J', 'K', 'A', 'mol',
    'W', 'm/s' , 'rad/s', 'Pa', 'cd' ,'V', 'N', 'C'
);

UPDATE measurement_units SET audit_trail = json_build_object(
    'created_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'created_by', 'admin', 'updated_at', to_char(current_timestamp,'YYYYMMDDHHMI'), 'last_updated_by', 'admin'
);

-- TODO UPDATE measurement_units SET multiplier_included=

CREATE TABLE measurement_unit_has_conversions (
    from_unit_id integer NOT NULL REFERENCES measurement_units(id),
    to_unit_id integer NOT NULL REFERENCES measurement_units(id),
    factor numeric NOT NULL DEFAULT 1,
    intercept numeric NOT NULL DEFAULT 0,
    formula text,
    reverse_formula text,
    audit_trail json,
    PRIMARY KEY(from_unit_id, to_unit_id)
);

CREATE INDEX measurement_unit_conversions_reverse_idx ON measurement_unit_has_conversions(to_unit_id, from_unit_id);

CREATE TABLE themes(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE themes
  ADD COLUMN
    tags json,
  ADD COLUMN
    linked_documents json
;

COMMENT ON TABLE themes IS 'Themes for climate change analysis';
COMMENT ON COLUMN themes.short_code IS 'Theme short code key, uppercased, in the form of a hierarchical path, e.g. /{root theme}/{subtheme}[/...]';
COMMENT ON COLUMN themes.tags IS 'Searchable tags, per language. Form is {"{lang}": [{"tag1"}, {"tag2"}...}]';

INSERT INTO themes(short_code, title, tags) VALUES
  ('/ENERGY', '{"fr": "énergie", "en": "energy"}', '{"fr": ["énergie","génération", "électricité", "fossile", "carburant", "centrale"], "en": ["energy", "power", "electricity", "fossile", "fuel", "power plant"]}'),
  ('/ENERGY/NUCLEAR', '{"fr": "énergie nucléaire", "en": "nuclear power"}', '{"fr": ["nucléaire", "atomique", "atome", "uranium", "électricité" ], "en": ["nuclear", "atomic", "atom", "uranium", "electricity"]}'),
  ('/ENERGY/GAZ', '{"fr": "gaz", "en": "gas"}', NULL),
  ('/ENERGY/OIL', '{"fr": "pétrole", "en": "pétrole"}', NULL),
  ('/ENERGY/RENEWABLES', '{"fr": "énergies renouvelables", "en": "renewable energy"}', NULL),
  ('/INDUSTRY', '{"fr": "industrie", "en": "industry"}', NULL),
  ('/INDUSTRY/METALLURGY', '{"fr": "métallurgie", "en": "metallurgy"}', NULL),
  ('/INDUSTRY/CHEMISTRY', '{"fr": "chimie", "en": "chemistry"}', NULL),
  ('/INDUSTRY/CEMENTS', '{"fr": "ciments", "en": "cements"}', NULL),
  ('/INDUSTRY/AUTOMOTIVE', '{"fr": "industrie automobile", "en": "automotive industry"}', NULL),
  ('/INDUSTRY/BATTERIES', '{"fr": "industrie des batteries", "en": "batteries industry"}', NULL),
  ('/INDUSTRY/OTHERS', '{"fr": "autres industries", "en": "other industries"}', NULL),
  ('/MOBILITY', '{"fr": "mobilité", "en": "mobility"}', NULL),
  ('/MOBILITY/COMMUTING', '{"fr": "mobilité quotidienne", "en": "commuting"}', NULL),
  ('/MOBILITY/LONG', '{"fr": "mobilité longue distance", "en": "long-distance mobility"}', NULL),
  ('/FREIGHT', '{"fr": "frêt", "en": "freight"}', NULL),
  ('/CONSTRUCTION', '{"fr": "construction", "en": "construction"}', NULL),
  ('/CONSTRUCTION/HOMES', '{"fr": "logement", "en": "housing"}', NULL),
  ('/CONSTRUCTION/CITIES', '{"fr": "villes et territoires", "en": "cities and land usage"}', NULL),
  ('/AGRO', '{"fr": "agriculture", "en": "agriculture"}', NULL),
  ('/AGRO/INDUSTRY', '{"fr": "industrie agro-alimentaire", "en": "agro-industry"}', NULL),
  ('/AGRO/FOREST', '{"fr": "exploitation forestière", "en": "forestry"}', NULL),
  ('/ECONOMY', '{"fr": "économie", "en": "economy"}', NULL),
  ('/ECONOMY/EMPLOYMENT', '{"fr": "emploi", "en": "employment"}', NULL),
  ('/ECONOMY/FINANCE', '{"fr": "finance", "en": "finance"}', NULL),
  ('/HEALTHCARE', '{"fr": "santé", "en": "healthcare"}', NULL),
  ('/CULTURAL', '{"fr": "culture", "en": "cultural"}', NULL),
  ('/DIGITAL', '{"fr": "usages numériques", "en": "digital"}', NULL)
;

-- make tags searchable
CREATE INDEX themes_by_tag ON data_sources USING gist((tags::text) gist_trgm_ops);

CREATE TABLE object_statuses(LIKE nomenclatures INCLUDING ALL);

COMMENT ON TABLE object_statuses IS 'Object status describes the stage an object has reached in the publication workflow';

INSERT INTO object_statuses(short_code, title) VALUES
  ('VALIDATED', '{"fr": "validé", "en": "validated"}'),
  ('PUBLISHED', '{"fr": "publié", "en": "published"}'),
  ('REJECTED', '{"fr": "rejeté", "en": "rejected"}'),
  ('PENDINGV', '{"fr": "en attente de validation", "en": "pending validation"}'),
  ('PENDINGP', '{"fr": "en attente de publication", "en": "pending publication"}')
;

CREATE TABLE series (
    id serial PRIMARY KEY,
    title json NOT NULL,
    description_short json NOT NULL DEFAULT '{}',
    description_long json NOT NULL DEFAULT '{}',
    time_period_id integer NOT NULL REFERENCES time_periods(id),
    measurement_unit_id integer NOT NULL REFERENCES measurement_units(id),
    input_composed_unit json,
    status_id integer NOT NULL REFERENCES object_statuses(id),
    status_change_reason json NOT NULL DEFAULT '{}',
    zone_id integer NOT NULL REFERENCES zones(id),
    data_source_id integer NOT NULL REFERENCES data_sources(id),
    tags json NOT NULL DEFAULT '{}',
    linked_documents json,
    audit_trail jsonb NOT NULL DEFAULT '{}'
);

-- make tags searchable
CREATE INDEX series_by_tag ON series USING gist((tags::text) gist_trgm_ops);

-- make audit_trail searchable
CREATE INDEX series_by_audit ON series USING gin(audit_trail);

CREATE TABLE series_has_themes (
    series_id integer NOT NULL REFERENCES series(id),
    theme_id integer NOT NULL REFERENCES themes(id),
    context_short json,
    context_long json,
    audit_trail jsonb NOT NULL DEFAULT '{}',
    PRIMARY KEY(series_id, theme_id)
);

CREATE TABLE series_has_owners (
    series_id integer NOT NULL REFERENCES series(id),
    owner_id uuid NOT NULL REFERENCES owners(id),
    role_id integer NOT NULL REFERENCES roles(id),
    context_short json,
    context_long json,
    audit_trail jsonb NOT NULL DEFAULT '{}',
    PRIMARY KEY(series_id, owner_id)
);

CREATE TABLE constants(LIKE nomenclatures INCLUDING ALL);
ALTER TABLE constants
  ADD COLUMN
    symbol text NOT NULL,
  ADD COLUMN
    value numeric NOT NULL,
  ADD COLUMN
    metadata json NOT NULL DEFAULT '{}',
  ADD COLUMN
    measurement_unit_id integer REFERENCES measurement_units(id),
  ADD CONSTRAINT constants_symbol_unique UNIQUE(symbol, measurement_unit_id)
;

COMMENT ON TABLE constants IS 'Mathematical and physical constants that may be used in formulae and conversions';
COMMENT ON COLUMN constants.symbol IS 'Usual symbol for this constant, unique for a given measurement unit';
COMMENT ON COLUMN constants.measurement_unit_id IS 'The measurement unit used for that constant. Mathematical constants don''t have an unit';
COMMENT ON COLUMN constants.metadata IS 'Additional data for rendering, such as aliases as UTF-8 glyphs';

INSERT INTO constants(short_code, symbol, value, title, measurement_unit_id) VALUES
('PI',        'π', PI()                , '{"fr":"le nombre pi", "en":"the pi number"}'                                        , NULL),
('EULER',     'e', EXP(1)              , '{"fr":"la constante d''Euler (base du logarithme népérien)", "en":"the Euler constant (naperian logarithm)"}', NULL),
('ELECTRON',  'e', -1.602176565 * 1E-19, '{"fr": "charge électrique de l''électron", "en": "eletrical charge of an electron"}', (SELECT id FROM measurement_units WHERE short_code='C'))
-- https://fr.wikipedia.org/wiki/Constante_physique#Liste_de_constantes_physiques
-- AVOGADRO
-- BOLTZMANN
-- PLANCK
-- LIGHT CELERITY
-- Coulomb
-- constante gravitationelle
-- pression atmosphérique standard
-- pesanteur terrestre au niveau de la mer
-- perméabilité magnétique du vide
-- permittivité diéléctrique du vide
-- impédance caractéristique du vide
-- ...
;

CREATE TABLE constant_has_measurement_domains (
    constant_id integer NOT NULL REFERENCES constants(id),
    measurement_domain_id integer REFERENCES measurement_domains(id),
    audit_trail json NOT NULL DEFAULT '{}',
    PRIMARY KEY(constant_id, measurement_domain_id)
);

CREATE TABLE series_produces_versions (
    series_id integer NOT NULL REFERENCES series(id),
    version text NOT NULL DEFAULT 'v0.0.0',
    versioned_series_id uuid NOT NULL DEFAULT uuid_generate_v4() UNIQUE,
    version_owner_id uuid NOT NULL REFERENCES owners(id),
    version_status_id integer NOT NULL REFERENCES object_statuses(id),
    version_data_source_id integer NOT NULL REFERENCES data_sources(id),
    version_parent_versioned_id uuid REFERENCES series_produces_versions(versioned_series_id),
    formula text,
    version_status_change_reason json NOT NULL DEFAULT '{}',
    version_title json,
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
    collapsed_values json,
    PRIMARY KEY(versioned_series_id, started_at)
);

COMMENT ON TABLE versioned_timeseries IS 'Holds all versions of all timeseries';
COMMENT ON COLUMN versioned_timeseries.value IS 'For simple series, the numerical value relevant over the period (started_at, started_at+{series period})';
COMMENT ON COLUMN versioned_timeseries.values IS 'For vector series and non-numerical series, the value relevant over the period (started_at, started_at+{series period}). Values is of the following form, all keys being optional: {"v": [{numeric, ...}]}, "k": [{text values, ...}, {key}: {value}...]';
COMMENT ON COLUMN versioned_timeseries.collapsed_values IS 'A timeseries collapsed in an array on a single row. Supported formats: [{"{date}": {value}}, ...] (ISO date), [{"{date}": {{values object}}}]';

CREATE VIEW versioned_timeseries_bounds AS
  SELECT
    versioned_series_id,
    -- TODO: case expression for collapsed series
    MIN(started_at) AS min_started_at,
    MAX(started_at) AS max_started_at
  FROM
    versioned_timeseries
  GROUP BY versioned_series_id
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

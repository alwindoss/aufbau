-- PRIMARY KEY(composite key)
-- org_id
-- entity_id
-- config_id
-- OTHER FIELDS
-- configuration VARCHAR(100)
CREATE TABLE config (
    org_id uuid NOT NULL,
    entity_id uuid NOT NULL,
    config_id VARCHAR(100) NOT NULL,
    config VARCHAR(100),
    PRIMARY KEY(org_id, entity_id, config_id)
);

INSERT INTO config (org_id, entity_id, config_id, config) values ('909a88ba-74c5-4a07-8e94-6687747d64af', '1e1239d6-d67c-43c3-ad89-30cfe2e56653', 'config_1', 'config_1_value');

SELECT * from config;
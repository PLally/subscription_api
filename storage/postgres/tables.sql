CREATE TABLE destinations(
    id SERIAL PRIMARY KEY NOT NULL,
    external_identifier TEXT NOT NULL,
    destination_type TEXT NOT NULL,
    UNIQUE (destination_type, external_identifier)
);

CREATE TABLE subscription_types(
    id SERIAL PRIMARY KEY NOT NULL,
    type TEXT NOT NULL,
    tags TEXT NOT NULL DEFAULT '',
    UNIQUE (type, tags)
);

CREATE TABLE subscriptions(
    id SERIAL PRIMARY KEY NOT NULL,
    subscription_type BIGINT NOT NULL REFERENCES subscription_types,
    destination BIGINT NOT NULL REFERENCES destinations,
    UNIQUE (subscription_type, destination)
);

ALTER TABLE subscriptions ADD  COLUMN last_item bigint NOT NULL default 0;
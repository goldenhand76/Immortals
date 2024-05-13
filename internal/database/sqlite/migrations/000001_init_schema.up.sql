CREATE TABLE IF NOT EXISTS nodes (
	id bigserial PRIMARY KEY, 
	name VARCHAR, 
	agent_id INTEGER, 
	client_id VARCHAR UNIQUE, 
	is_online INTEGER
);

CREATE TABLE IF NOT EXISTS sensors (
    id bigserial PRIMARY KEY,
    name VARCHAR,
    topic VARCHAR UNIQUE,
    node_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS actuators (
    id bigserial PRIMARY KEY,
    name VARCHAR,
    topic VARCHAR UNIQUE,
    node_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE sensors ADD FOREIGN KEY (node_id) REFERENCES nodes (id);

ALTER TABLE actuators ADD FOREIGN KEY (node_id) REFERENCES nodes (id);

CREATE INDEX ON sensors (topic);

CREATE INDEX ON actuators (topic);
CREATE TABLE IF NOT EXISTS
    nfe (
        id SERIAL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        raw_xml TEXT,
	    nfe_id VARCHAR(255),
	    cnpj   VARCHAR(255),
	    vnf    VARCHAR(255),
        PRIMARY KEY (id)
    );
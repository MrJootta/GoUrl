CREATE TABLE IF NOT EXISTS shortlinks.url (
	code VARCHAR(6) NOT NULL,
	url varchar(255) NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;
CREATE UNIQUE INDEX url_code_IDX USING BTREE ON shortlinks.url (code);

CREATE TABLE IF NOT EXISTS shortlinks.visits (
     code varchar(6) NOT NULL,
     `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)
    ENGINE=InnoDB
    DEFAULT CHARSET=latin1
    COLLATE=latin1_swedish_ci;
CREATE INDEX visits_code_IDX USING BTREE ON shortlinks.visits (code);
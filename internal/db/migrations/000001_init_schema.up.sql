START TRANSACTION;

CREATE TABLE IF NOT EXISTS check_infos (
    id bigint NOT NULL AUTO_INCREMENT,
    check_type tinyint unsigned NOT NULL,
    block_number bigint unsigned NOT NULL,
    block_hash char(64) NOT NULL,
    created_at                  datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP   COMMENT 'created time',
    updated_at                  datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    PRIMARY KEY (id),
    KEY index_check_infos_on_block_number (block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

COMMIT;

CREATE TABLE Warehouses (
    id BIGINT AUTO_INCREMENT,
    user_id VARCHAR(64) NOT NULL,
    currency VARCHAR(32) NOT NULL,
    amount  BIGINT      NOT NULL DEFAULT 0,
    version BIGINT      NOT NULL DEFAULT 0,  -- optimistic lock
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_id, currency)
);
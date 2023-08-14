CREATE TABLE IF NOT EXISTS Photo (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    imageUrl VARCHAR(2083) NOT NULL,
    shootingDate DATE NOT NULL,
    createdAt DATE NOT NULL,
    updatedAt DATE NOT NULL,

    PRIMARY KEY(id)
)
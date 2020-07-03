CREATE TABLE user (
    `id` BIGINT(20) UNSIGNED NOT NULL,
--     `uid` VARCHAR(255) NOT NULL DEFAULT '',

    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `password_hash` VARBINARY(254) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY (name),
--     UNIQUE KEY (uid),

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE entry (
    `id` BIGINT UNSIGNED NOT NULL,

    `url` VARBINARY(512) NOT NULL,
    `title` VARCHAR(512) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY (url(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

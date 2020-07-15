CREATE TABLE user (
    `id` BIGINT UNSIGNED NOT NULL,

    `name` VARCHAR(32) NOT NULL DEFAULT '',
    `user_name` VARCHAR(32) NOT NULL DEFAULT '',
    `email` VARCHAR(255) NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY (name),

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE study_group (
    `id` BIGINT UNSIGNED NOT NULL,

    `name` VARCHAR(255) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE event (
    `id` BIGINT UNSIGNED NOT NULL,

    `study_group_id` BIGINT UNSIGNED NOT NULL,

    `name` VARCHAR(255) NOT NULL,
    `event_date` DATETIME(6) NOT NULL,
    `description` VARCHAR(5000) NOT NULL,
    `place` VARCHAR(255) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE membership (
    `id` BIGINT UNSIGNED NOT NULL,

    `user_id` BIGINT UNSIGNED NOT NULL,
    `study_group_id` BIGINT UNSIGNED NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE,

    FOREIGN KEY (study_group_id)
        REFERENCES study_group(id)
        ON DELETE CASCADE,

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE administrator (
    `id` BIGINT UNSIGNED NOT NULL,

    `user_id` BIGINT UNSIGNED NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE,

    FOREIGN KEY (event_id)
        REFERENCES event(id)
        ON DELETE CASCADE,

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE participant (
    `id` BIGINT UNSIGNED NOT NULL,

    `user_id` BIGINT UNSIGNED NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE,

    FOREIGN KEY (event_id)
        REFERENCES event(id)
        ON DELETE CASCADE,

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE speaker (
    `id` BIGINT UNSIGNED NOT NULL,

    `user_id` BIGINT UNSIGNED NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE,

    FOREIGN KEY (event_id)
        REFERENCES event(id)
        ON DELETE CASCADE,

    KEY (created_at),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
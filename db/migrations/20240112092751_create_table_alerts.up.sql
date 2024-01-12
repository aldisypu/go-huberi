CREATE TABLE todos
(
    id          BIGINT       NOT NULL auto_increment,
    background_color       VARCHAR(7) NOT NULL DEFAULT #22c55e,
    highlight_color       VARCHAR(7) NOT NULL DEFAULT #9333ea,
    text_color       VARCHAR(7) NOT NULL DEFAULT #ffffff,
    duration    INT NOT NULL DEFAULT 5000,
    is_border BOOLEAN NOT NULL DEFAULT true,
    is_unit BOOLEAN NOT NULL DEFAULT false,
    created_at  bigint    NOT NULL,
    updated_at  bigint    NOT NULL,
    PRIMARY KEY (id)
) engine = innodb;
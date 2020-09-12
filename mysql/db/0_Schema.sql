DROP DATABASE IF EXISTS isuumo;
CREATE DATABASE isuumo;

DROP TABLE IF EXISTS isuumo.estate;
DROP TABLE IF EXISTS isuumo.chair;

CREATE TABLE isuumo.estate
(
    id          INTEGER             NOT NULL PRIMARY KEY,
    name        VARCHAR(64)         NOT NULL,
    description VARCHAR(4096)       NOT NULL,
    thumbnail   VARCHAR(128)        NOT NULL,
    address     VARCHAR(128)        NOT NULL,
    latitude    DOUBLE PRECISION    NOT NULL,
    longitude   DOUBLE PRECISION    NOT NULL,
    rent        INTEGER             NOT NULL,
    door_height INTEGER             NOT NULL,
    door_width  INTEGER             NOT NULL,
    features    VARCHAR(64)         NOT NULL,
    popularity  INTEGER             NOT NULL,
    INDEX `lat_long_pop` (`latitude`, `longitude`) USING BTREE, 
    INDEX `reverse_lat_long` (`longitude`, `latitude`) USING BTREE, 
    INDEX `door_height_key` (`door_height` ) USING BTREE,
    INDEX `door_width_key` (`door_width` ) USING BTREE,
    -- INDEX `rent_key` (`rent` ) USING BTREE,
    INDEX `pop` (`popularity` ) USING BTREE,
    -- INDEX `features_30` (`features`(30))
    FULLTEXT `fulltext_in_features` (`features`) 
);

    -- INDEX `popularity_key` (`popularity` ) USING BTREE,
    -- INDEX `lang_long_ASC` (`latitude`  ASC , `longitude`ASC ) USING BTREE,
    -- INDEX `long_lang_ASC` (`longitude` ASC , `latitude` ASC ) USING BTREE,
    -- INDEX `lang_long_DSC` (`latitude`  DESC ,`longitude`ASC ) USING BTREE,
CREATE TABLE isuumo.chair
(
    id          INTEGER         NOT NULL PRIMARY KEY,
    name        VARCHAR(64)     NOT NULL,
    description VARCHAR(4096)   NOT NULL,
    thumbnail   VARCHAR(128)    NOT NULL,
    price       INTEGER         NOT NULL,
    height      INTEGER         NOT NULL,
    width       INTEGER         NOT NULL,
    depth       INTEGER         NOT NULL,
    color       VARCHAR(64)     NOT NULL,
    features    VARCHAR(64)     NOT NULL,
    kind        VARCHAR(64)     NOT NULL,
    popularity  INTEGER         NOT NULL,
    stock       INTEGER         NOT NULL,
    INDEX `price_key` (`price`),
    INDEX `height_width` (`height`, `width`),
    INDEX `depth` (`depth`)
);

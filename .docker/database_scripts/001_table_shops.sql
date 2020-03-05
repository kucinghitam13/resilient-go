BEGIN;
    CREATE SEQUENCE IF NOT EXISTS "public"."shops_id_seq"
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
COMMIT;

BEGIN;
    CREATE TABLE IF NOT EXISTS "public"."shops" (
        "shop_id" BIGINT DEFAULT nextval("shop_id_seq"::regclass) NOT NULL,
        "name" VARCHAR(30) NOT NULL,
        "city" VARCHAR(30) NOT NULL,
        "create_time" Timestamp Without Time Zone,
        "update_time" Timestamp Without Time Zone,
        PRIMARY KEY ("shop_id")
    )

    COMMENT ON COLUMN shops.shop_id "unique identification for shop";
    COMMENT ON COLUMN shops.name "name of the shop";
COMMIT;
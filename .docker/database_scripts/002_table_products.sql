USE "playground_go";

BEGIN;
    CREATE SEQUENCE IF NOT EXISTS "public"."products_id_seq"
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
COMMIT;


BEGIN;
    CREATE TABLE IF NOT EXISTS "public"."products" (
        "product_id" BIGINT DEFAULT nextval("products_id_seq"::regclass) NOT NULL,
        "shop_id" BIGINT NOT NULL,
        "name" VARCHAR(70) NOT NULL,
        "description" VARCHAR(256),
        "price" NUMERIC(15,2) NOT NULL,
        "stock" SMALLINT NOT NULL DEFAULT 0,
        "create_time" Timestamp Without Time Zone,
        "update_time" Timestamp Without Time Zone,
        PRIMARY KEY ("product_id")
    );

    COMMENT ON COLUMN products.product_id "identification for production";
    COMMENT ON COLUMN products.shop_id "shop identifaction to which this product belongs to, coming from shop service";
    COMMENT ON COLUMN products.name "name of the product";
    COMMENT ON COLUMN products.description "description about the product";
    COMMENT ON COLUMN products.price "price of the product";
    COMMENT ON COLUMN products.stock "available stock of the product";
COMMIT;

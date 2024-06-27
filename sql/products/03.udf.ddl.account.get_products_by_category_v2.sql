CREATE OR REPLACE FUNCTION account.get_products_by_category_v2(category_name text)
 RETURNS TABLE(product_id uuid, product_name character varying, category text, price numeric, description text)
 LANGUAGE plpgsql
AS $function$
BEGIN
    RETURN QUERY
    SELECT p.product_id, p.product_name, p.category, p.price, p.description
    FROM account.products p
    WHERE lower(p.category) = lower(category_name);
END;
$function$
;

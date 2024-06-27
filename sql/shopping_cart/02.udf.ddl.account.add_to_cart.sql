CREATE OR REPLACE FUNCTION account.add_to_cart(params jsonb)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
    INSERT INTO account.shopping_cart (
    	user_id, 
    	product_id, 
    	quantity
    )
    VALUES (
        (params->>'user_id')::uuid,
        (params->>'product_id')::uuid,
        COALESCE((params->>'quantity')::integer, 1) 
    );
END;
$function$
;

CREATE OR REPLACE FUNCTION account.create_product(params jsonb)
 RETURNS uuid
 LANGUAGE plpgsql
AS $function$
DECLARE
    ret_id UUID;
BEGIN
 
    INSERT INTO account.products (
    	product_name, 
    	category, 
    	price,
    	description
    )
    VALUES (
    	params->>'product_name',
    	params->>'category',
   		(params->>'price')::numeric,
   		params->>'description'
    )
    RETURNING product_id INTO ret_id;

    RETURN ret_id;
END;
$function$
;

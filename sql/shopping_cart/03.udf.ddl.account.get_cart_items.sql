CREATE OR REPLACE FUNCTION account.get_cart_items(p_user_id uuid)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result jsonb;
BEGIN
    SELECT jsonb_build_object(
        'cart_items', jsonb_agg(jsonb_build_object(
            'cart_id', sc.cart_id,
            'product_id', sc.product_id,
            'product_name', p.product_name,
            'category', p.category,
            'price', p.price,
            'quantity', sc.quantity
        ))
    )
    INTO result
    FROM account.shopping_cart sc
    JOIN account.products p ON sc.product_id = p.product_id
    WHERE sc.user_id = p_user_id;

    RETURN COALESCE(result, '{}'::jsonb);
END;
$function$
;

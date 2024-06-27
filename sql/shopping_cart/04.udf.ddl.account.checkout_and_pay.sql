CREATE OR REPLACE FUNCTION account.checkout_and_pay(params jsonb)
 RETURNS uuid
 LANGUAGE plpgsql
AS $function$
DECLARE
    v_transaction_id UUID;
    p_user_id UUID;
    p_payment_method character varying;
BEGIN
    p_user_id := (params ->> 'user_id')::uuid;
    p_payment_method := params ->> 'payment_method';


    INSERT INTO account.transactions (user_id, total_amount, payment_method)
    SELECT
        p_user_id,
        SUM(p.price * sc.quantity),
        p_payment_method
    FROM account.shopping_cart sc
    JOIN account.products p ON sc.product_id = p.product_id
    WHERE sc.user_id = p_user_id
    GROUP BY p_user_id, p_payment_method
    RETURNING transaction_id INTO v_transaction_id;

    INSERT INTO account.transaction_details (transaction_id, product_id, quantity, price)
    SELECT
        v_transaction_id,
        sc.product_id,
        sc.quantity,
        p.price
    FROM account.shopping_cart sc
    JOIN account.products p ON sc.product_id = p.product_id
    WHERE sc.user_id = p_user_id;

    DELETE FROM account.shopping_cart WHERE user_id = p_user_id;

    RETURN v_transaction_id;
END;
$function$
;

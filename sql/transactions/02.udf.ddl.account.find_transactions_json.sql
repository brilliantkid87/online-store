CREATE OR REPLACE FUNCTION account.find_transactions_json(params jsonb)
 RETURNS json
 LANGUAGE plpgsql
AS $function$
DECLARE
    transactions_data JSON;
BEGIN
    SELECT json_agg(user_transactions)
    INTO transactions_data
    FROM (
        SELECT json_build_object(
            'user_id', u.user_id,
            'username', u.user_name ,
            'email', u.email,
            'transactions', (
                SELECT json_agg(
                    json_build_object(
                        'transaction_id', t.transaction_id,
                        'timestamp', t.transaction_date,
                        'amount', t.total_amount,
                        'payment_method', t.payment_method,
                        'products', (
                            SELECT json_agg(
                                json_build_object(
                                    'product_id', p.product_id,
                                    'product_name', p.product_name,
                                    'category', p.category,
                                    'price_per_unit', td.price,
                                    'quantity', td.quantity
                                )
                            )
                            FROM account.transaction_details td
                            LEFT JOIN account.products p ON td.product_id = p.product_id
                            WHERE td.transaction_id = t.transaction_id
                        )
                    )
                )
                FROM account.transactions t
                WHERE t.user_id = u.user_id
            )
        ) AS user_transactions
        FROM account.users u
        WHERE
            (params ->> 'user_id' IS NULL OR u.user_id = (params ->> 'user_id')::UUID)
    ) AS users_data;

    RETURN transactions_data;
END;
$function$
;

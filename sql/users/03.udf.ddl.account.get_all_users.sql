CREATE OR REPLACE FUNCTION account.get_all_users(params jsonb)
 RETURNS jsonb
 LANGUAGE plpgsql
AS $function$
DECLARE
    result jsonb;
BEGIN
    SELECT jsonb_agg(jsonb_build_object(
        'user_id', users.user_id,
        'email', users.email,
        'user_name', users.user_name,
        'password', users."password"
    ))
    INTO result
    FROM account.users users
    WHERE (params ->> 'email' IS NULL OR users.email = params ->> 'email')
      AND (params ->> 'user_name' IS NULL OR users.user_name = params ->> 'user_name');

    RETURN COALESCE(result, '[]'::jsonb);
END;
$function$
;

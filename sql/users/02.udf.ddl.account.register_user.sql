CREATE OR REPLACE FUNCTION account.register_user(params jsonb)
 RETURNS uuid
 LANGUAGE plpgsql
AS $function$
declare 
	ret_id UUID;
BEGIN
	INSERT INTO account.users (
		email,
		user_name,
		"password" 
	)
    VALUES (
    	params->>'email', 
    	params->>'user_name',
    	params->>'password'
    )
    RETURNING user_id INTO ret_id;


    RETURN ret_id;
END;
$function$
;

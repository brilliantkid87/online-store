CREATE OR REPLACE FUNCTION account.remove_from_cart(params jsonb)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
DECLARE
  v_user_id UUID;
  v_cart_id UUID;
  v_cart_exists BOOLEAN;
BEGIN

  SELECT (params ->> 'user_id')::UUID INTO v_user_id;
  SELECT (params ->> 'cart_id')::UUID INTO v_cart_id;

  IF v_user_id IS NULL OR v_cart_id IS NULL THEN
    RAISE EXCEPTION 'Missing required parameters: user_id or cart_id';
  END IF;

  SELECT EXISTS (
    SELECT 1 FROM account.shopping_cart
    WHERE user_id = v_user_id AND cart_id = v_cart_id
  ) INTO v_cart_exists;

  IF NOT v_cart_exists THEN
    RAISE EXCEPTION 'User does not have a cart with cart_id: %', v_cart_id;
  END IF;

  DELETE FROM account.shopping_cart
  WHERE user_id = v_user_id
    AND cart_id = v_cart_id;
END;
$function$
;

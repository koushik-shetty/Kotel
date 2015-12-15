DO
$$
DECLARE
   varA text := lower(':cmd_line_arg') ;
BEGIN
   PERFORM something(varA);
END
$$;
CREATE OR REPLACE FUNCTION public.update_mhsw(
    id integer,
    name text,
    address text,
    old text)
  RETURNS boolean AS
$BODY$
begin 
update mhsw
set name = coalesce(update_mhsw.name, mhsw.name),
    address = coalesce(update_mhsw.address, mhsw.address),
    old = coalesce(update_mhsw.old, mhsw.old) 
WHERE mhsw.id = update_mhsw.id;
return found;
end;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION public.update_mhsw(integer, text, text, text)
  OWNER TO postgres;
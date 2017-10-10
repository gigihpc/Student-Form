create or replace function search_mhsw
(
param text, 
fill text
) returns table
(
id int,
name text,
address text,
old text
) as $$
begin
if param = 'name' then
	return query select * from mhsw where mhsw.name = fill;
elsif param = 'address' then
	return query select * from mhsw where mhsw.address = fill;
elsif param = 'old' then
	return query select * from mhsw where mhsw.old = fill;
else 
	return query select * from mhsw where mhsw.id = cast((coalesce(fill,'0')) as integer);
end if;
end;
$$ language plpgsql;
create or replace function add_users2(name text, email text, password text, country text)
returns uuid as $$
declare id uuid;
begin
insert into users values(name, email, password, country) returning id;
end;
$$ language plpgsql;

select add_users2('yes!!','test','test', 'ok')

create or replace function update_mhsw(
id int,
name text,
address text,
old text) returns boolean as $$
begin 
update mhsw
set name = coalesce(update_mhsw.name, mhsw.name),
    address = coalesce(update_mhsw.address, mhsw.address),
    old = coalesce(update_mhsw.old, mhsw.old) 
WHERE mhsw.id = update_mhsw.id;
return found;
end;
$$ language plpgsql;

select update_mhsw(5, 'mark', 'italy', '23')

create table coba
(
id uuid primary key,
name text
)
create table coba1
(
id serial primary key,
name text
)
//generate uuid
SELECT md5(random()::text || clock_timestamp()::text)::uuid
insert into coba values(md5(random()::text || clock_timestamp()::text)::uuid,'test')
insert into coba1 (name) values('gigih')
insert into coba (name) values('master')
select * from users where name = 'gigih'

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

select * from search_mhsw('id','6')
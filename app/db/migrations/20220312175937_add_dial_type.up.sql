alter table dials
    add type int default 0 not null;

create index dials_type_index
    on dials (type);
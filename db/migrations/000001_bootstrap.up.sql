create table if not exists articles (
    id integer not null primary key,
    filename varchar(512) not null
);
create unique index if not exists idx_article_filename on articles (filename);

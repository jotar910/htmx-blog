create table if not exists articles (
    id integer primary key autoincrement,
    title   varchar(128) not null,
    filename varchar(512) not null,
	image   varchar(512) not null,
	summary varchar(1024) not null,
	timestamp int not null
);
create index if not exists idx_article_filename on articles (filename);

create table if not exists article_tags (
    id integer primary key autoincrement,
    article_id int not null,
    tag_name varchar(128) not null,
    foreign key (article_id) references articles(id)
);

create table if not exists article_carousel (
    id integer primary key autoincrement,
    article_id int not null,
    foreign key (article_id) references articles(id)
);

create table if not exists article_highlights (
    id integer primary key autoincrement,
    article_id int not null,
    foreign key (article_id) references articles(id)
);

create table if not exists article_views (
    id integer primary key autoincrement,
    article_id int not null,
    views bigint not null default 0,
    foreign key (article_id) references articles(id)
);


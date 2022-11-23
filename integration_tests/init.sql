-- SCHEMA START
create type semester as enum ('FALL', 'SPRING', 'SUMMER');

create type subscription_tier as enum ('BRONZE', 'SILVER', 'GOLD', 'PLATINUM');

create table sponsors
(
    id          serial,
    name        varchar           not null,
    tier        subscription_tier not null,
    since       date              not null,
    description varchar,
    website     varchar,
    logo_url    varchar,
    constraint sponsors_pk
        primary key (id, name)
);

create unique index sponsors_id_uindex
    on sponsors (id);

create unique index sponsors_name_uindex
    on sponsors (name);

create table terms
(
    id       serial,
    year     integer  not null,
    semester semester not null,
    constraint terms_pk
        primary key (id)
);

create unique index terms_id_uindex
    on terms (id);

create table hackathons
(
    id         serial,
    term_id    serial,
    start_date timestamp not null,
    end_date   timestamp not null,
    constraint hackathons_pk
        primary key (id),
    constraint hackathons_terms_id_fk
        foreign key (term_id) references terms
);

create unique index hackathons_id_uindex
    on hackathons (id);

create unique index hackathons_term_id_uindex
    on hackathons (term_id);

create table pronouns
(
    id         serial,
    subjective varchar not null,
    objective  varchar not null,
    constraint pronouns_pk
        primary key (id)
);

create unique index pronouns_id_uindex
    on pronouns (id);

create table users
(
    id             serial,
    email          varchar not null,
    phone_number   varchar,
    last_name      varchar not null,
    age            integer,
    pronoun_id     integer,
    first_name     varchar not null,
    role           varchar not null,
    oauth_uid      varchar not null,
    oauth_provider varchar not null,
    constraint users_pk
        primary key (id, oauth_uid),
    constraint users_pronouns_id_fk
        foreign key (pronoun_id) references pronouns
);

create unique index users_email_uindex
    on users (email);

create unique index users_phone_number_uindex
    on users (phone_number);

create table hackathon_participants
(
    user_id       integer not null,
    hackathon_id  integer not null,
    accepted_date timestamp
);

create table hackathon_sponsors
(
    hackathon_id integer not null,
    sponsor_id   integer not null
);

create table events
(
    id           serial,
    hackathon_id integer   not null,
    location     varchar   not null,
    start_date   timestamp not null,
    end_date     timestamp not null,
    name         varchar   not null,
    description  varchar   not null,
    constraint events_pk
        primary key (id),
    constraint events_hackathons_id_fk
        foreign key (hackathon_id) references hackathons
);

-- SCHEMA END


-- INTEGRATION TEST DATA START

-- TestDatabaseRepository_GetSponsorWithQueryable & TestDatabaseRepository_GetSponsor
INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Billy Bob LLC'::varchar, 'PLATINUM'::subscription_tier, '2022-11-09'::date,
        'loves coding'::varchar, 'billybob.com'::varchar, null::varchar); -- ID = 1

-- TestDatabaseRepository_CreateSponsor, TestDatabaseRepository_UpdateWebsite, TestDatabaseRepository_UpdateSince, TestDatabaseRepository_UpdateTier MUTABLE
INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Joe Shmoe Woodworking'::varchar, 'BRONZE'::subscription_tier, '2022-10-09'::date,
        'does wood'::varchar, 'joeshmoe.com'::varchar, null::varchar); -- ID = 2

-- TestDatabaseRepository_GetSponsors
INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Microsoft'::varchar, 'PLATINUM'::subscription_tier, '2000-10-10'::date,
        'does stuff'::varchar, 'microsoft.com'::varchar, null::varchar); -- ID = 3

INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Apple'::varchar, 'GOLD'::subscription_tier, '2000-10-10'::date,
        'does stuff'::varchar, 'apple.com'::varchar, null::varchar); -- ID = 4

INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Bing'::varchar, 'PLATINUM'::subscription_tier, '2000-10-10'::date,
        'does stuff'::varchar, 'bing.com'::varchar, null::varchar); -- ID = 5

INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Oracle'::varchar, 'BRONZE'::subscription_tier, '2000-10-10'::date,
        'does stuff'::varchar, 'oracle.com'::varchar, null::varchar); -- ID = 6

INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('UrMom'::varchar, 'SILVER'::subscription_tier, '2000-10-10'::date,
        'does stuff'::varchar, 'urmom.com'::varchar, null::varchar); -- ID = 7

-- TestDatabaseRepository_UpdateSponsor

INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('abcdef'::varchar, 'SILVER'::subscription_tier, '2000-10-10'::date,
        'does stuff'::varchar, 'urmom.com'::varchar, null::varchar); -- ID = 8

-- TestDatabaseRepository_DeleteSponsor
INSERT INTO public.sponsors (name, tier, since, description, website, logo_url)
VALUES ('Johnson''s Reality'::varchar, 'PLATINUM'::subscription_tier, '2000-10-10'::date,
        'does games'::varchar, 'urmom.com'::varchar, null::varchar); -- ID = 9

-- INTEGRATION TEST DATA END
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
    id       serial
        constraint terms_pk
            primary key,
    year     integer  not null,
    semester semester not null
);

create unique index terms_id_uindex
    on terms (id);

create table hackathons
(
    id         serial
        constraint hackathons_pk
            primary key,
    term_id    serial
        constraint hackathons_terms_id_fk
            references terms,
    start_date timestamp not null,
    end_date   timestamp not null
);

create unique index hackathons_id_uindex
    on hackathons (id);

create unique index hackathons_term_id_uindex
    on hackathons (term_id);

create table pronouns
(
    id         serial
        constraint pronouns_pk
            primary key,
    subjective varchar not null,
    objective  varchar not null
);

create unique index pronouns_id_uindex
    on pronouns (id);

create table hackathon_sponsors
(
    hackathon_id integer not null
        constraint hackathon_sponsors_hackathons_null_fk
            references hackathons,
    sponsor_id   integer not null
        constraint hackathon_sponsors_sponsors_null_fk
            references sponsors (id)
);

create table events
(
    id           serial
        constraint events_pk
            primary key,
    hackathon_id integer   not null
        constraint events_hackathons_id_fk
            references hackathons,
    location     varchar   not null,
    start_date   timestamp not null,
    end_date     timestamp not null,
    name         varchar   not null,
    description  varchar   not null
);

create table api_keys
(
    user_id integer   not null
        constraint api_keys_pk
            primary key,
    key     varchar   not null,
    created timestamp not null
);

create table users
(
    id                  serial
        constraint users_pk
            primary key
        constraint users_api_keys_user_id_fk
            references api_keys
            deferrable initially deferred,
    email               varchar not null,
    phone_number        varchar,
    last_name           varchar not null,
    age                 integer,
    pronoun_id          integer
        constraint users_pronouns_id_fk
            references pronouns
            deferrable initially deferred,
    first_name          varchar not null,
    role                varchar not null,
    oauth_uid           varchar not null
        constraint users_oauth_uid_unique
            unique,
    oauth_provider      varchar not null,
    years_of_experience double precision,
    shirt_size          varchar not null,
    race                character varying[],
    gender              varchar
);

create unique index users_email_uindex
    on users (email);

create unique index users_phone_number_uindex
    on users (phone_number);

create table hackathon_applications
(
    id                        serial
        constraint hackathon_applications_pk
            primary key,
    user_id                   integer                 not null
        constraint hackathon_applications_users_null_fk
            references users,
    hackathon_id              integer                 not null
        constraint hackathon_applications_hackathons_null_fk
            references hackathons,
    why_attend                character varying[]     not null,
    what_do_you_want_to_learn character varying[]     not null,
    share_info_with_sponsors  boolean                 not null,
    application_status        varchar                 not null,
    created_time              timestamp default now() not null,
    status_change_time        timestamp
);

create table mailing_addresses
(
    user_id       integer             not null
        constraint mailing_addresses_pk
            primary key
        constraint mailing_addresses_users_null_fk
            references users,
    country       varchar             not null,
    state         varchar             not null,
    city          varchar             not null,
    postal_code   varchar             not null,
    address_lines character varying[] not null
);

alter table users
    add constraint users_mailing_addresses_user_id_fk
        foreign key (id) references mailing_addresses
            deferrable initially deferred;

create table mlh_terms
(
    user_id         integer not null
        constraint mlh_terms_pk
            primary key
        constraint mlh_terms_users_null_fk
            references users,
    send_messages   boolean not null,
    share_info      boolean not null,
    code_of_conduct boolean not null
);

alter table users
    add constraint users_mlh_terms_user_id_fk
        foreign key (id) references mlh_terms
            deferrable initially deferred;

create table education_info
(
    user_id         integer   not null
        constraint education_info_pk
            primary key
        constraint education_info_users_null_fk
            references users,
    name            varchar   not null,
    major           varchar   not null,
    graduation_date timestamp not null,
    level           varchar
);

alter table users
    add constraint users_education_info_user_id_fk
        foreign key (id) references education_info
            deferrable initially deferred;

create table event_attendance
(
    event_id integer                 not null
        constraint event_attendance_events_null_fk
            references events,
    user_id  integer                 not null
        constraint event_attendance_users_null_fk
            references users,
    time     timestamp default now() not null,
    constraint event_attendance_pk
        primary key (event_id, user_id)
);

create table meals
(
    hackathon_id integer             not null
        constraint meals_hackathons_null_fk
            references hackathons,
    user_id      integer             not null
        constraint meals_users_null_fk
            references users,
    meals        character varying[] not null,
    constraint meals_pk
        primary key (hackathon_id, user_id)
);

create table hackathon_checkin
(
    hackathon_id integer   not null
        constraint hackathon_attendance_hackathons_null_fk
            references hackathons,
    user_id      integer   not null
        constraint hackathon_attendance_users_null_fk
            references users,
    time         timestamp not null,
    constraint hackathon_checkin_pk
        primary key (hackathon_id, user_id)
);

create unique index api_keys_key_uindex
    on api_keys (key);


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
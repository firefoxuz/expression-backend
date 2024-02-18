create table public.expressions
(
    id            serial                              not null
        constraint expressions_pk
            primary key,
    expression    text                                not null,
    result        bigint,
    is_processing boolean   default false             not null,
    is_time_limit boolean   default null,
    is_valid      boolean   default null,
    is_finished   boolean   default false,
    time_limit    integer   default 200               not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null,
    finished_at   timestamp default null
);

comment
on column public.expressions.expression is 'expression to calculate';

comment
on column public.expressions.result is 'result of expression';

comment
on column public.expressions.is_processing is 'is calculations in processing';

comment
on column public.expressions.is_valid is 'is expression valid or not';

comment
on column public.expressions.time_limit is 'time limit to calculate an expression';

alter table public.expressions
    owner to expression_user;
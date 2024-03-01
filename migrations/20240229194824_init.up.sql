do
$$
begin
execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
end;
$$;


CREATE TABLE IF NOT EXISTS web_hooks
(
    id              bigserial NOT NULL CONSTRAINT webhooks_pk PRIMARY KEY,
    callback_url    text NOT NULL,
    http_method     text DEFAULT 'POST',
    params          jsonb,
    body            text,
    end_time        timestamp with time zone not null
)
CREATE TABLE IF NOT EXISTS reports
(
    id            BIGSERIAL                NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL,

    start_height  DECIMAL(65, 0)           NOT NULL,
    end_height    DECIMAL(65, 0)           NOT NULL,
    success_count INT,
    error_count   INT,
    error_msg     TEXT,
    duration      BIGINT,
    completed_at  TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY (id)
);

-- Indexes

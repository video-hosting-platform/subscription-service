BEGIN;

CREATE TABLE users (
    id INT PRIMARY KEY,
    subscription_count INT NOT NULL
);

CREATE TABLE subscribes (
    user_id INT references users(id), 
    subscriber_id INT references users(id),
    subscribed_at timestamp not null,
    PRIMARY KEY(user_id, subscriber_id)
);

COMMIT;
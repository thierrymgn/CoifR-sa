CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    user_type TEXT CHECK(user_type IN('admin', 'client', 'salon_owner')) NOT NULL
);

CREATE TABLE IF NOT EXISTS salons (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    description TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS hairdressers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    salon_id INTEGER NOT NULL,
    FOREIGN KEY(salon_id) REFERENCES salons(id)
);

CREATE TABLE IF NOT EXISTS slots (
    id SERIAL PRIMARY KEY,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    hairdresser_id INTEGER NOT NULL,
    FOREIGN KEY(hairdresser_id) REFERENCES hairdressers(id)
);

CREATE TABLE IF NOT EXISTS reservations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    slot_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(slot_id) REFERENCES slots(id)
);

CREATE INDEX IF NOT EXISTS idx_hairdressers_salon_id ON hairdressers(salon_id);
CREATE INDEX IF NOT EXISTS idx_slots_hairdresser_id ON slots(hairdresser_id);
CREATE INDEX IF NOT EXISTS idx_salons_user_id ON salons(user_id);
CREATE INDEX IF NOT EXISTS idx_reservations_slot_id ON reservations(slot_id);
CREATE INDEX IF NOT EXISTS idx_reservations_user_id ON reservations(user_id);
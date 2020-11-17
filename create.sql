CREATE TABLE ra(
    id INTEGER PRIMARY KEY,
    ra_name TEXT,
    ra_career TEXT
);

CREATE TABLE bed(
    id INTEGER PRIMARY KEY,
    room_num INTEGER,
    bed_name TEXT,
    ra1_id INTEGER,
    ra2_id INTEGER,
    bed_checked_out INTEGER,
    bed_time_checked_out INTEGER,
    FOREIGN KEY(ra1_id) REFERENCES ra(id),
    FOREIGN KEY(ra2_id) REFERENCES ra(id)
);

-- CREATE TABLE room (
--     room_num INTEGER PRIMARY KEY,
--     FOREIGN KEY(left_id) REFERENCES bed(id),
--     FOREIGN KEY(right_id) REFERENCES bed(id)
-- );
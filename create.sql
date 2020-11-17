CREATE TABLE ra(
    id INTEGER PRIMARY KEY,
    ra_name TEXT,
    ra_career TEXT
);

CREATE TABLE bed(
    id INTEGER PRIMARY KEY,
    room_num INTEGER,
    room_wing TEXT,
    bed_name TEXT,
    ra1_id INTEGER DEFAULT  0,
    ra2_id INTEGER DEFAULT  0,
    bed_checked_out INTEGER DEFAULT 0,
    bed_time_checked_out INTEGER DEFAULT -1,
    comments TEXT DEFAULT '',
    FOREIGN KEY(ra1_id) REFERENCES ra(id),
    FOREIGN KEY(ra2_id) REFERENCES ra(id)
);

-- CREATE TABLE room (
--     room_num INTEGER PRIMARY KEY,
--     FOREIGN KEY(left_id) REFERENCES bed(id),
--     FOREIGN KEY(right_id) REFERENCES bed(id)
-- );
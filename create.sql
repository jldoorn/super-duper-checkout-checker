CREATE TABLE ra(
    id INTEGER PRIMARY KEY,
    ra_name TEXT,
    ra_career TEXT
)

CREATE TABLE bed(
    id INTEGER PRIMARY KEY,
    FOREIGN KEY(ra_id) REFERENCES ra(id),
    bed_checked_out INTEGER,
    bed_time_checked_out INTEGER
);

CREATE TABLE room (
    room_num INTEGER PRIMARY KEY,
    FOREIGN KEY(left_id) REFERENCES bed(id),
    FOREIGN KEY(right_id) REFERENCES bed(id)
);
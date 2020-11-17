import sqlite3

if __name__ == '__main__':
    db = sqlite3.connect("beds.db")
    beds = 0
    for room in range(10):
        for bed in range(2):
            if bed == 0:
                name = "left"
            else:
                name = "right"
            sql = "INSERT INTO bed(id, room_num, room_wing, bed_name) VALUES(?,?,?,?)"
            db.execute(sql, (beds, room, "ne2", name))
            db.commit()
            beds += 1
    db.close()
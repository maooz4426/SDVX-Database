CREATE TABLE levels (
    level_id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    music_id INTEGER NOT NULL ,
    level_name VARCHAR(255) NOT NULL ,
    level_value INTEGER NOT NULL,
    created_At TIMESTAMP NOT NULL,
    updated_At TIMESTAMP NOT NULL,
    FOREIGN KEY (music_id) REFERENCES musics(music_id)
)
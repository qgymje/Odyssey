#!bin/sh

mysqldump -uodyssey -p odyssey --no-data | sed 's/ AUTO_INCREMENT=[0-9]*//g' > odyssey.sql

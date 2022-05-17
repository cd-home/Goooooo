echo $1 $2 $3 $4 $5
/usr/local/mysql-8.0.28-macos11-x86_64/bin/mysql -h$1 -P$2 -u$3 -p$4  $5 <<EOF 
    source ./internal/scripts/database.sql;
    SELECT * FROM casbin_rule limit 10;
    commit;
EOF 
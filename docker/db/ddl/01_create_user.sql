CREATE USER 'balance_user'@'%' IDENTIFIED BY 'balance_pass';
GRANT SELECT,INSERT,UPDATE,DELETE,EXECUTE,SHOW VIEW ON diary.* TO 'balance_user'@'%';

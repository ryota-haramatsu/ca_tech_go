FROM mysql:5.7
EXPOSE 3306

#MySQLログイン情報
ENV MYSQL_ROOT_USER root
ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE ca_tech
ENV MYSQL_USER ca_tech
ENV MYSQL_PASSWORD ca_tech

#MySQL設定ファイルをイメージ内にコピー
# ADD ./my.cnf /etc/mysql/conf.d/my.cnf

#docker runに実行される
# CMD ["mysqld"]
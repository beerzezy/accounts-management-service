FROM postgres:15.2

COPY /_build/script/pg/*.sql /docker-entrypoint-initdb.d/

# Set host time zone 
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
FROM golang:1.18

# buat folder app
RUN mkdir /app

# set direktori utama
WORKDIR /app

# copy seluruh file ke app
ADD . .

# SET TIME ZONE
RUN echo "Asia/Jakarta" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

# buat exe
RUN go build -o main

# run exe
CMD [ "./main" ]

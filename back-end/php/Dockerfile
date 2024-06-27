FROM php:7-alpine

RUN   apk update \                                                                                                                                                                                                                        
  &&   apk add ca-certificates wget \                                                                                                                                                                                                      
  &&   update-ca-certificates
RUN docker-php-ext-install mysqli && \
    docker-php-ext-install pdo_mysql
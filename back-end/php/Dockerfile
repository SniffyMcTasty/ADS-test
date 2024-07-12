FROM php:7-alpine

RUN   apk update \
  &&   apk add ca-certificates wget \
  &&   update-ca-certificates \

RUN docker-php-ext-install mysqli && \
    docker-php-ext-install pdo_mysql

ENV APP_PATH=/var/www

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
COPY . .


# Install composer
COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

ENV COMPOSER_ALLOW_SUPERUSER=1

# Install PHP library
RUN composer install

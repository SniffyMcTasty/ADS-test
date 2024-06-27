# -*- coding: utf-8 -*-

import os
from dotenv import load_dotenv

basedir = os.path.abspath(os.path.dirname(__file__))
load_dotenv(os.path.join(basedir, '.env'))


class Config(object):
    """
    Common configurations
    """

    DEBUG = eval(os.environ.get('DEBUG'))

    # sql alchemy config
    SQLALCHEMY_DATABASE_URI = os.environ.get('SQLALCHEMY_DATABASE_URI')
    # SQLALCHEMY_BINDS = None
    # SQLALCHEMY_NATIVE_UNICODE = None
    SQLALCHEMY_ECHO = eval(os.environ.get('SQLALCHEMY_ECHO'))
    SQLALCHEMY_RECORD_QUERIES = eval(os.environ.get('SQLALCHEMY_RECORD_QUERIES'))

    @staticmethod
    def init_app(app):
        pass

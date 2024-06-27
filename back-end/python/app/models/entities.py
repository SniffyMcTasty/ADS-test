# -*- coding: utf-8 -*-
"""
    app.models.entities
    ~~~~~~~~~~~~~~

    All the entity of the database

    :copyright: (c) 2023 by Automotive Data Solution.
"""

from app import db


class VehicleMake(db.Model):
    __tablename__ = 'vehicle_make'
    __table_args__ = {'extend_existing': True}

    vehicle_make_id = db.Column('vehicle_make_id', db.Integer, primary_key=True, autoincrement=True)
    name = db.Column('name', db.String(255), nullable=False)
    url = db.Column('url', db.String(255), nullable=False)

# -*- coding: utf-8 -*-
"""
    app.repositories.vehicle
    ~~~~~~~~~~~~~~

    Repository related to the vehicle make

    :copyright: (c) 2023 by Automotive Data Solution.
"""
from app.models.entities import VehicleMake


class VehicleMakeMapper:

    def find_all(self):
        q = VehicleMake.query

        objs = q.all()

        if len(objs) == 0:
            return []

        return objs

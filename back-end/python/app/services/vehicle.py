# -*- coding: utf-8 -*-
"""
    app.services.vehicle
    ~~~~~~~~~~~~~~

    Service for the vehicle

    :copyright: (c) 2023 by Automotive Data Solution.
"""

from app.repositories.vehicle import VehicleMakeMapper


class VehicleService:

    def __init__(self):
        pass

    @staticmethod
    def find_all_vehicle_make():
        vehicle_make_mapper = VehicleMakeMapper()

        makes = vehicle_make_mapper.find_all()

        result = []
        for make in makes:
            dto = dict(
                id=make.vehicle_make_id,
                name=make.name,
                url=make.url
            )

            result.append(dto)

        return result

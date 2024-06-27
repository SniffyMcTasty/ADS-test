# -*- coding: utf-8 -*-
"""
    app.api.vehicle
    ~~~~~~~~~~~~~~

    Endpoint for the vehicle

    :copyright: (c) 2023 by Automotive Data Solution.
"""

from . import api
from flask import jsonify
from app.services.vehicle import VehicleService


@api.route('/vehicle-makes', methods=['GET'])
def vehicle_make_list():
    lst_vehicle_make = VehicleService.find_all_vehicle_make()

    res = dict(vehicle_makes=lst_vehicle_make)

    return jsonify(res)

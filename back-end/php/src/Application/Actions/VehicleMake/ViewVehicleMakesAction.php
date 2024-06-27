<?php
declare(strict_types=1);

namespace App\Application\Actions\VehicleMake;

use Psr\Http\Message\ResponseInterface as Response;

class ViewVehicleMakesAction extends VehicleMakeAction
{
    /**
     * {@inheritdoc}
     */
    protected function action(): Response
    {
        $vehicleMakeId = (int) $this->resolveArg('id');
        $vehicleMake = $this->vehicleMakeRepository->findOne($vehicleMakeId);

        $this->logger->info("Vehicle make of id `${vehicleMakeId}` was viewed.");

        return $this->respondWithJSON($vehicleMake);
    }
}

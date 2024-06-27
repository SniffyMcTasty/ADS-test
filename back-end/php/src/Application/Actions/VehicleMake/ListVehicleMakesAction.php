<?php
declare(strict_types=1);

namespace App\Application\Actions\VehicleMake;

use Psr\Http\Message\ResponseInterface as Response;

class ListVehicleMakesAction extends VehicleMakeAction
{
    /**
     * {@inheritdoc}
     */
    protected function action(): Response
    {
        $vehicleMakes = $this->vehicleMakeRepository->findAll();

        $this->logger->info("Vehicle makes list was viewed.");

        return $this->respondWithJSON(['vehicle-makes' => $vehicleMakes]);
    }
}

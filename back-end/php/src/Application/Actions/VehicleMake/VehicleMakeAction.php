<?php
declare(strict_types=1);

namespace App\Application\Actions\VehicleMake;

use App\Application\Actions\Action;
use App\Domain\VehicleMake\VehicleMakeRepository;
use Psr\Log\LoggerInterface;

abstract class VehicleMakeAction extends Action
{
    /**
     * @var VehicleMakeRepository
     */
    protected $vehicleMakeRepository;

    /**
     * @param LoggerInterface $logger
     * @param VehicleMakeRepository $vehicleMakeRepository
     */
    public function __construct(LoggerInterface $logger,
                                VehicleMakeRepository $vehicleMakeRepository
    ) {
        parent::__construct($logger);
        $this->vehicleMakeRepository = $vehicleMakeRepository;
    }
}

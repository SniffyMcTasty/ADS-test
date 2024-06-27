<?php
declare(strict_types=1);

namespace App\Domain\VehicleMake;

use App\Domain\DomainException\DomainRecordNotFoundException;

class VehicleMakeNotFoundException extends DomainRecordNotFoundException
{
    public $message = 'The vehicle make you requested does not exist.';
}

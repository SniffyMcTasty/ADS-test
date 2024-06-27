<?php
declare(strict_types=1);

namespace App\Domain\VehicleMake;

use App\Domain\DomainException\DomainRecordNotFoundException;

class VehicleMakeListNotFoundException extends DomainRecordNotFoundException
{
    public $message = 'The vehicle make list you requested does not exist.';
}

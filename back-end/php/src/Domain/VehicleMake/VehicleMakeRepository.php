<?php

namespace App\Domain\VehicleMake;

use PDO;

/**
 * Repository.
 */
class VehicleMakeRepository
{
    /**
     * @var PDO The database connection
     */
    private $connection;

    /**
     * Constructor.
     *
     * @param PDO $connection The database connection
     */
    public function __construct(PDO $connection)
    {
        $this->connection = $connection;
    }

    /**
     * Convert an associative array to a VehicleMake object
     *
     * @param array $data
     *
     * @return VehicleMake
     */
    public function convertArrayToObject(array $data): VehicleMake
    {
        $vehicleMake = new VehicleMake();
        $vehicleMake->loadFromData($data);

        return $vehicleMake;
    }

    /**
     * Get all vehicle make rows.
     *
     * @return array 
     */
    public function findAll(): array
    {
        $sql = "SELECT * FROM vehicle_make;";

        $stmt = $this->connection->prepare($sql);
        $stmt->execute();
        $records = $stmt->fetchAll();

        if (!count($records)) {
            throw new VehicleMakeListNotFoundException();
        }

        return array_map(function($record) {
            return $this->convertArrayToObject($record)->jsonSerialize();
        }, $records);
    }

    /**
     * Find one row.
     *
     * @param int $vehicleMakeId The vehicle make ID
     *
     * @return int The new ID
     */
    public function findOne(int $vehicleMakeId): array
    {
        $row = ['vehicle_make_id' => $vehicleMakeId];

        $sql = "SELECT * 
                FROM vehicle_make 
                WHERE 
                    vehicle_make_id=:vehicle_make_id;";

        $stmt = $this->connection->prepare($sql);
        $stmt->execute($row);

        $record = $stmt->fetch();

        if (!$record) {
            throw new VehicleMakeNotFoundException();
        }

        return $this->convertArrayToObject($record)->jsonSerialize();
    }
}
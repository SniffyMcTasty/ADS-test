<?php
declare(strict_types=1);

namespace Tests\Domain\VehicleMake;

use App\Domain\VehicleMake\VehicleMake;
use Tests\TestCase;

class VehicleMakeTest extends TestCase
{
    public function vehicleMakeProvider()
    {
        return [
            [1, 'Mercedes', 'mercedes.com'],
            [2, 'Renault', 'renault.fr'],
            [3, 'Mini', 'mini.de'],
        ];
    }

    /**
     * @dataProvider vehicleMakeProvider
     * @param int    $id
     * @param string $name
     * @param string $url
     */
    public function testGetters(int $id, string $name, string $url)
    {
        $vehicleMake = new VehicleMake();
        $vehicleMake->setId($id);
        $vehicleMake->setName($name);
        $vehicleMake->setUrl($url);

        $this->assertEquals($id, $vehicleMake->getId());
        $this->assertEquals($name, $vehicleMake->getName());
        $this->assertEquals($url, $vehicleMake->getUrl());
    }

    /**
     * @dataProvider vehicleMakeProvider
     * @param int    $id
     * @param string $name
     * @param string $url
     */
    public function testJsonSerialize(int $id, string $name, string $url)
    {
        $vehicleMake = new VehicleMake();
        $vehicleMake->setId($id);
        $vehicleMake->setName($name);
        $vehicleMake->setUrl($url);

        $expectedPayload = json_encode([
            'id' => $id,
            'name' => $name,
            'url' => $url,
        ]);

        $this->assertEquals($expectedPayload, json_encode($vehicleMake));
    }
}

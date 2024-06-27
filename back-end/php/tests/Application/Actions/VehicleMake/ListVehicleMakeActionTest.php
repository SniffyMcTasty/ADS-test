<?php
declare(strict_types=1);

namespace Tests\Application\Actions\VehicleMake;

use App\Application\Actions\ActionPayload;
use App\Domain\VehicleMake\VehicleMakeRepository;
use App\Domain\VehicleMake\VehicleMake;
use DI\Container;
use Tests\TestCase;

class ListVehicleMakeActionTest extends TestCase
{
    public function testAction()
    {
        $app = $this->getAppInstance();

        /** @var Container $container */
        $container = $app->getContainer();

        $vehicleMake = new VehicleMake();

        $vehicleMakeRepositoryProphecy = $this->prophesize(VehicleMakeRepository::class);
        $vehicleMakeRepositoryProphecy
            ->findAll()
            ->willReturn([$vehicleMake])
            ->shouldBeCalledOnce();

        $container->set(VehicleMakeRepository::class, $vehicleMakeRepositoryProphecy->reveal());

        $request = $this->createRequest('GET', '/api/vehicle-makes');
        $response = $app->handle($request);

        $payload = (string) $response->getBody();
        $expectedPayload = ['vehicle-makes' => [$vehicleMake]];
        $serializedPayload = json_encode($expectedPayload, JSON_PRETTY_PRINT);

        $this->assertEquals($serializedPayload, $payload);
    }
}

<?php
declare(strict_types=1);

namespace Tests\Application\Actions\VehicleMake;

use App\Application\Actions\ActionError;
use App\Application\Actions\ActionPayload;
use App\Application\Handlers\HttpErrorHandler;
use App\Domain\VehicleMake\VehicleMake;
use App\Domain\VehicleMake\VehicleMakeNotFoundException;
use App\Domain\VehicleMake\VehicleMakeRepository;
use DI\Container;
use Slim\Middleware\ErrorMiddleware;
use Tests\TestCase;

class ViewVehicleMakeActionTest extends TestCase
{
    public function testAction()
    {
        $app = $this->getAppInstance();

        /** @var Container $container */
        $container = $app->getContainer();

        $vehicleMake = new VehicleMake();

        $vehicleMakeRepositoryProphecy = $this->prophesize(VehicleMakeRepository::class);
        $vehicleMakeRepositoryProphecy
            ->findOne(1)
            ->willReturn([$vehicleMake])
            ->shouldBeCalledOnce();

        $container->set(VehicleMakeRepository::class, $vehicleMakeRepositoryProphecy->reveal());

        $request = $this->createRequest('GET', '/api/vehicle-makes/1');
        $response = $app->handle($request);

        $payload = (string) $response->getBody();
        $expectedPayload = [$vehicleMake];
        $serializedPayload = json_encode($expectedPayload, JSON_PRETTY_PRINT);

        $this->assertEquals($serializedPayload, $payload);
    }

    public function testActionThrowsUserNotFoundException()
    {
        $app = $this->getAppInstance();

        $callableResolver = $app->getCallableResolver();
        $responseFactory = $app->getResponseFactory();

        $errorHandler = new HttpErrorHandler($callableResolver, $responseFactory);
        $errorMiddleware = new ErrorMiddleware($callableResolver, $responseFactory, true, false, false);
        $errorMiddleware->setDefaultErrorHandler($errorHandler);

        $app->add($errorMiddleware);

        /** @var Container $container */
        $container = $app->getContainer();

        $vehicleMakeRepositoryProphecy = $this->prophesize(VehicleMakeRepository::class);
        $vehicleMakeRepositoryProphecy
            ->findOne(1)
            ->willThrow(new VehicleMakeNotFoundException())
            ->shouldBeCalledOnce();

        $container->set(VehicleMakeRepository::class, $vehicleMakeRepositoryProphecy->reveal());

        $request = $this->createRequest('GET', '/api/vehicle-makes/1');
        $response = $app->handle($request);

        $payload = (string) $response->getBody();
        $expectedError = new ActionError(ActionError::RESOURCE_NOT_FOUND, 'The vehicle make you requested does not exist.');
        $expectedPayload = new ActionPayload(404, null, $expectedError);
        $serializedPayload = json_encode($expectedPayload, JSON_PRETTY_PRINT);

        $this->assertEquals($serializedPayload, $payload);
    }
}

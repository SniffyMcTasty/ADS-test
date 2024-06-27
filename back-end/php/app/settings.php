<?php
declare(strict_types=1);

use App\Application\Settings\Settings;
use App\Application\Settings\SettingsInterface;
use DI\ContainerBuilder;
use Monolog\Logger;

return function (ContainerBuilder $containerBuilder) {
    $rootPath = realpath(__DIR__ . '/..');

    $configIni = parse_ini_file($rootPath . "/config/application.ini", true);

    // Global Settings Object
    $containerBuilder->addDefinitions([
        SettingsInterface::class => function () use ($rootPath, $configIni) {
            return new Settings([
                'displayErrorDetails' => true, // Should be set to false in production
                'logger' => [
                    'name' => 'backend_phptest',
                    'path' => isset($_ENV['docker']) ? 'php://stdout' : __DIR__ . '/../logs/app.log',
                    'level' => Logger::DEBUG,
                ],
                'db' => [
                    'driver' => $configIni['database']['driver'],
                    'host' => $configIni['database']['host'],
                    'database' => $configIni['database']['database'],
                    'username' => $configIni['database']['username'],
                    'password' => $configIni['database']['password'],
                    'charset' => $configIni['database']['charset'],
                    'collation' => $configIni['database']['collation'],
                    'flags' => [
                        // Turn off persistent connections
                        PDO::ATTR_PERSISTENT => false,
                        // Enable exceptions
                        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
                        // Emulate prepared statements
                        PDO::ATTR_EMULATE_PREPARES => true,
                        // Set default fetch mode to array
                        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
                        // Set character set
                        PDO::MYSQL_ATTR_INIT_COMMAND => 'SET NAMES utf8mb4 COLLATE utf8mb4_unicode_ci'
                    ]
                ],
            ]);
        }
    ]);
};

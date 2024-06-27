<?php
declare(strict_types=1);

namespace App\Domain\VehicleMake;

use JsonSerializable;

class VehicleMake implements JsonSerializable
{
    /**
     * @var int|null
     */
    private $id;

    /**
     * @var string
     */
    private $name;

    /**
     * @var string
     */
    private $url;

    /**
     * Constructor
     */
    public function __construct()
    {
        $this->id = 0;
        $this->name = '';
        $this->url = '';
    }

    /**
     * @return int|null
     */
    public function getId(): ?int
    {
        return $this->id;
    }

    /**
     * @param $id int
     */
    public function setId(int $id): void
    {
        $this->id = $id;
    }

    /**
     * @return string
     */
    public function getName(): string
    {
        return $this->name;
    }

    /**
     * @param $name string
     */
    public function setName(string $name): void
    {
        $this->name = $name;
    }

    /**
     * @return string
     */
    public function getUrl(): string
    {
        return $this->url;
    }

    /**
     * @param $url string
     */
    public function setUrl(string $url): void
    {
        $this->url = $url;
    }

    /**
     * @param $data array
     */
    public function loadFromData(array $data): void
    {
        if (array_key_exists('vehicle_make_id', $data) && $data['vehicle_make_id'] != null) {
            $this->setId((int)$data['vehicle_make_id']);
        }

        if (array_key_exists('name', $data) && $data['name'] != null) {
            $this->setName($data['name']);
        }

        if (array_key_exists('url', $data) && $data['url'] != null) {
            $this->setUrl($data['url']);
        }
    }

    /**
     * @return array
     */
    public function jsonSerialize()
    {
        return [
            'id' => $this->getId(),
            'name' => $this->getName(),
            'url' => $this->getUrl(),
        ];
    }
}

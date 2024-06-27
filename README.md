# Short test

## Requirements

* A single page "app" that displays years and vehicle models in a grid format coming from the api call `http://localhost:8080/api/vehicle-coverage`.
* A grid with the contents of `years` and `vehicle-models` (as in the mockup below).

![Database structure](./data/mockup.png "Page Mockup")

* Set the corresponding box to be blue (as in the mockup) if the entry for that vehicle model and year exists in the `coverage` segment of the json. And grey if it doesn't exist.
* When clicking on a vehicle-model/year box, it is toggled (set the state to 0 in the database if it was previously 1 or 0  if it was not...) and the visual displays the new state.

It will take at least 2 hours but no more than 4 to complete the test.

### Example

Once the list of vehicles are loaded into the grid

- The years for the RLX are `"coverage":{"RLX":[2012,2011,2010]}`
- After clicking on the 2013 column next to RLX: `"coverage":{"RLX":[2013,2012,2011,2010]}`
- The visual should now show a blue box for the year 2013 next to RLX.
- Clicking on the 2010 column next to RLX: 
`"coverage":{"RLX":[2013,2012,2011]}`
- The visual should now show a grey box for the year 2010 next to RLX.

Note that this does not change the json file, only the loaded object.

## Evaluation

This will be evaluated on the compliance to the supplied visual (mockup.png), and the code quality.

You will have the opportunity to justify your decisions during the interview.

## Installation

### Install Parcel

Parcel is required for this project to run properly.

```sh
npm install -g parcel-bundler
```

### Install dependencies

```sh
npm install
```

## Run dev environment

```sh
npm run dev
```

Navigate to http://localhost:1234

## Backend

* Create a call to get the list of vehicle model by vehicle make 

```bash
GET http://localhost:8080/api/vehicle-makes/:vehicleMakeId/vehicle-models
```

Response
```json
{
    "models": [
        {
            "id": 2,
            "name": "MDX"
        },
        {
            "id": 6,
            "name": "TL"
        }
    ]
}
```


* Create a call to get the list of vehicle years by vehicle make

```bash
GET http://localhost:8080/api/vehicle-makes/:vehicleMakeId/vehicle-years
```

Response
```json
{
    "years": [2018, 2017, 2016, 2015, 2014, 2013, 2012, 2011, 2010]
}
```


Create a call to get the vehicle coverage

```bash
GET http://localhost:8080/api/vehicle-makes/:vehicleMakeId/vehicle-coverage
```

Response
```json 
{
  "vehicle-models": ["ILX", "MDX", "RDX", "RLX", "TL", "TLX", "TSX"],
  "years": [2018, 2017, 2016, 2015, 2014, 2013, 2012, 2011, 2010],
  "coverage": [
      {
          "vehicle_id": 1,
          "vehicle_year": 2017,
          "vehicle_model_id": 
      },
      
    "ILX": [2017, 2016, 2015, 2014],
    "MDX": [2017, 2016, 2015, 2014],
    "RDX": [2011, 2010],
    "RLX": [2012, 2011, 2010],
    "TL": [2014, 2013, 2012, 2011, 2010],
    "TLX": [2016, 2015, 2014, 2013],
    "TSX": [2017, 2015]
  ]
}
```

* Create a call to update the state of the vehicle in the database












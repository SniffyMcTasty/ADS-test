# Short test

## Provided

A json file, containing vehicle data, and a png of the visual we want this component to use.

## Requirements

- A single page "app" that displays years and vehicle models in a grid format coming from the supplied json file.
- A grid with the contents of `years` and `vehicle-models` (as in the mockup.png).
- Set the corresponding box to be blue (as in the mockup.png) if the entry for that vehicle model and year exists in the `coverage` segment of the json. And grey if it doesn't exist.
- When clicking on a vehicle-model/year box, it is toggled (removes the entry from the javascript `coverage` object if it previously existed or adds it if it did not...) and the visual displays the new state.

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

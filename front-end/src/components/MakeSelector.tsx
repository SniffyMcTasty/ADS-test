import React from "react";
import type { VehicleMake, VehicleMakeData } from "../types/vehicle";

interface Props {
  makes?: VehicleMakeData;
  selected: VehicleMake | undefined;
  onChange: (make: VehicleMake) => void;
}

const MakeSelector: React.FC<Props> = ({
  makes,
  selected,
  onChange
}) => {
  return (
    <div className="make-selector">
      <select
        value={selected?.id || 0}
        onChange={(e) => {
          const selectedMake = makes?.vehicle_makes?.find(m => m.id === parseInt(e.target.value));
          if (selectedMake) onChange(selectedMake);
        }}
      >
        {
          Object.entries(makes?.vehicle_makes || []).map(([id, make]) => (
          <option key={id} value={make.id}>
            {make.name}
          </option>
        ))}
      </select>
    </div>
  );
};

export default MakeSelector;

import React from "react";
import type { VehicleMakeData } from "../types/vehicle";

interface Props {
  makes?: VehicleMakeData;
  selected: number;
  onChange: (make: number) => void;
}

const MakeSelector: React.FC<Props> = ({
  makes,
  selected,
  onChange
}) => {
  return (
    <div className="make-selector">
      <select
        value={selected}
        onChange={(e) => {
          onChange(e.target.value ? parseInt(e.target.value) : 0);
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

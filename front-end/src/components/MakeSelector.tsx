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
        onChange={(e) => onChange(parseInt(e.target.value))}
      >
        {Object.entries(makes || []).map(([id, name]) => (
          <option key={id} value={id}>
            {name}
          </option>
        ))}
      </select>
    </div>
  );
};

export default MakeSelector;

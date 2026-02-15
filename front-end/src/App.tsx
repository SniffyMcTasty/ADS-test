import React, { useEffect, useState } from "react";
import CoverageGrid from "./components/CoverageGrid";
import MakeSelector from "./components/MakeSelector";
import "./styles/App.scss";
import { getMakes, login } from "./api/api";
import type { VehicleMakeData } from "./types/vehicle";

const App: React.FC = () => {
  const [makes, setMakes] = useState<VehicleMakeData>();
  const [selectedMake, setSelectedMake] =useState<number>(0);

  useEffect(() => {
    login()
      .then(() => {
        console.log("Login successful");
        return getMakes();
      })
      .then((data) => {
        setMakes(data);
      })
      .catch((err) => {
        console.error("Error during initialization:", err);
      });
  }, []); // Empty dependency array = runs once on mount

  return (
    <div className="App">
      <MakeSelector
        makes={makes}
        selected={selectedMake}
        onChange={setSelectedMake}
      />

      <CoverageGrid selectedMake={selectedMake} />
    </div>
  );
};

export default App;

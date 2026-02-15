import React, { useEffect, useState } from "react";
import CoverageGrid from "./components/CoverageGrid";
import MakeSelector from "./components/MakeSelector";
import "./styles/App.scss";
import { getMakes, login } from "./api/api";
import type { VehicleMake, VehicleMakeData } from "./types/vehicle";
import Spinner from "./components/Spinner";

const App: React.FC = () => {
  const [makes, setMakes] = useState<VehicleMakeData>();
  const [selectedMake, setSelectedMake] =useState<VehicleMake>();
  const [ready, setReady] = useState(false);

  useEffect(() => {
    login()
      .then(() => {
        console.log("Login successful");
        return getMakes();
      })
      .then((data) => {
        setMakes(data);
        const firstMakeId = data.vehicle_makes?.[0]?.id || 0;
        if (firstMakeId) setSelectedMake(data.vehicle_makes?.[0]);
        setReady(true);
      })
      .catch((err) => {
        console.error("Error during initialization:", err);
      });
  }, []); // Empty dependency array = runs once on mount

  if (!ready) return <Spinner />;

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

import React, { useEffect, useState } from "react";
import {
  getModels,
  getYears,
  getCoverage,
  toggleCoverage
} from "../api/api";
import Spinner from "./Spinner";
import "../styles/CoverageGrid.scss";
import type { VehicleModelData, VehicleYearData } from "../types/vehicle";

interface Props {
  selectedMake: number;
}

const CoverageGrid: React.FC<Props> = ({
  selectedMake
}) => {
  const [models, setModels] = useState<VehicleModelData>();
  const [years, setYears] = useState<VehicleYearData>();
  const [coverage, setCoverage] = useState<
    Record<string, number[]>
  >({});
  const [loading, setLoading] = useState<boolean>(true);
  const [cellLoading, setCellLoading] =
    useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const fetchAll = async () => {
    setLoading(true);
    setError(null);

    try {
      const modelsData = await getModels(selectedMake);
      const yearsData = await getYears(selectedMake);
      const coverageData = await getCoverage(selectedMake);

      setModels(modelsData);
      setYears(yearsData);
      setCoverage(coverageData.coverage);
    } catch (err: any) {
      setError(err.message || "Something went wrong");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (!selectedMake) return; // Skip if no make selected
    fetchAll();
  }, [selectedMake]);

  const handleToggle = async (
    model: number,
    year: number
  ) => {
    const key = `${model}-${year}`;
    setCellLoading(key);
    setError(null);

    try {
      await toggleCoverage(selectedMake, model, year);

      const updated = await getCoverage(selectedMake);
      setCoverage(updated.coverage);
    } catch (err: any) {
      setError(err.message || "Toggle failed");
    } finally {
      setCellLoading(null);
    }
  };

  const isActive = (model: string, year: number) =>
    coverage[model]?.includes(year);

  if (loading) return <Spinner />;

  return (
    <div className="grid-container">
      <h2>{selectedMake}</h2>

      {error && <div className="error">{error}</div>}

      <table>
        <thead>
          <tr>
            <th>Model</th>
            {Object.values(years || []).map((year) => (
              <th key={year}>{year}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {Object.values(models || []).map((model) => (
            <tr key={model.id}>
              <td className="model-name">
                {model.name}
              </td>
              {Object.values(years || []).map((year) => {
                const key = `${model.id}-${year}`;
                return (
                  <td
                    key={year}
                    className={
                      isActive(model.id, year)
                        ? "active"
                        : "inactive"
                    }
                    onClick={() =>
                      !cellLoading &&
                      handleToggle(model.id, year)
                    }
                  >
                    {cellLoading === key
                      ? "..."
                      : ""}
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default CoverageGrid;

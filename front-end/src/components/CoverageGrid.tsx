import React, { useEffect, useState } from "react";
import {
  getModels,
  getYears,
  getCoverage,
  toggleCoverage
} from "../api/api";
import Spinner from "./Spinner";
import "../styles/CoverageGrid.scss";
import type { VehicleCoverageData, VehicleMake, VehicleModel, VehicleModelData, VehicleYearData } from "../types/vehicle";

interface Props {
  selectedMake: VehicleMake | undefined;
}

const CoverageGrid: React.FC<Props> = ({
  selectedMake
}) => {
  const [modelList, setModels] = useState<VehicleModelData>();
  const [yearList, setYears] = useState<VehicleYearData>();
  const [coverage, setCoverage] = useState<VehicleCoverageData>();
  const [loading, setLoading] = useState<boolean>(false);
  const [cellLoading, setCellLoading] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const fetchAll = async () => {
    setLoading(true);
    setError(null);

    try {
      const [modelsData, yearsData, coverageData] = await Promise.all([
        getModels(selectedMake?.id || 0),
        getYears(selectedMake?.id || 0),
        getCoverage(selectedMake?.id || 0),
      ]);

      setModels(modelsData);
      setYears(yearsData);
      setCoverage(coverageData);
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
    model: VehicleModel,
    year: number
  ) => {
    const key = `${model}-${year}`;
    setCellLoading(key);
    setError(null);

    try {
      await toggleCoverage(selectedMake?.id || 0, model.id, year);

      const updated = coverage?.coverage[model.name]?.includes(year) ? coverage.coverage[model.name].filter((y) => y !== year)
        : [...(coverage?.coverage[model.name] || []), year];
      setCoverage(prev => ({
        ...prev,
        coverage: {
          ...prev?.coverage,
          [model.name]: updated
        }
      }));
    } catch (err: any) {
      setError(err.message || "Toggle failed");
    } finally {
      setCellLoading(null);
    }
  };

  const isActive = (model: string, year: number) =>
    coverage?.coverage[model]?.includes(year) ?? false;

  if (loading) return <Spinner />;

  return (
    <div className="grid-wrapper">
      {error && <div className="grid-error">{error}</div>}

      <div className="grid-scroll">
        <table className="coverage-table">
          <thead>
            <tr>
              <th
                className="logo-cell"
                rowSpan={(Object.entries(modelList ?? {})?.length ?? 0) + 1}
              >
                {selectedMake?.name && (
                  <img
                    src={new URL(`../assets/logos/logo-${selectedMake.id}.png`, import.meta.url)}
                    alt={selectedMake.name}
                    className="make-logo"
                  />
                )}
              </th>

              {Object.values(yearList?.vehicle_years ?? {}).map((year) => (
                <th key={year} className="year-header">
                  <span>{year}</span>
                </th>
              ))}
            </tr>
          </thead>

          <tbody>
            {Object.entries(modelList?.vehicle_models ?? {}).map(([id, model], rowIndex) => (
              <tr key={id} className={rowIndex % 2 === 0 ? "row-even" : "row-odd"}>
                <td className="model-name">{model.name}</td>

                {Object.values(yearList?.vehicle_years ?? {}).map((year) => {
                  const key = `${model.id}-${year}`;
                  const active = isActive(model.name, year);
                  const isLoadingCell = cellLoading === key;

                  return (
                    <td
                      key={year}
                      className={`coverage-cell ${active ? "cell-active" : "cell-inactive"} ${isLoadingCell ? "cell-loading" : ""}`}
                      onClick={() => !cellLoading && handleToggle(model, year)}
                      title={`${model.name} – ${year}`}
                    >
                      {isLoadingCell && <span className="cell-spinner" />}
                    </td>
                  );
                })}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default CoverageGrid;

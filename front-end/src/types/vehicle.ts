// Core domain types
export interface VehicleMake {
  id: number;
  name: string;
  url: string;
}

export interface VehicleModel {
  id: number;
  name: string;
}

export interface Vehicle {
  id: number;
  make: number;
  model: string;
  year: number;
  state: number;
  updated: string;
}

// API Data Wrappers (mirror Go structs)
export interface VehicleMakeData {
  vehicle_makes: VehicleMake[];
}

export interface VehicleModelData {
  vehicle_models: VehicleModel[];
}

export interface VehicleYearData {
  vehicle_years: number[];
}

export interface VehicleCoverageData {
  coverage: Record<string, number[]>;
}

export interface PostMessageData {
  message: string;
}

export interface LoginData {
  token: string;
}

import type { APIResponse } from "../types/api";
import type {
  VehicleModelData,
  VehicleYearData,
  VehicleCoverageData,
  LoginData,
  VehicleMakeData
} from "../types/vehicle";

const API_BASE = import.meta.env.VITE_API_BASE;
const USERNAME = import.meta.env.VITE_API_USERNAME;
const PASSWORD = import.meta.env.VITE_API_PASSWORD;

let token: string | null = null;

async function apiFetch<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  const res = await fetch(`${API_BASE}${endpoint}`, {
    ...options,
  });

  if (!res.ok) {
    throw new Error(`Network error: ${res.status}`);
  }

  const json: APIResponse<T> = await res.json();

  if (!json.success) {
    const errorMsg =
      json.error?.message || "API error";
    throw new Error(errorMsg);
  }

  return json.data as T;
  
}

export const login = async () => {
  const response = await apiFetch<LoginData>(
    "/login", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify({
        username: USERNAME,
        password: PASSWORD
    })
  });

  token = response.token;
};

const authHeaders = () => ({
  "Content-Type": "application/json",
  Authorization: `Bearer ${token}`
});

export const getMakes = async () => {
    return apiFetch<VehicleMakeData>(
        "/vehicle-makes",
        {
            method: "GET",
            headers: authHeaders()
        }
    );
};

export const getModels = async (make: number) => {
    return apiFetch<VehicleModelData>(
        `/vehicle-makes/${make}/vehicle-models`,
        {
            method: "GET",
            headers: authHeaders()
        }
    );
};

export const getYears = async (make: number) => {
    return apiFetch<VehicleYearData>(
        `/vehicle-makes/${make}/vehicle-years`,
        {
            method: "GET",
            headers: authHeaders()
        }
    );
};

export const getCoverage = async (make: number) => {
    return apiFetch<VehicleCoverageData>(
        `/vehicle-makes/${make}/vehicle-coverage`,
        {
            method: "GET",
            headers: authHeaders()
        }
    );
};

export const toggleCoverage = async (
  make: number,
  model: number,
  year: number
) => {
    return apiFetch<VehicleCoverageData>(
        `/vehicle-makes/${make}/change-coverage/${model}/${year}`,
        {
            method: "POST",
            headers: authHeaders(),
            body: JSON.stringify({ make, model, year })
        }
    );
};

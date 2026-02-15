export interface APIErrorResponse {
  code: number;
  message: string;
}

export interface APIResponse<T> {
  success: boolean;
  data?: T;
  error?: APIErrorResponse;
}

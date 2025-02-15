export type ApiResponse<T> = {
  data: T;
  message: string;
  status: number;
};

export type ErrorBody = {
  code: string;
  message: string;
};

export type ErrorResponse = ApiResponse<ErrorBody>;

export type ApiResponseType<T> = ApiResponse<T> | ErrorResponse;

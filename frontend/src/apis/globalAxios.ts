import axios, { AxiosError } from "axios";

const BASE_API_URL = import.meta.env.VITE_REACT_APP_API_URL;

const apiClient = axios.create({
  baseURL: BASE_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export default apiClient;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const isAxiosError = (error: any): error is AxiosError =>
  !!error.isAxiosError;

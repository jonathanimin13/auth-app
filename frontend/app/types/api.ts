export interface Response<TData> {
  message?: string;
  data?: TData;
  errors?: FetchError[];
}

export interface FetchError {
  field: string;
  detail: string;
}

export interface User {
  username: string;
  access_token: string;
}

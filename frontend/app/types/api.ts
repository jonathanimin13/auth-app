export interface Response<TData> {
  message?: string;
  data?: TData;
  errors?: FetchError[];
}

export interface FetchError {
  field: string;
  detail: string;
}

export interface Token {
  access_token: string;
}

export interface User {
  email: string;
  username: string;
}

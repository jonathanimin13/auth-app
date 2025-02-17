import { Response } from "@/app/types/api";

export const fetchLogin = async (loginData: {
  email: string;
  password: string;
}) => {
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/auth/login`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
        credentials: "include",
      }
    );

    if (!response.ok) {
      const errorData: Response<undefined> = await response.json();

      const errorMessages = errorData.errors
        ?.map((error) => error.detail)
        .join(", ");

      throw new Error(errorMessages);
    }

    return response.json();
  } catch (error) {
    if (error instanceof Error && error.message !== "Failed to fetch") {
      throw error;
    }

    throw new Error("internal server error");
  }
};

export const fetchRegister = async (registerData: {
  username: string;
  email: string;
  password: string;
}) => {
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/auth/register`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(registerData),
        credentials: "include",
      }
    );

    if (!response.ok) {
      const errorData: Response<undefined> = await response.json();

      const errorMessages = errorData.errors
        ?.map((error) => error.detail)
        .join(", ");

      throw new Error(errorMessages);
    }

    return response.json();
  } catch (error) {
    if (error instanceof Error && error.message !== "Failed to fetch") {
      throw error;
    }

    throw new Error("internal server error");
  }
};

export const fetchLogout = async () => {
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/auth/logout`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      }
    );

    if (!response.ok) {
      const errorData: Response<undefined> = await response.json();

      const errorMessages = errorData.errors
        ?.map((error) => error.detail)
        .join(", ");

      throw new Error(errorMessages);
    }

    return response.json();
  } catch (error) {
    if (error instanceof Error && error.message !== "Failed to fetch") {
      throw error;
    }

    throw new Error("internal server error");
  }
};

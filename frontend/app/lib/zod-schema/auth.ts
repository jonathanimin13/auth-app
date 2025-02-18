import { z } from "zod";

export const loginSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(1, "Password is required"),
});

const passwordSchema = z.string().refine(
  (password) => {
    if (password.length < 8) return false;

    let hasUpper = false;
    let hasLower = false;
    let hasNumber = false;
    let hasSpecial = false;
    let hasWhiteSpace = false;

    for (let char of password) {
      if (/[A-Z]/.test(char)) hasUpper = true;
      if (/[a-z]/.test(char)) hasLower = true;
      if (/\d/.test(char)) hasNumber = true;
      if (/[!@#$%^&*(),.?":{}|<>]/.test(char)) hasSpecial = true;
      if (/\s/.test(char)) hasWhiteSpace = true;
    }

    return hasUpper && hasLower && hasNumber && hasSpecial && !hasWhiteSpace;
  },
  {
    message:
      "Password must be at least 8 characters long, contain an uppercase letter, a lowercase letter, a number, a special character, and cannot contain whitespace.",
  }
);

export const registerSchema = z.object({
  username: z.string().min(1, "Username is required"),
  email: z.string().email("Invalid email address"),
  password: passwordSchema,
});

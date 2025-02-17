"use client";

import { loginSchema } from "@/app/lib/zod-schema/auth";
import { Button } from "@/app/ui/button/button";
import { ErrorCard } from "@/app/ui/error-card/error-card";
import { Input } from "@/app/ui/Input/input";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { ZodFormattedError } from "zod";

export const Login = () => {
  const [validationErrors, setValidationErrors] = useState<
    ZodFormattedError<
      {
        email: string;
        password: string;
      },
      string
    >
  >();
  const [authErrors, setAuthErrors] = useState<string>();
  const [showPassword, setShowPassword] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const router = useRouter();

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    setIsLoading(true);
    setAuthErrors(undefined);
    setValidationErrors(undefined);

    const formData = new FormData(e.target as HTMLFormElement);

    const validationResult = loginSchema.safeParse({
      email: formData.get("email"),
      password: formData.get("password"),
    });

    if (validationResult.error) {
      const formatedError = validationResult.error.format();
      setValidationErrors(formatedError);
      setIsLoading(false);
      return;
    }

    const result = await signIn("credentials", {
      redirect: false,
      email: formData.get("email"),
      password: formData.get("password"),
    });

    if (result?.error) {
      setAuthErrors(result.error);
      setIsLoading(false);

      return;
    }

    router.push("/");
  };

  return (
    
      <div className="p-6 md:p-16 flex flex-col gap-4">
        <h1 className="text-3xl font-bold text-center">Login</h1>
        <form className="flex flex-col gap-4" onSubmit={handleLogin}>
          {authErrors && <ErrorCard errors={authErrors} />}
          <div className="flex flex-col gap-4">
            <div>
              <Input
                type="text"
                name="email"
                placeholder="example@mail.com"
                label="Email"
                required={true}
              />
              {validationErrors?.email &&
                validationErrors.email._errors.map((error, idx) => {
                  return (
                    <p key={idx} className="text-red-600 pt-1 block">
                      {error}
                    </p>
                  );
                })}
            </div>
            <div>
              <Input
                type={showPassword ? "text" : "password"}
                name="password"
                placeholder="Password"
                label="Password"
                required={true}
                button={
                  <p className="text-blue-500 font-semibold">
                    {showPassword ? "Hide" : "Show"}
                  </p>
                }
                onButtonClick={() => setShowPassword((prev) => !prev)}
              />
              {validationErrors?.password &&
                validationErrors.password._errors.map((error, idx) => {
                  return (
                    <p key={idx} className="text-red-600 pt-1 block">
                      {error}
                    </p>
                  );
                })}
            </div>
          </div>
          <Button disabled={isLoading}>
            {isLoading ? "Loading..." : "Login"}
          </Button>
        </form>
        <p onClick={() => router.push("/")}>Click here to register!</p>
      </div>
  );
};

export default Login;

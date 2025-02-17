"use client";

import { fetchLogin } from "@/app/lib/actions/auth";
import { loginSchema } from "@/app/lib/zod-schema/auth";
import { Button } from "@/app/ui/button/button";
import { ErrorCard } from "@/app/ui/error-card/error-card";
import { Input } from "@/app/ui/Input/input";
import { useMutation } from "@tanstack/react-query";

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
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const router = useRouter();

  const { mutate, isPending, isError, error } = useMutation({
    mutationFn: fetchLogin,
    onSuccess: () => {
      router.push("/");
    },
  });

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setValidationErrors(undefined);

    const formData = new FormData(e.target as HTMLFormElement);

    const validationResult = loginSchema.safeParse({
      email: formData.get("email"),
      password: formData.get("password"),
    });

    if (validationResult.error) {
      const formatedError = validationResult.error.format();
      setValidationErrors(formatedError);
      return;
    }

    mutate({
      email: formData.get("email") as string,
      password: formData.get("password") as string,
    });
  };

  return (
    <div className="p-6 md:p-16 flex flex-col gap-4">
      <h1 className="text-3xl font-bold text-center">Login</h1>
      <form className="flex flex-col gap-4" onSubmit={handleLogin}>
        <div className="flex flex-col gap-4">
          {isError && <ErrorCard errors={error.message} />}
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
        <Button disabled={isPending}>
          {isPending ? "Loading..." : "Login"}
        </Button>
      </form>
      <p onClick={() => router.push("/")}>Click here to register!</p>
    </div>
  );
};

export default Login;

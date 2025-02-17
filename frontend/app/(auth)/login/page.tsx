"use client";

import { handleLogin } from "@/app/lib/actions/login";
import { Button } from "@/app/ui/button/button";
import { Input } from "@/app/ui/Input/input";
import { useActionState, useState } from "react";

export const Login = () => {
  const [state, login, isPending] = useActionState(handleLogin, null);
  const [showPassword, setShowPassword] = useState<boolean>(false);

  return (
    <div className="p-6 md:p-16 flex flex-col gap-4">
      <h1 className="text-3xl font-bold text-center">Login</h1>
      <form className="flex flex-col gap-4" action={login}>
        <div className="flex flex-col gap-4">
          <div>
            <Input
              type="text"
              name="email"
              placeholder="example@mail.com"
              label="Email"
              required={true}
            />
            {state?.validationerrors.email &&
              state.validationerrors.email.map((error, idx) => {
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
            {state?.validationerrors.password &&
              state.validationerrors.password.map((error, idx) => {
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
    </div>
  );
};

export default Login;

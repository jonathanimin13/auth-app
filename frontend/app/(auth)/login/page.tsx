"use client";

import { Button } from "@/app/ui/button/button";
import { Input } from "@/app/ui/Input/input";
import { useState } from "react";

export const Login = () => {
  const [showPassword, setShowPassword] = useState<boolean>(false);

  return (
    <div className="p-6 md:p-16 flex flex-col gap-4">
      <h1 className="text-3xl font-bold text-center">Login</h1>
      <form className="flex flex-col gap-4">
        <div className="flex flex-col gap-4">
          <Input
            type="text"
            name="email"
            placeholder="example@mail.com"
            label="Email"
            required={true}
          />
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
        </div>
        <Button>Login</Button>
      </form>
    </div>
  );
};

export default Login;

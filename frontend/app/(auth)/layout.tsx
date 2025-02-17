"use client";

import { useRouter } from "next/navigation";

export const AuthLayout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const router = useRouter();
  if (localStorage.getItem("access-token")) {
    router.push("/");
  }

  return (
    <div className="w-full h-screen flex justify-center items-center">
      <div className="w-[80%] md:w-[610px] mx-auto border border-black rounded-lg">
        {children}
      </div>
    </div>
  );
};

export default AuthLayout;

"use client";

import { useMutation } from "@tanstack/react-query";
import Button from "../button/button";
import { fetchLogout } from "@/app/lib/actions/auth";
import { useRouter } from "next/navigation";
import toast, { Toaster } from "react-hot-toast";

export const Navbar = () => {
  const router = useRouter();

  const { mutate, isPending, isError, error } = useMutation({
    mutationFn: fetchLogout,
    onSuccess: () => {
      router.push("/login");
    },
    onError: (error) => {
      toast.error(error.message);
    },
  });

  const handleLogout = async () => {
    mutate();
  };

  return (
    <div className="border-b">
      <div className="w-[80%] p-4 mx-auto flex justify-between items-center">
        <p className="font-semibold text-3xl">Auth.app</p>
        <div className="w-fit">
          <Button onClick={handleLogout} disabled={isPending}>
            {isPending ? "loading..." : "Logout"}
          </Button>
        </div>
        <Toaster />
      </div>
    </div>
  );
};

export default Navbar;

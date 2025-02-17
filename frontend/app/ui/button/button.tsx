import { ReactNode } from "react";

interface buttonProps {
  children: ReactNode;
}

export const Button = ({ children }: buttonProps) => {
  return (
    <button className="w-full bg-blue-600 text-white font-semibold p-3 rounded-lg">
      {children}
    </button>
  );
};

export default Button;

import { ReactNode } from "react";

interface buttonProps {
  children: ReactNode;
  disabled?: boolean;
}

export const Button = ({ children, disabled = false }: buttonProps) => {
  return (
    <button
      className="w-full bg-blue-600 text-white font-semibold p-3 rounded-lg disabled:bg-gray-400 disabled:text-gray-600 disabled:cursor-not-allowed"
      disabled={disabled}
    >
      {children}
    </button>
  );
};

export default Button;

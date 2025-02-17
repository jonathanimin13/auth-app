import { ReactNode } from "react";

interface buttonProps {
  children: ReactNode;
  disabled?: boolean;
  onClick?: () => void;
}

export const Button = ({
  children,
  disabled = false,
  onClick,
}: buttonProps) => {
  return (
    <button
      className="w-full bg-blue-600 text-white font-semibold p-3 rounded-lg disabled:bg-gray-400 disabled:text-gray-600 disabled:cursor-not-allowed"
      disabled={disabled}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Button;

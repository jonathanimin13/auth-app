import { ReactNode } from "react";

interface inputProps {
  type: "text" | "password";
  name: string;
  placeholder: string;
  label?: string;
  required?: boolean;
  button?: ReactNode;
  onButtonClick?: () => void;
}

export const Input = ({
  type,
  name,
  placeholder,
  label,
  required = false,
  button,
  onButtonClick,
}: inputProps) => {
  return (
    <div className="flex flex-col gap-2">
      {label && (
        <label htmlFor={name} className="">
          {label}
          {required && <span className="text-red-600 pl-1">*</span>}
        </label>
      )}
      <div className="relative">
        <input
          type={type}
          id={name}
          name={name}
          placeholder={placeholder}
          className={`w-full border border-black py-2 ${
            button ? "pl-4 pr-16" : "px-4"
          } rounded-lg`}
        />
        {button && (
          <div
            className="absolute right-3 top-1/2 -translate-y-1/2 cursor-pointer"
            onClick={onButtonClick}
          >
            {button}
          </div>
        )}
      </div>
    </div>
  );
};

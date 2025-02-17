import { CapitalizeFirstLetter } from "@/app/utils/formatter";

interface errorCardProps {
  errors: string;
}

export const ErrorCard = ({ errors }: errorCardProps) => {
  const errs = errors.split(", ").map((error, idx) => {
    return <p key={idx}>{CapitalizeFirstLetter(error)}</p>;
  });

  return (
    <div className="bg-red-200 border border-red-600 text-red-600 font-semibold rounded-lg px-4 py-2">
      {errs}
    </div>
  );
};

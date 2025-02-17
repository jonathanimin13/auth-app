import { CapitalizeFirstLetter } from "@/app/utils/formatter";

interface errorCardProps {
  errors: string;
}

export const ErrorCard = ({ errors }: errorCardProps) => {
  const renderErrors = () => {
    if (errors.includes("internal server error")) {
      return <p>Sorry, something went wrong. Please try again!</p>;
    }

    const errs = errors.split(", ").map((error, idx) => {
      return <p key={idx}>{CapitalizeFirstLetter(error)}</p>;
    });
    return errs
  };
  return (
    <div className="bg-red-200 border border-red-600 text-red-600 font-semibold rounded-lg px-4 py-2">
      {renderErrors()}
    </div>
  );
};

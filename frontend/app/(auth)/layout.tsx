export const AuthLayout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  return (
    <div className="w-full h-screen flex justify-center items-center">
      <div className="w-[80%] md:w-[610px] mx-auto border border-black rounded-lg">
        {children}
      </div>
    </div>
  );
};

export default AuthLayout;

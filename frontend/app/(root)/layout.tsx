import Navbar from "../ui/navbar/navbar";

export const UserLayout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  return (
    <div className="w-full h-screen flex flex-col">
      <Navbar />
      <div className="grow">{children}</div>
    </div>
  );
};

export default UserLayout;

import { cookies } from "next/headers";
import { Response, User } from "../types/api";

export default async function Home() {
  const cookieStore = await cookies();
  const token = cookieStore.get("access-token");
  const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_API_URL}/auth/user`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token?.value}`,
    },
  });
  const data: Response<User> = await response.json();

  return (
    <div className="w-[90%] md:w-[80%] h-full mx-auto p-4 flex flex-col justify-center items-center gap-4">
      <p className="font-semibold text-xl md:text-2xl text-center">
        Welcome, {data.data?.username}
      </p>
      <p className="font-semibold text-3xl md:text-4xl text-center">
        This is the homepage
      </p>
    </div>
  );
}

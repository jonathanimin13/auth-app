import { NextRequest, NextResponse } from "next/server";

export const middleware = async (req: NextRequest) => {
  const protectedRoute = ["/"];
  const privateRoute = ["/login", "/register"];

  const currentRoute = req.nextUrl.pathname;
  const isProtectedRoute = protectedRoute.includes(currentRoute);
  const isPrivateRoute = privateRoute.includes(currentRoute);

  const token = req.cookies.get("access-token");

  if (isProtectedRoute && !token) {
    const url = req.nextUrl.clone();
    url.pathname = "/login";
    return NextResponse.redirect(url);
  }

  if (isPrivateRoute && token) {
    const url = req.nextUrl.clone();
    url.pathname = "/";
    return NextResponse.redirect(url);
  }

  if (isProtectedRoute && token) {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_BACKEND_API_URL}/auth/verify-token`,
        {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token?.value}`,
          },
        }
      );

      if (!response.ok) {
        const url = req.nextUrl.clone();
        url.pathname = "/login";
        const res = NextResponse.redirect(url);
        res.cookies.set("access-token", "", { maxAge: 0 });
        return res;
      }
    } catch (error) {
      const url = req.nextUrl.clone();
      url.pathname = "/login";
      const res = NextResponse.redirect(url);
      res.cookies.set("access-token", "", { maxAge: 0 });
      return res;
    }
  }

  return NextResponse.next();
};

export default middleware;

import { NextRequest, NextResponse } from "next/server";

export const middleware = (req: NextRequest) => {
  console.log("middleware here");

  const protectedRoute = ["/"];
  const privateRoute = ["/login", "/register"];
  
  const currentRoute = req.nextUrl.pathname;
  const isProtectedRoute = protectedRoute.includes(currentRoute);
  const isPrivateRoute = privateRoute.includes(currentRoute);

  const accessToken = req.cookies.get("access-token");

  if (isProtectedRoute && !accessToken) {
    const url = req.nextUrl.clone();
    url.pathname = "/login";
    return NextResponse.redirect(url);
  }

  if (isPrivateRoute && accessToken) {
    const url = req.nextUrl.clone();
    url.pathname = "/";
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
};

export default middleware;

import { NextRequest } from "next/server";

export const middleware = async (req: NextRequest) => {
  const protectedRoute = ['/']
  const currentRoute = req.nextUrl.pathname
  const isProtectedRoute = protectedRoute.includes(currentRoute)

  console.log(req.cookies)
  

}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image).*)']
}
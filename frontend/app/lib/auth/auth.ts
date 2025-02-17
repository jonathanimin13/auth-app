import { FetchError } from "@/app/types/api";
import NextAuth, { DefaultSession } from "next-auth";
import Credentials from "next-auth/providers/credentials";

declare module "next-auth/jwt" {
  interface JWT {
    username: string;
    access_token: string;
  }
}

declare module "next-auth" {
  interface Session {
    user: {
      username?: string;
      access_token?: string;
    } & DefaultSession["user"];
  }

  interface Account {
    username: string;
    access_token: string;
  }
}

const handlers = NextAuth({
  providers: [
    Credentials({
      name: "credentials",
      credentials: {
        email: {},
        password: {},
      },

      async authorize(credentials) {
        const response = await fetch("http://localhost:8080/api/auth/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            email: credentials?.email,
            password: credentials?.password,
          }),
        });

        const data = await response.json();

        if (!response.ok) {
          const errors = data.errors as FetchError[];
          const errorMessage = errors.map((error) => error.detail).join(", ");
          throw new Error(errorMessage);
        }

        return data.data;
      },
    }),
  ],
  session: {
    strategy: "jwt",
  },
  callbacks: {
    async jwt({ token, account }) {
      if (account) {
        token.username = account.username;
        token.access_token = account.access_token;
      }
      return token;
    },
    async session({ session, token }) {
      session.user.username = token.username;
      session.user.access_token = token.access_token;
      return session;
    },
  },
});

export { handlers as GET, handlers as POST };

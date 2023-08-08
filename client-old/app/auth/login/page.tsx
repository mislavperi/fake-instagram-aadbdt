import LoginForm from "@/components/loginForm";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

type User = {
  firstName: string;
  lastName: string;
  email: string;
  username: string;
  plan: Plan;
};

type Plan = {
  planName: string;
  uploadLimiSizeKb: number;
  dailyUploadLimit: number;
  cost: number;
};

async function whoami(accessToken: string, refreshToken: string) {
  const res = await fetch("http://localhost:8080/user/whoami", {
    method: "GET",
    headers: {
      Accept: "application/json",
      Access: accessToken,
      Refresh: refreshToken,
    },
  });

  if (!res.ok) {
    return undefined;
  }

  return res.json();
}

export default async function Login() {
  const cookieJar = cookies();

  const accessToken = cookieJar.get("accessToken")?.value || "";
  const refreshToken = cookieJar.get("refreshToken")?.value || "";

  const data: User = await whoami(accessToken, refreshToken);

  if (data != undefined) {
    if (data?.plan.planName === "") {
      redirect("/welcome/plan");
    }
  }

  return data === undefined ? <LoginForm /> : undefined;
}

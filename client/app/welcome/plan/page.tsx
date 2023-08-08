import PlanSelection from "@/components/planSelection";
import { cookies } from "next/headers";

type Plan = {
  planName: string;
  cost: number;
  uploadLimitSizeKb: number;
  dailyUploadLimit: number;
};

async function getPlans() {
  const res = await fetch("http://localhost:8080/plans/get");

  if (!res.ok) {
  }
  return res.json();
}

export default async function Plan() {
  const cookieStore = cookies();
  const data: Plan[] = await getPlans();

  const accessToken = cookieStore.get("accessToken")?.value || "";
  const refreshToken = cookieStore.get("refreshToken")?.value || "";

  return (
    <PlanSelection
      plans={data}
      accessToken={accessToken}
      refreshToken={refreshToken}
    />
  );
}

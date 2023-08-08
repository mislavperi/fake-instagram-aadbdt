"use client";

import { Button } from "@/components/ui/button";
import { useState } from "react";

type Plan = {
  planName: string;
  cost: number;
  uploadLimitSizeKb: number;
  dailyUploadLimit: number;
};

export default function PlanSelection({ plans, accessToken, refreshToken }: { plans: Plan[], accessToken: string, refreshToken: string }) {
  const [selectedPlan, setSelectedPlan] = useState<Plan>();

  async function selectPlan() {
    const result = await fetch("http://localhost:8080/user/select", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Authorization": accessToken,
        "Refresh": refreshToken,
      },
      body: JSON.stringify(selectedPlan),
    });
    if (!result.ok) {
    }
  }

  return (
    <main className="flex flex-col w-full h-screen">
      <div className="flex">
        {plans != null
          ? plans.map((plan) => {
              return (
                <div
                  className={`bg-zinc-100 w-max p-10 m-5 items-center flex flex-col rounded-xl border-black border ${
                    plan === selectedPlan ? "bg-green-100" : ""
                  }`}
                  key={plan.planName}
                >
                  <p className="text-center p-2 text-5xl">{plan.planName}</p>
                  <p className="text-center">Price: {plan.cost}</p>
                  <p className="text-center">
                    Daily upload limit: {plan.dailyUploadLimit}
                  </p>
                  <p className="text-center">
                    Max upload size: {plan.uploadLimitSizeKb} MB
                  </p>
                  <Button
                    className="w-max mt-16"
                    onClick={() => setSelectedPlan(plan)}
                  >
                    Pick this plan
                  </Button>
                </div>
              );
            })
          : null}
      </div>

      <div className="w-full flex justify-end">
        <Button onClick={selectPlan}>Select a plan</Button>
      </div>
    </main>
  );
}

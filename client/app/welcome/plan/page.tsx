"use client";

import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import Link from "next/link"

type Plan = {
  PlanName: string;
  Cost: number;
  UploadLimitSizeKb: number;
  DailyUploadLimit: number;
};

export default function Plan() {
  useEffect(() => {
    fetch("http://localhost:8080/plans/get")
      .then((res) => res.json())
      .then((res) => setPlans(res));
  }, []);

  async function selectPlan() {
    const result = await fetch("http://localhost:8080/user/select", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      body: JSON.stringify(selectedPlan)
    })
    if (!result.ok) {

    }
  }

  const [plans, setPlans] = useState<Plan[]>([]);
  const [selectedPlan, setSelectedPlan] = useState<Plan>();

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
                  key={plan.PlanName}
                >
                  <p className="text-center p-2 text-5xl">{plan.PlanName}</p>
                  <p className="text-center">Price: {plan.Cost}</p>
                  <p className="text-center">
                    Daily upload limit: {plan.DailyUploadLimit}
                  </p>
                  <p className="text-center">
                    Max upload size: {plan.UploadLimitSizeKb} MB
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
           <p onClick={selectPlan}>Select a plan</p>
      </div>
    </main>
  );
}

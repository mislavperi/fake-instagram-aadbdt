"use client"

import { Button } from "@/components/ui/button"
import { useState } from "react"

type Plans = {
    title: string;
    price: number;
    uploadLimit: number;
    maxUploadSize: number;
}

export default function Plan(){
    const [planss, setPlans] = useState<Plans[]>([])

    const plans: { title: string, price: number, uploadLimit: number, maxUploadSize: number }[] = [
        {
            title: "FREE",
            price: 0,
            uploadLimit: 20,
            maxUploadSize: 3000
        },
        {
            title: "FREE",
            price: 0,
            uploadLimit: 20,
            maxUploadSize: 3000
        },
        {
            title: "FREE",
            price: 0,
            uploadLimit: 20,
            maxUploadSize: 3000
        },
    ]

    return (
        <main className="flex w-full h-screen items-center justify-center">
            {plans.map(plan => {
                return (
                    <div className="bg-zinc-100 w-max p-10 m-5 items-center flex flex-col rounded-xl border-black border">
                        <p className="text-center p-2 text-5xl">{plan.title}</p>
                        <p className="text-center">Price: {plan.price}</p>
                        <p className="text-center">Daily upload limit: {plan.uploadLimit}</p>
                        <p className="text-center">Max upload size: {plan.maxUploadSize} MB</p>
                        <Button className="w-max mt-16">Pick this plan</Button>
                    </div>
                )
            })}
        </main>
    )
}
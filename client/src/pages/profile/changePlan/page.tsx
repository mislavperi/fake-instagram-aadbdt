// @ts-nocheck
import { Flex, Text, Button, Wrap, WrapItem } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import Plan from "../../../types/plan";

export default function ChangePlan() {
  const cookieJar = new Cookies();
  const [plans, setPlans] = useState<Plan[]>();
  const [selectedPlan, setSelectedPlan] = useState<Plan>();

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  useEffect(() => {
    fetch("/api/plans/get", {
      headers: {
        Authorization: cookieJar.get("accessToken"),
        Refresh: cookieJar.get("refreshToken"),
        Accept: "application/json"
      }
    })
      .then((res) => res.json())
      .then((res) => setPlans(res));
  }, []);

  const selectPlan = () => {
    fetch("/api/user/select", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Authorization: accessToken,
        Refresh: refreshToken,
      },
      body: JSON.stringify(selectedPlan),
    });
  };

  return plans != null ? (
    <Flex width="100vw" justify="center" direction="column">
      <Wrap width="100vw" justify="center">
        {plans.map((plan) => {
          return (
            <WrapItem key={plan.planID}>
              <Flex
                direction="column"
                m={10}
                key={plan.planName}
                align="center"
                p={15}
                background={plan === selectedPlan ? "green" : ""}
              >
                <Text fontSize="72px" p={0} m={0}>
                  {plan.planName}
                </Text>
                <Text>Cost: {plan.cost} euro</Text>
                <Text>Daily upload limit: {plan.dailyUploadLimit} images</Text>
                <Text>Upload limit size: {plan.uploadLimitSizeKb} kb</Text>
                <Button onClick={() => setSelectedPlan(plan)}>
                  Select this plan
                </Button>
              </Flex>
            </WrapItem>
          );
        })}
      </Wrap>

      <Button onClick={selectPlan}>Select a plan</Button>
    </Flex>
  ) : null;
}

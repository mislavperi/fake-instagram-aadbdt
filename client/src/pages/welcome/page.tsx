import { Flex, Text, Button } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { useNavigate } from "react-router-dom";

export default function Welcome() {
  const cookieJar = new Cookies();
  const [plans, setPlans] = useState<Plan[]>();
  const [selectedPlan, setSelectedPlan] = useState<Plan>();
  const navigate = useNavigate();

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  useEffect(() => {
    fetch("http://localhost:8080/plans/get")
      .then((res) => res.json())
      .then((res) => setPlans(res));
  }, []);

  const selectPlan = () => {
    fetch("http://localhost:8080/user/select", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Authorization: accessToken,
        Refresh: refreshToken,
      },
      body: JSON.stringify(selectedPlan),
    }).then((res) => {
      if (res.ok) {
        navigate("/login");
      }
    });
  };

  type Plan = {
    planName: string;
    cost: number;
    uploadLimitSizeKb: number;
    dailyUploadLimit: number;
  };

  return plans != null ? (
    <Flex width="100vw" justify="center" direction="column">
      <Flex width="100vw" justify="center">
        {plans.map((plan) => {
          return (
            <Flex
              direction="column"
              m={10}
              key={plan.planName}
              border="1px solid black"
              borderRadius="10px"
              align="center"
              p={15}
              background={plan === selectedPlan ? "green": ""}

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
          );
        })}
      </Flex>

      <Button onClick={selectPlan}>Select a plan</Button>
    </Flex>
  ) : null;
}

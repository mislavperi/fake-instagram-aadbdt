import { useState, useEffect } from "react";
import Cookies from "universal-cookie";
import { useLocation } from "react-router-dom";
import {
  VStack,
  Text,
  Stat,
  StatLabel,
  StatNumber,
  Avatar,
  Button,
  Input,
  HStack,
  Flex,
  Wrap,
  WrapItem,
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
} from "@chakra-ui/react";
import ExpandedStatistics from "../../../../types/expandedStatistics";
import Plan from "../../../../types/plan";
import Log from "../../../../types/log";

export default function UserStatistics() {
  const [user, setUser] = useState<ExpandedStatistics>({});
  const [plans, setPlans] = useState<Plan[]>({});
  const [selectedPlan, setSelectedPlan] = useState<Plan>();
  const [logs, setLogs] = useState<Log[]>();

  const cookieJar = new Cookies();

  const location = useLocation();
  const id = location.state;

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  const updateUserValue = (value: string, field: string) => {
    const objectToUpdate = structuredClone(user);
    objectToUpdate.user[field] = value;
    setUser(objectToUpdate);
  };

  useEffect(() => {
    fetch(`http://localhost:8080/admin/statistics?id=${id}`, {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json",
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res) => setUser(res));
      }
    });
  }, []);

  useEffect(() => {
    fetch("http://localhost:8080/plans/get", {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json",
      },
    })
      .then((res) => res.json())
      .then((res) => setPlans(res));
  }, []);

  useEffect(() => {
    fetch(`http://localhost:8080/admin/userlogs?id=${id}`, {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json"
      }
    })
    .then(res => {
      if (res.ok) {
        res.json().then(res => setLogs(res))
      }
    })
  })

  const updateUserPlan = () => {
    fetch(`http://localhost:8080/admin/changePlan?id=${id}&planId=${selectedPlan?.planID}`, {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
      }
    })
    .then(res => {
      if (res.ok) {
        console.log(res)
      }
    })
  }

  return Object.keys(user).length != 0 ? (
    <Flex align="center" justify="center" wrap="wrap" background="white">
      <VStack align="center">
        <Avatar src={user.user.pictureURI} />
        <HStack>
          <Text>Username:</Text>
          <Input
            value={user.user.username}
            onChange={(v) => updateUserValue(v.target.value, "username")}
          />
        </HStack>
        <HStack>
          <Text>E-mail:</Text>
          <Input
            value={user.user.email}
            onChange={(v) => updateUserValue(v.target.value, "email")}
          />
        </HStack>{" "}
        <HStack>
          <Text>First name:</Text>
          <Input
            value={user.user.firstName}
            onChange={(v) => updateUserValue(v.target.value, "firstName")}
          />
        </HStack>{" "}
        <HStack>
          <Text>Last name:</Text>
          <Input
            value={user.user.lastName}
            onChange={(v) => updateUserValue(v.target.value, "lastName")}
          />
        </HStack>{" "}
        <Text>Plan name: {user.plan.planName}</Text>
        <Text>{user.user.lastName}</Text>
        <Stat align="center">
          <StatLabel>Total consumption kb</StatLabel>
          <StatNumber>
            {user.totalConsumptionKb} / {user.plan.uploadLimitSizeKb}
          </StatNumber>
        </Stat>
        <Stat align="center">
          <StatLabel>Daily uploads</StatLabel>
          <StatNumber>{user.totalDailyUploadCount || 0} </StatNumber>
        </Stat>
        <Stat align="center">
          <StatLabel>Total images uploaded</StatLabel>
          <StatNumber>{user.totalConsumptionCount}</StatNumber>
        </Stat>
        <Button>Apply changes</Button>
      </VStack>
      <Flex direction="column">
        <HStack>
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
                  <Text>
                    Daily upload limit: {plan.dailyUploadLimit} images
                  </Text>
                  <Text>Upload limit size: {plan.uploadLimitSizeKb} kb</Text>
                  <Button onClick={() => setSelectedPlan(plan)}>
                    Select this plan
                  </Button>
                </Flex>
              </WrapItem>
            );
          })}
        </HStack>
        <Button onClick={updateUserPlan}>Change plan for user</Button>
      </Flex>
      <TableContainer>
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th>Log ID</Th>
              <Th>Action</Th>
              <Th>Timestamp</Th>
            </Tr>
          </Thead>
          <Tbody>
          {logs?.length !== 0 ? logs?.map(log => {
            const timestamp = new Date(log.timestamp)
          return (
            <Tr>
              <Td>{log.id}</Td>
              <Td>{log.action}</Td>
              <Td>{timestamp.toLocaleString()}</Td>
            </Tr>
          )
        }) : null}
          </Tbody>

        </Table>
      </TableContainer>
    </Flex>
  ) : null;
}

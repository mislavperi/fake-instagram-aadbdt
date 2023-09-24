// @ts-nocheck
import { Flex, Stat, StatLabel, StatNumber, useToast } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import Stats from "../../../types/statistics";

export default function Statistics() {
  const [statistics, setStatistics] = useState<Stats>({});

  const toast = useToast();

  const cookieJar = new Cookies();

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  useEffect(() => {
    fetch("/api/statistics/get", {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json",
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res) => setStatistics(res));
      } else {
        res.json().then((res) => {
          toast({
            description: res,
          });
        });
      }
    });
  }, []);

  return Object.keys(statistics).length != 0 ? (
    <Flex>
      <Stat>
        <StatLabel>Total consumption kb</StatLabel>
        <StatNumber>
          {statistics.totalConsumptionKb} / {statistics.plan.uploadLimitSizeKb}
        </StatNumber>
      </Stat>
      <Stat>
        <StatLabel>Daily uploads</StatLabel>
        {statistics.dailyUploadCount != 0 ? (
          <StatNumber>{statistics.dailyUploadCount} </StatNumber>
        ) : (
          <StatNumber>0</StatNumber>
        )}
      </Stat>
      <Stat>
        <StatLabel>Total images uploaded</StatLabel>
        <StatNumber>{statistics.totalConsumptionCount}</StatNumber>
      </Stat>
    </Flex>
  ) : null;
}

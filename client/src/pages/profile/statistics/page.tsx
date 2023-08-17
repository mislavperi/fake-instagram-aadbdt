import {
  Flex,
  Stat,
  StatLabel,
  StatNumber,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import Stats from "../../../types/statistics";

export default function Statistics() {
  const [statistics, setStatistics] = useState<Stats>({});

  const cookieJar = new Cookies();

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  useEffect(() => {
    fetch("http://localhost:8080/statistics/get", {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json",
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res) => setStatistics(res));
      }
    });
  }, []);

  return Object.keys(statistics).length != 0 ? (<Flex>
    <Stat>
        <StatLabel>Total consumption kb</StatLabel>
        <StatNumber>{statistics.totalConsumptionKb} / {statistics.plan.uploadLimitSizeKb}</StatNumber>
    </Stat>
    <Stat>
        <StatLabel>Daily uploads</StatLabel>
        <StatNumber>{statistics.totalDailyUploadCount || 0} </StatNumber>
    </Stat>
    <Stat>
        <StatLabel>Total images uploaded</StatLabel>
        <StatNumber>{statistics.totalConsumptionCount}</StatNumber>
    </Stat>
  </Flex>) : null;
}

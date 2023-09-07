// @ts-nocheck
import {
  Button,
  Text,
  Wrap,
  WrapItem,
  Container,
  Box,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
} from "@chakra-ui/react";

import ChangePlan from "./changePlan/page";
import Admin from "./admin/page";

import { useEffect, useState } from "react";
import User from "../../types/user";
import { Link } from "react-router-dom";

import Cookies from "universal-cookie"
import Statistics from "./statistics/page";

interface Picture {
  [key: string]: any;
  id: number;
  title: string;
  description: string;
  pictureURI: string;
  uploadDateTime: string;
  hashtags: string[];
  user: User;
}

export default function Profile() {
  const [pictures, setPictures] = useState<Picture[] | []>([]);

  const cookieJar = new Cookies();

  useEffect(() => {
    fetch("/api/picture/userImages", {
      headers: {
        Authorization: cookieJar.get("accessToken"),
        Refresh: cookieJar.get("refreshToken"),
        Accept: "application/json"
      }
    })
      .then((res) => res.json())
      .then((res) => setPictures(res));
  }, []);
  return (
    <div>
      <Tabs variant="enclosed">
        <TabList>
          <Tab>My images</Tab>
          <Tab>Update Plan</Tab>
          <Tab>My statistics</Tab>
          <Tab>Admin</Tab>
        </TabList>
        <TabPanels>
          <TabPanel>
            {pictures !== null ? (
              <Wrap justify="flex-start">
                {pictures.map((picture) => {
                  return (
                    <WrapItem key={picture.id}>
                      <img src={picture.pictureURI} width="600px" />
                      <Container>
                      useNavigate  <Box>
                          <Text fontSize="18px">Title: {picture.title}</Text>
                        </Box>
                        <Box>
                          <Text fontSize="14px">
                            Description: {picture.description}
                          </Text>
                        </Box>
                        <Link to={`/edit/${picture.id}`} state={picture.id}>
                          Edit image
                        </Link>
                      </Container>
                    </WrapItem>
                  );
                })}
              </Wrap>
            ) : null}
          </TabPanel>
          <TabPanel>
            <ChangePlan />
          </TabPanel>
          <TabPanel>
            <Statistics />
          </TabPanel>
          <TabPanel>
            <Admin />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </div>
  );
}

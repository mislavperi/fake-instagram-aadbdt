// @ts-nocheck
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import Cookies from "universal-cookie";
import Picture from "../../../../types/picture";
import {
  Flex,
  Input,
  IconButton,
  Text,
  Stack,
  VStack,
  WrapItem,
  Wrap,
} from "@chakra-ui/react";
import { EditIcon } from "@chakra-ui/icons";
import { useNavigate } from "react-router-dom";


export default function UserPictures() {
  const [pictures, setPictures] = useState<Picture[]>([]);
  const location = useLocation();
  const userID = location.state;
  const cookieJar = new Cookies();
  const navigate = useNavigate()

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  useEffect(() => {
    fetch(`/api/admin/userPictures?id=${userID}`, {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json",
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res) => setPictures(res));
      }
    });
  }, []);

  return pictures.length != 0 ? (
    <Wrap align="center">
      {pictures.map((picture) => {
        return (
          <WrapItem>
            <VStack key={picture.id}>
              <img src={picture.pictureURI} alt={picture.description} width={400} />
              <Text>Title: {picture.title}</Text>
              <IconButton aria-label="edit" icon={<EditIcon />} onClick={() => navigate(`/edit/${picture.id}`, {state: picture.id})}></IconButton>
            </VStack>
          </WrapItem>
        );
      })}
    </Wrap>
  ) : null;
}

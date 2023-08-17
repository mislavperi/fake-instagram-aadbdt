import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import User from "../../../types/user";
import {
  Flex,
  Text,
  Stack,
  Avatar,
  IconButton,
  VStack,
} from "@chakra-ui/react";
import { EditIcon, AttachmentIcon } from "@chakra-ui/icons";
import { useNavigate }  from "react-router-dom";

export default function Admin() {
  const [users, setUsers] = useState<User[]>([]);
  const cookieJar = new Cookies();
  const navigate = useNavigate();

  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  useEffect(() => {
    fetch("http://localhost:8080/admin/users", {
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
        Accept: "application/json",
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res) => setUsers(res));
      }
    });
  }, []);

  return users.length != 0 ? (
    <Flex>
      {users.map((user) => {
        return (
          <Stack direction="row" align="center" key={user.username}>
            <Avatar src={user.pictureURI} />
            <VStack align="left">
              <Text>First name:{user.firstName}</Text>
              <Text>Last name:{user.lastName}</Text>
              <Text>Username: {user.username}</Text>
            </VStack>
            <IconButton aria-label="edit-user" icon={<EditIcon />} onClick={() => navigate(`/statistics/${user.id}`, {state: user.id})}></IconButton>
            <IconButton aria-label="edit-user" icon={<AttachmentIcon />} onClick={() => navigate(`/userimages/${user.id}`, {state: user.id})}></IconButton>
          </Stack>
        );
      })}
    </Flex>
  ) : null;
}

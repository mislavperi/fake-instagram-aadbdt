import {
  Box,
  Flex,
  Avatar,
  Button,
  useColorModeValue,
  Stack,
  Text,
  HStack,
} from "@chakra-ui/react";
import { Outlet } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { useUserContext } from "../context/userContext";

interface NavLink {
  title: string;
  path: string;
}

const Links: NavLink[] = [
  {
    title: "Home",
    path: "/home",
  },
  {
    title: "Upload image",
    path: "/upload",
  },
];

const NavLink = ({ link }: { link: NavLink }) => {
  const navigate = useNavigate();

  return (
    <Box
      as="a"
      px={2}
      py={1}
      rounded={"md"}
      _hover={{
        textDecoration: "none",
        bg: useColorModeValue("gray.200", "gray.700"),
      }}
      onClick={() => navigate(link.path)}
    >
      <Text>{link.title}</Text>
    </Box>
  );
};

export default function Simple() {
  const navigate = useNavigate();
  const { user } = useUserContext();

  return (
    <div style={{ height: "100vh", background: "white" }}>
      <Flex
        justify="space-between"
        width="100vw"
        align="center"
        py={5}
        borderBottom="1px solid black"
      >
        <HStack>
          {Links.map((link: NavLink) => {
            return <NavLink link={link}/>;
          })}
        </HStack>
        <Stack direction="row" align="center">
          {user.username !== "" ? (
            <Text onClick={() => navigate("/profile")}>{user.username}</Text>
          ) : (
            <Button onClick={() => navigate("/login")}>Log in</Button>
          )}
          <Avatar
            mx={2}
            size="md"
            name="Dan Abrahmov"
            src="https://bit.ly/dan-abramov"
            borderRadius="50%"
          />
        </Stack>
      </Flex>
      <Box>
        <Outlet />
      </Box>
    </div>
  );
}

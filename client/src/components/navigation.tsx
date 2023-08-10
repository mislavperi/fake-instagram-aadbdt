import {
  Box,
  Flex,
  Avatar,
  Button,
  useColorModeValue,
  Stack,
  Text,
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
];

const NavLink = ({ link }: { link: NavLink }) => {
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
    >
      {link.title}
    </Box>
  );
};

export default function Simple() {
  const navigate = useNavigate();
  const { user } = useUserContext();

  return (
    <div style={{ height: "100vh" }}>
      <Flex
        justify="space-between"
        width="100vw"
        bg="red"
        align="center"
        py={5}
      >
        <Box>
          {Links.map((link: NavLink) => {
            return <NavLink link={link} />;
          })}
        </Box>
        <Stack direction="row" align="center">
          {user.username !== "" ? (
            <Text>{user.username}</Text>
          ) : (
            <Button onClick={() => navigate("/login")}>Log in</Button>
          )}

          <Avatar
            name="Dan Abrahmov"
            src="https://bit.ly/dan-abramov"
            w={64}
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

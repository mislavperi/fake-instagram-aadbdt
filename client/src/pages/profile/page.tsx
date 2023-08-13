import { Button, Text, Wrap, WrapItem, Container, Box } from "@chakra-ui/react";

import { useUserContext } from "../../context/userContext";
import { useEffect, useState } from "react";
import User from "../../types/user";
import { Link } from "react-router-dom";

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
  const { user } = useUserContext();

  const [pictures, setPictures] = useState<Picture[] | []>([]);

  useEffect(() => {
    fetch("http://localhost:8080/picture/userImages")
      .then((res) => res.json())
      .then((res) => setPictures(res));
  }, []);
  return (
    <div>
      {pictures.length != 0 ? (
        <Wrap justify="flex-start">
          {pictures.map((picture) => {
            return (
              <WrapItem key={picture.id}>
                <img src={picture.pictureURI} width="50%" />
                <Container bg="red">
                  <Box>
                    <Text>Title: {picture.title}</Text>
                  </Box>
                  <Box>
                    <Text>Description: {picture.description}</Text>
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
    </div>
  );
}

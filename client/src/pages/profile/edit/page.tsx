// @ts-nocheck
import {
  Flex,
  Text,
  Editable,
  EditableInput,
  EditablePreview,
  Button,
  useToast
} from "@chakra-ui/react";

import ChakraTagInput from "../../../components/inputTag";
import PartialPicture from "../../../types/partialPicture";

import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { useLocation } from "react-router-dom";

export default function Edit() {
  const cookies = new Cookies();
  const location = useLocation();
  const id = location.state;
  const toast = useToast()

  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [pictureURI, setPictureURI] = useState<string>("");
  const [hashtags, setHashtags] = useState<string[]>([]);

  useEffect(() => {
    fetch(`/api/picture/info?id=${id}`, {
      method: "GET",
      headers: {
        Authorization: cookies.get("accessToken"),
        Refresh: cookies.get("refreshToken"),
        Accept: "application/json",
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res: PartialPicture) => {
          updateStates(
            res.title,
            res.description,
            res.pictureURI,
            res.hashtags
          );
        });
      }
    });
  }, []);

  const submitChanges = () => {
    console.log(JSON.stringify(hashtags))
    const body = JSON.stringify({
      id: id,
      title: title,
      description: description,
      hashtags: hashtags
    })
    fetch("/api/picture/update", {
      method: "POST",
      headers: {
        Authorization: cookies.get("accessToken"),
        Refresh: cookies.get("refreshToken"),
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: body
    })
    .then(res => {
      if (res.ok) {
        toast({
          "description": "Successfully updated image"
        })
      }
    })
  }

  const updateStates = (
    title: string,
    description: string,
    pictureURI: string,
    pictureHashtags: any
  ) => {
    setTitle(title);
    setDescription(description);
    setPictureURI(pictureURI);
    if (pictureHashtags.length != 0)
    setHashtags(pictureHashtags);
  };

  return title !== "" ? (
    <Flex
      direction="column"
      align="center"
    >
      <img src={pictureURI} alt={description} width={600} />
      <Text fontSize="18px" p={0} m={0}>
        {title}
      </Text>
      <Editable value={description} onChange={setDescription}>
        <EditablePreview />
        <EditableInput />
      </Editable>
      <Text fontStyle="bold">Hashtags:</Text>
      <ChakraTagInput tags={hashtags} setTags={setHashtags} />
      <Button onClick={submitChanges}>Update image</Button>
    </Flex>
  ) : null;
}

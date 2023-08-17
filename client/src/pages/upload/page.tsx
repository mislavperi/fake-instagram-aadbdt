import Cookies from "universal-cookie";

import { useContext, useState, useCallback, SyntheticEvent, useEffect } from "react";

import {
  Button,
  Input,
  FormLabel,
  Container,
  RadioGroup,
  Radio,
  Stack,
  NumberInput,
  NumberInputField,
  Flex
} from "@chakra-ui/react";

import { UserContext } from "../../context/userContext";

import ChakraTagInput from "../../components/inputTag";

export default function UploadImage() {
  const cookieJar = new Cookies();
  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");

  const user = useContext(UserContext);

  useEffect(() => {
    fetch("http://localhost:8080/consumption/get", {
      headers: {
        Accept: "application/json",
        Authorization: accessToken,
        Refresh: refreshToken
      }
    })
  },[])

  const [picture, setPicture ] = useState<string>();
  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [hashtags, setHashtags] = useState<string[]>([]);
  const [format, setFormat] = useState<string>("jpeg");
  const [height, setHeight] = useState<string>("0");
  const [width, setWidth] = useState<string>("0");
  const handleTagsChange = useCallback(
    (event: SyntheticEvent, tags: string[]) => {
      setHashtags(tags);
    },
    []
  );

  function handleChange(e: any) {
    e.preventDefault();
    setPicture(e.target.files[0]);
  }

  function submitFile() {
    const formData = new FormData();
    if (picture != undefined) {
      formData.append("file", picture);
      formData.append("title", title);
      formData.append("description", description);
      formData.append("hashtags", JSON.stringify(hashtags));
      formData.append("user", JSON.stringify(user));
      formData.append("format", format);
      formData.append("height", height);
      formData.append("width", width);
    }

    fetch("http://localhost:8080/picture/upload", {
      method: "POST",
      headers: {
        Authorization: accessToken,
        Refresh: refreshToken,
      },
      body: formData,
    });
  }

  return (
    <Flex display="inline-flex" direction="column">
      <input type="file" onChange={handleChange} />
      <FormLabel>Title of image</FormLabel>
      <Input type="text" onChange={(e) => setTitle(e.target.value)} />
      <FormLabel>Description</FormLabel>
      <Input type="text" onChange={(e) => setDescription(e.target.value)} />
      <Container>
        <ChakraTagInput
          tags={hashtags}
          setTags={setHashtags}
        />
      </Container>
      <RadioGroup onChange={setFormat} value={format}>
        <Stack direction="row">
          <Radio value="jpeg">JPEG</Radio>
          <Radio value="bmp">BMP</Radio>
          <Radio value="png">PNG</Radio>
        </Stack>
      </RadioGroup>
        <FormLabel>Height</FormLabel>
        <NumberInput value={width} onChange={(value) => setHeight(value)}>
          <NumberInputField/>
        </NumberInput>

        <FormLabel>Width</FormLabel>
        <NumberInput value={width} onChange={(value) => setWidth(value)} >
          <NumberInputField />
        </NumberInput>
        <Button onClick={submitFile}>Upload</Button>

    </Flex>
  );
}

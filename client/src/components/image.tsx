// @ts-nocheck
import {
  Text,
  HStack,
  Button,
  IconButton,
  Popover,
  PopoverTrigger,
  PopoverContent,
  PopoverHeader,
  PopoverBody,
  PopoverFooter,
  PopoverArrow,
  PopoverCloseButton,
  RadioGroup,
  Radio,
  Stack,
  Flex,
  NumberInput,
  FormLabel,
  NumberInputField,
  Slider,
  SliderTrack,
  SliderFilledTrack,
  SliderThumb,
  SliderMark,
  Box,
} from "@chakra-ui/react";
import Zoom from "react-medium-image-zoom";
import "react-medium-image-zoom/dist/styles.css";

import { DownloadIcon } from "@chakra-ui/icons";

import { useState } from "react";

const Image = ({
  id,
  title,
  description,
  url,
  dateTime,
  hashtags,
}: {
  id: number;
  title: string;
  description: string;
  url: string;
  dateTime: string;
  hashtags: string[];
}) => {
  const [format, setFormat] = useState<string>("jpeg");
  const [height, setHeight] = useState<string>("0");
  const [width, setWidth] = useState<string>("0");
  const [blur, setBlur] = useState<number>(0);
  const [sapia, setSapia] = useState<number>(0);

  const getEditedImage = () => {
    fetch("/api/picture/edited", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/octet-stream",
      },
      body: JSON.stringify({
        id: id,
        format: format,
        height: height,
        width: width,
        blur: blur / 100,
        sapia: sapia / 100,
      }),
    })
      .then((res) => res.blob())
      .then((blob) => {
        const link = document.createElement("a");
        const url = URL.createObjectURL(blob);
        link.href = url;
        link.download = `111.${format}`;
        link.click();
      });
  };

  const labelStyles = {
    mt: "2",
    ml: "-2.5",
    fontSize: "sm",
  };

  return (
    <Flex
      direction="column"
      width="fit-content"
      maxWidth={600}
      p={5}
      borderRadius="5px"
    >
      <Zoom>
        <img src={url} alt={description} />
      </Zoom>
      <Text fontSize="18px" p={0} m={0}>
        {title}
      </Text>
      <Text fontSize="14px" p={0} m={0}>
        {description}
      </Text>
      <Text fontSize="10px" p={0} m={0}>
        {dateTime}
      </Text>
      <HStack>
        {hashtags.map((hashtag) => {
          return (
            <Text fontSize="12px" p={0} m={0} color="grey" key={hashtag}>
              #{hashtag}
            </Text>
          );
        })}
        <Popover>
          <PopoverTrigger>
            <IconButton
              mx={5}
              aria-label="download-image"
              icon={<DownloadIcon />}
            ></IconButton>
          </PopoverTrigger>
          <PopoverContent>
            <PopoverArrow />
            <PopoverCloseButton />
            <PopoverHeader>Download image</PopoverHeader>
            <PopoverBody>
              <FormLabel>Height</FormLabel>
              <NumberInput
                value={height}
                size="sm"
                onChange={(value) => setHeight(value)}
              >
                <NumberInputField />
              </NumberInput>

              <FormLabel>Width</FormLabel>
              <NumberInput
                value={width}
                size="sm"
                onChange={(value) => setWidth(value)}
              >
                <NumberInputField />
              </NumberInput>
              <Box pt={6} pb={2} my={2}>
                <Text>Sapia: </Text>
                <Slider
                  aria-label="slider-ex-6"
                  onChange={(val) => setSapia(val)}
                  defaultValue={0}
                >
                  <SliderMark value={25} {...labelStyles}>
                    25%
                  </SliderMark>
                  <SliderMark value={50} {...labelStyles}>
                    50%
                  </SliderMark>
                  <SliderMark value={75} {...labelStyles}>
                    75%
                  </SliderMark>
                  <SliderMark
                    value={sapia}
                    textAlign="center"
                    bg="blue.500"
                    color="white"
                    mt="-10"
                    ml="-5"
                    w="12"
                    borderRadius="10px"
                  >
                    {sapia}%
                  </SliderMark>
                  <SliderTrack>
                    <SliderFilledTrack />
                  </SliderTrack>
                  <SliderThumb />
                </Slider>
              </Box>
              <Box pt={6} pb={2} my={2}>
                <Text>Blur: </Text>
                <Slider
                  aria-label="slider-ex-6"
                  onChange={(val) => setBlur(val)}
                  defaultValue={0}
                >
                  <SliderMark value={25} {...labelStyles}>
                    25%
                  </SliderMark>
                  <SliderMark value={50} {...labelStyles}>
                    50%
                  </SliderMark>
                  <SliderMark value={75} {...labelStyles}>
                    75%
                  </SliderMark>
                  <SliderMark
                    value={blur}
                    textAlign="center"
                    bg="blue.500"
                    color="white"
                    mt="-10"
                    ml="-5"
                    w="12"
                    borderRadius="10px"
                  >
                    {blur}%
                  </SliderMark>
                  <SliderTrack>
                    <SliderFilledTrack />
                  </SliderTrack>
                  <SliderThumb />
                </Slider>
              </Box>
              <RadioGroup onChange={setFormat} value={format} my="2px">
                <Stack direction="row">
                  <Radio value="jpeg">JPEG</Radio>
                  <Radio value="bmp">BMP</Radio>
                  <Radio value="png">PNG</Radio>
                </Stack>
              </RadioGroup>
            </PopoverBody>
            <PopoverFooter>
              <Button m={2} size="sm">
                <a href={url} download>
                  Get image
                </a>
              </Button>
              <Button m={2} size="sm" onClick={getEditedImage}>
                Download edited version
              </Button>
            </PopoverFooter>
          </PopoverContent>
        </Popover>
      </HStack>
    </Flex>
  );
};

export default Image;

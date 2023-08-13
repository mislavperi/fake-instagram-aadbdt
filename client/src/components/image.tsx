import { Flex, Text, HStack } from "@chakra-ui/react";

const Image = ({
  title,
  description,
  url,
  dateTime,
  hashtags,
}: {
  title: string;
  description: string;
  url: string;
  dateTime: string;
  hashtags: string[];
}) => {
  return (
    <Flex
      direction="column"
      bg="red"
      width="fit-content"
      border="1px solid black"
    >
      <img src={url} alt={description} width={300} height={300} />
      <Text fontSize="16px" p={0} m={0}>
        {title}
      </Text>
      <Text fontSize="16px" p={0} m={0}>
        {description}
      </Text>
      <Text fontSize="12px" p={0} m={0}>
        {dateTime}
      </Text>
      <HStack>
        {hashtags.map((hashtag) => {
          return (
            <Text fontSize="12px" p={0} m={0}>
              {hashtag}
            </Text>
          );
        })}
      </HStack>
    </Flex>
  );
};

export default Image;

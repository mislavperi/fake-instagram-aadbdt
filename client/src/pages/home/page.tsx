import { Flex, Text, HStack } from "@chakra-ui/react";
import {  useUserContext } from "../../context/userContext";

export default function Home() {
  const { user } = useUserContext()

  const dummyImage: {
    title: string;
    description: string;
    url: string;
    dateTime: string;
    hashtags: string[];
  }[] = [
    {
      title: "image",
      description: "image",
      url: "https://images.unsplash.com/photo-1638486071992-536e48c8fa3e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2072&q=80",
      dateTime: "01.01.2021",
      hashtags: ["#asgag", "#asgagsaga"],
    },
  ];

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
      <Flex direction="column" bg="red" width="fit-content" border="1px solid black">
        <img src={url} alt={description} width={300} />
        <Text fontSize="16px" p={0} m={0}>{title}</Text>
        <Text fontSize="16px" p={0} m={0}>{description}</Text>
        <Text fontSize="12px" p={0} m={0}>{dateTime}</Text>
        <HStack>
          {hashtags.map(hashtag => {
            return (
              <Text fontSize="12px" p={0} m={0}>{hashtag}</Text>
            )
          })}
          </HStack>
      </Flex>
    );
  };

  return (
    <main>
      <div>
        {dummyImage.map((image) => (
          <Image
            title={image.title}
            description={image.description}
            url={image.url}
            dateTime={image.dateTime}
            hashtags={image.hashtags}
          />
        ))}
      </div>
    </main>
  );
}

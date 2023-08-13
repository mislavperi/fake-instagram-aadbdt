import {
  Flex,
  Text,
  HStack,
  Container,
  Input,
  Button,
  Wrap,
  WrapItem,
} from "@chakra-ui/react";
import { useState, useEffect } from "react";
import Picture from "../../types/picture";
import { RangeDatepicker } from "chakra-dayzed-datepicker";
import InputTag from "../../components/inputTag";
import Image from "../../components/image"

interface Filter {
  [key: string]: any;
  title?: string;
  description?: string;
  dateRange?: DateRange;
  hashtags?: string[];
  user?: string;
}

export default function Home() {
  const [pictures, setPictures] = useState<Picture[]>([]);
  const [hashtags, setHashtags] = useState<string[]>([]);
  const [filter, setFilter] = useState<Filter>({});
  const [selectedDates, setSelectedDates] = useState<Date[]>([
    new Date(),
    new Date(),
  ]);

  const getPictures = () => {
    const queryParam: string = JSON.stringify(filter);

    fetch(`http://localhost:8080/picture/get?filter=${queryParam}`)
      .then((res) => res.json())
      .then((res) => setPictures(res));
  };
  useEffect(() => {
    getPictures();
  }, [true]);

  const updateFilter = (value: string, field: string) => {
    const newFilter: Filter = filter;
    if (value == "") {
      delete newFilter[field];
    } else {
      newFilter[field] = value;
    }
    console.log(newFilter);
    setFilter(newFilter);
  };

  return (
    <main>
      <Container>
        <Wrap>
          <WrapItem>
            <HStack>
              <Text>Title: </Text>
              <Input onChange={(e) => updateFilter(e.target.value, "title")} />
            </HStack>
          </WrapItem>
          <WrapItem>
            <HStack>
              <Text>Description: </Text>
              <Input onChange={(e) => updateFilter(e.target.value, "title")} />
            </HStack>
          </WrapItem>
          <WrapItem>
            <HStack>
              <Text>User</Text>
              <Input onChange={(e) => updateFilter(e.target.value, "title")} />
            </HStack>
          </WrapItem>
          <WrapItem>
            <HStack>
              <Text>Select date range: </Text>
              <RangeDatepicker
                selectedDates={selectedDates}
                onDateChange={setSelectedDates}
              />
            </HStack>
          </WrapItem>
        </Wrap>
        <HStack>
          <Text>Hashtags:</Text>
          <InputTag tags={hashtags} setTags={setHashtags} />
        </HStack>
        <HStack>
          <Button onClick={getPictures}>Apply filter</Button>
        </HStack>
      </Container>

      <Wrap>
        <WrapItem>
          {pictures.length !== 0
            ? pictures.map((image) => (
                <Image
                  key={image.id}
                  title={image.title}
                  description={image.description}
                  url={image.pictureURI}
                  dateTime={image.uploadDateTime}
                  hashtags={image.hashtags}
                />
              ))
            : null}
        </WrapItem>
      </Wrap>
    </main>
  );
}

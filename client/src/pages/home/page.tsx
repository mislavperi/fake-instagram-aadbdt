import {
  Text,
  HStack,
  Input,
  Button,
  Wrap,
  WrapItem,
} from "@chakra-ui/react";
import { useState, useEffect } from "react";
import Picture from "../../types/picture";
import { SingleDatepicker } from "chakra-dayzed-datepicker";
import InputTag from "../../components/inputTag";
import Image from "../../components/image";

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
  const [fromDate, setFromDate] = useState<Date>();
  const [toDate, setToDate] = useState<Date>();

  const getPictures = () => {

    filter.dateRange = {gte: null, lte: null}
    if (fromDate !== undefined && filter.dateRange != undefined) {
      filter.dateRange.gte = fromDate
    }
    if (toDate !== undefined && filter.dateRange != undefined) {
      filter.dateRange.lte = toDate
    }

    const queryParam: string = JSON.stringify(filter);

    fetch(`http://localhost:8080/public/picture/get?filter=${queryParam}`)
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
    setFilter(newFilter);
  };

  return (
    <main style={{ backgroundColor:"white" }}>
      <Wrap m={5}>
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
            <Text>Select from date: </Text>
            <SingleDatepicker
              date={fromDate}
              onDateChange={setFromDate}
            />
          </HStack>
          <HStack>
            <Text>Select to date: </Text>
            <SingleDatepicker
              date={toDate}
              onDateChange={setToDate}
            />
          </HStack>
        </WrapItem>
      </Wrap>
      <HStack m={5}>
        <Text>Hashtags:</Text>
        <InputTag tags={hashtags} setTags={setHashtags} />
      </HStack>
      <HStack m={5}>
        <Button onClick={getPictures}>Apply filter</Button>
      </HStack>
      <Wrap>
        {pictures !== null
          ? pictures.map((image) => {
              const timestamp = new Date(image.uploadDateTime);
              return (
                <WrapItem key={image.id}>
                  <Image
                    id={image.id}
                    title={image.title}
                    description={image.description}
                    url={image.pictureURI}
                    dateTime={timestamp.toLocaleString()}
                    hashtags={image.hashtags}
                  />
                </WrapItem>
              );
            })
          : null}
      </Wrap>
    </main>
  );
}

// @ts-nocheck
import { useState, useEffect, useRef } from "react";
import { Box, Flex, Tag, TagCloseButton } from "@chakra-ui/react";
import styles from "./input.module.css";

export default function InputTag({tags, setTags}: {tags: string[], setTags: any}) {
  const [sizeInput, setSizeInput] = useState(() => 1);
  const ref_input = useRef<any>(null);

  useEffect(() => {
    ref_input.current.focus(); // auto focus input
    function handleKeyUp(event:any) {
      const newText = ref_input.current.value.trim().replace(",", "");
      switch (event.key) {
        case ",":
          if (newText.length > 0) {
            const dataInputTemp = [...tags];
            dataInputTemp.push(newText);
            setTags(() => [...dataInputTemp]);
            ref_input.current.value = "";
          } else {
            ref_input.current.value = "";
          }
          break;
        case "Enter":
          if (newText.length > 0) {
            const dataInputTemp = [...tags];
            dataInputTemp.push(newText);
            setTags(() => [...dataInputTemp]);
            ref_input.current.value = "";
          }
          break;
        case "Backspace":
          if (tags.length > 0 && newText.length === 0) {
            const dataInputTemp = [...tags];
            dataInputTemp.pop();
            setTags([...dataInputTemp]);
          }
          break;
        default:
          break;
      }
    }
    window.addEventListener("keyup", handleKeyUp);
    return () => window.removeEventListener("keyup", handleKeyUp);
  }, [sizeInput, tags]);

  const handleChangeInput = (e: any) => {
    let value = e.target.value;
    if (value.trim().length > 0) {
      setSizeInput(value.length);
    } else {
      ref_input.current.value = "";
    }
  };
  function handleDelItem(index: number) {
    const dataInputTemp = [...tags];
    dataInputTemp.splice(index, 1);
    setTags(() => [...dataInputTemp]);
  }
  return (
  
    <div className={styles.wrap}>
      <Flex align="center" onClick={() => ref_input.current.focus()}>
        <Box>
          {tags.map((text, i) => (
            <Tag
              key={i + "_" + text}
              colorScheme="cyan"
              className={styles.item_text}
            >
              {text}
              <TagCloseButton onClick={() => handleDelItem(i)} />
            </Tag>
          ))}
          <input
            ref={ref_input}
            onChange={handleChangeInput}
            className={styles.input}
            size={sizeInput}
          />
        </Box>
      </Flex>
    </div>
  );
}

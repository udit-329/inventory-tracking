import {
    Box,
    Button,
    Flex,
    Text,
    useColorModeValue,
    useBreakpointValue,
  } from "@chakra-ui/react";
  
  import DesktopNav from "./DesktopNav";
  import FileSaver from 'file-saver';
  
  const exportData = async (url: string, method: string) => {
    fetch(url, { 
      method: method,
    }).then(function(response) {
      return response.blob();
    }).then(function(blob) {
      FileSaver.saveAs(blob, 'data.csv');
    })
  }

  export default function Navbar() {  
    return (
      <Box>
        <Flex
          bg={useColorModeValue("white", "gray.800")}
          color={useColorModeValue("gray.600", "white")}
          minH="60px"
          py={{ base: 2 }}
          px={{ base: 4 }}
          borderBottom={1}
          borderStyle="solid"
          borderColor={useColorModeValue("gray.200", "gray.900")}
          align="center"
        >
          <Flex flex={{ base: 1 }} justify={{ base: "center", md: "start" }}>
            <Text
              textAlign={useBreakpointValue({ base: "center", md: "left" })}
              fontFamily="heading"
              fontSize="xl"
              fontWeight={800}
              color="green"
            >
              InvTrack
            </Text>    
            <Flex display={{ base: "none", md: "flex" }} ml={10}>
              <DesktopNav />
            </Flex>
          </Flex>   
          <Button colorScheme='green' variant='solid' onClick={() => exportData(`https://inventory-tracking-uk.herokuapp.com/export`, "GET")}>
              Export
            </Button>  
        </Flex>
      </Box>
    );
  }
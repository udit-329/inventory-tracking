import React, { useEffect, useState } from 'react';
import Link from 'next/link'
import { Box, Button } from "@chakra-ui/react";

const getItems = async () => {
  const url = "https://inventory-tracking-uk.herokuapp.com/get";
  const data = await (
      await fetch(url, {
        method: "GET",
        mode: "cors",
      })
    ).json();
  return (data);
};

const onClick = async (url: string, method: string, setData: any) => {
  const response = await (
    await fetch(url, {
      method: method,
    })
  )
  setData(await getItems())
}



const Home = () => {

  const [data, setData] = useState([]);

  useEffect(() => {
    (async () => setData(await getItems()))();
  }, []);

  return (
    <div>
      {data.map((properties) => (
        <Box key={properties["ID"]} minW='sm' maxW='sm' borderWidth='1px' rounded='md' borderRadius='lg' float='left' marginLeft='2%' marginTop='5px' overflow='hidden'>
          <Box p='6'>
            <Box
              mt='1'
              fontWeight='semibold'
              as='h4'
            >
              {properties["Name"]}
            </Box>
            
            <Box color='gray.900' fontSize='m'>
              {`Quantity: ${properties["Quantity"]}`}
            </Box>
            <Box color='gray.900' fontSize='m'>
            {`Location: ${properties["Location"]}`}
            </Box>
            <Box
              color='gray.500'
              letterSpacing='wide'
              fontSize='xs'
              textTransform='uppercase'
            >
              {`product id: ${properties["ID"]}`}
            </Box>
            <div>
                <Link href={`update/${properties["ID"]}`}>
                  <Button colorScheme='green' variant='solid'>
                    Update
                  </Button>
                </Link>

                <Button colorScheme='red' marginLeft='2%' variant='solid' onClick={() => onClick(`https://inventory-tracking-uk.herokuapp.com/delete/${properties["ID"]}`, "DELETE", setData)}>
                  Delete
                </Button>
            </div>
          </Box>
        </Box>
      ))}
    </div>
  );
};

export default Home;

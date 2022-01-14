import {
    Box,
    Stack,
    Link,
    useColorModeValue,
  } from "@chakra-ui/react";
  
  import NAV_ITEMS from "./navItems";
  

  const DesktopNav = () => {
    const linkColor = useColorModeValue("gray.600", "gray.200");
    const linkHoverColor = useColorModeValue("gray.800", "white");
  
    return (
      <Stack direction="row" spacing={4}>
        {NAV_ITEMS.map((navItem) => (
          <Box key={navItem.label}>
            <Link
                p={2}
                href={navItem.href}
                fontSize="sm"
                fontWeight={500}
                color={linkColor}
                _hover={{
                textDecoration: "none",
                color: linkHoverColor,
                }}
            >
                {navItem.label}
            </Link>
          </Box>
        ))}
      </Stack>
    );
  };
  
  export default DesktopNav;
  
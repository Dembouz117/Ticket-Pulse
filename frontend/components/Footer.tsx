import { Box } from "@chakra-ui/react";

const Footer = () => {
    return (
        <Box as="footer" borderTop="1px solid" borderColor="gray.400" p={4} fontSize={"sm"} fontWeight={"thin"} textColor={"gray.400"} backgroundColor={"black"}>
            <p>© 2023 Ticket Pulse. All rights reserved.</p>
        </Box>
    );
};

export default Footer;

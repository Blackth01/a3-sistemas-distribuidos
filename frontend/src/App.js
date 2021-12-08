import { useState } from "react";
import { Box, Button, ButtonGroup, Container, Typography } from "@mui/material";
import Products from "./pages/Products";

function App() {
  // products | raw_material | inputs
  const [currentPage, setCurrentPage] = useState(null);

  const handleChangePage = (page) => {
    setCurrentPage(page);
  };

  return (
    <Container>
      <Typography variant="h2" textAlign="center" marginTop={2}>
        Gerenciamento de produtos
      </Typography>
      <Box
        sx={{ display: "flex" }}
        alignItems="center"
        justifyContent="center"
        marginY={5}
      >
        <ButtonGroup
          variant="contained"
          aria-label="outlined primary button group"
        >
          <Button onClick={() => handleChangePage("products")}>Produtos</Button>
          <Button onClick={() => handleChangePage("raw_material")}>
            Matéria Prima
          </Button>
          <Button onClick={() => handleChangePage("inputs")}>Insumos</Button>
        </ButtonGroup>
      </Box>
      <Box>{currentPage === "products" && <Products />}</Box>
    </Container>
  );
}

export default App;
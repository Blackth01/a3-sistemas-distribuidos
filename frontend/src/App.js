import { useState } from "react";
import { Box, Button, ButtonGroup, Container, Typography } from "@mui/material";
import Products from "./pages/Products";
import RawMaterials from "./pages/RawMaterial";
import Inputs from "./pages/Inputs";
import Reports from "./pages/Reports";

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
          <Button onClick={() => handleChangePage("report_all")}>Relatório total</Button>
          <Button onClick={() => handleChangePage("report_valuable")}>Relatório mais valiosos</Button>
        </ButtonGroup>
      </Box>
      <Box>
        {(() => {
          if (currentPage === "products") {
            return <Products />;
          } else if (currentPage === "raw_material") {
            return <RawMaterials />;
          } else if (currentPage === "inputs"){
            return <Inputs />;
          } else if(currentPage === "report_all"){
            return <Reports reportType={"all"} />;
          } else if(currentPage === "report_valuable"){
            return <Reports reportType={"valuable"} />;
          }
        })()}
      </Box>
    </Container>
  );
}

export default App;

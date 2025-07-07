import axios from "axios";

const DEFAULT_URL = "http://localhost:8080";

type Product = {name: string};

const getProduct = async (id: number): Promise<Product | undefined> => {
  const result = await axios(DEFAULT_URL + `/products/${id}`);
  if (result.status == 200) {
    return new Promise(result.data);
  }
};

export { getProduct };

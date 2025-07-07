import axios from "axios";

const DEFAULT_URL = "http://localhost:8080";

enum ProductCategory {
  FRUIT,
  VEGETABLE,
  MEAT,
  FISH,
  DAIRY,
  SAUCE,
  GRAIN,
  OIL,
  BEVERAGE,
  SNACK,
  SPICE,
}

type Product = {
  id: number;
  category: ProductCategory;
  name: string;
  description: string;
  calories: number;
  fats: number;
  protein: number;
  carbs: number;
};

const getProduct = async (id: number): Promise<Product | undefined> => {
  const result = await axios(DEFAULT_URL + `/products/${id}`);
  if (result.status == 200) {
    return new Promise(result.data);
  }
};

export { getProduct };

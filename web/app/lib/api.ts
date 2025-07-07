import axios from "axios";

const DEFAULT_URL = "http://localhost:8080";

enum ProductCategory {
  FRUIT = "fruit",
  VEGETABLE = "vegetable",
  MEAT = "meat",
  FISH = "fish",
  DAIRY = "dairy",
  SAUCE = "sauce",
  GRAIN = "grain",
  OIL = "oil",
  BEVERAGE = "beverage",
  SNACK = "snack",
  SPICE = "spice",
}

enum DishCategory {
  SALAD = "salad",
  SOUP = "soup",
  MAIN = "main",
  DESSERT = "dessert",
  SANDWICH = "sandwich",
  WRAP = "wrap",
  PIZZA = "pizza",
  PASTA = "pasta",
  BURGER = "burger",
  BREAKFAST = "breakfast",
}

export type Product = {
  id: number;
  category: ProductCategory;
  name: string;
  description: string;
  calories: number;
  fats: number;
  protein: number;
  carbs: number;
};

export type Ingredient = {
  product: Product;
  weight: number;
};

export type Dish = {
  id: number;
  category: DishCategory;
  name: string;
  description: string;
  ingredients: Ingredient[];
};

const getProductsFromFilter = async (filter: string): Promise<Product[] | undefined> => {
  const { data, status } = await axios<Product[]>(DEFAULT_URL + `/products?filter=${filter}`);
  if (status == 404) {
    return;
  }
  return data;
}

const getProduct = async (id: number): Promise<Product | undefined> => {
  const { data, status } = await axios(DEFAULT_URL + `/product/${id}`);
  if (status == 400) {
    return;
  }
  return data;
};

const addProduct = async (pr: Product): Promise<number | undefined> => {
  const {data, status} = await axios.post<Product>(DEFAULT_URL + "/dishes", pr)
  if (status == 400) {
    return 
  }
  return data.id
};

const getDish = async (id: number): Promise<Dish | undefined> => {
  const { data, status } = await axios(DEFAULT_URL + `/dish/${id}`);
  if (status == 400) {
    return;
  }
  return data;
};

export { getProduct, getProductsFromFilter, getDish, addProduct, ProductCategory, DishCategory };

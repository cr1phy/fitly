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

export enum ProductType {
  RAW = "raw",
  READY = "ready",
  SEMI = "semi",
  COMPOSITE = "composite",
}

export enum Unit {
  GRAM = "g",
  KILOGRAM = "kg",
  LITER = "l",
  MILLILITER = "ml",
  PIECE = "pcs",
  TABLESPOON = "tbsp",
  TEASPOON = "tsp",
  CUP = "cup",
  PACKAGE = "pack",
  BOTTLE = "bottle",
  CAN = "can",
  SLICE = "slice",
  BUNCH = "bunch",
  CLOVE = "clove",
}

export type Product = {
  id: number;
  name: string;
  description: string;
  category: ProductCategory;
  type: ProductType;
  brand: string;
  image_url: string | null;
  calories: number;
  fats: number;
  protein: number;
  carbs: number;
  is_organic: boolean;
  is_vegetarian: boolean;
  is_vegan: boolean;
  is_gluten_free: boolean;
};

export type ProductAlias = {
  id: number;
  product_id: number;
  alias: string;
};

export type NutritionInfo = {
  calories: number;
  fats: number;
  protein: number;
  carbs: number;
  weight: number;
};

export type Ingredient = {
  id?: number;
  dish_id?: number;
  product_id: number;
  product?: Product;
  weight: number;
  unit?: Unit;
};

export type Dish = {
  id: number;
  category: DishCategory;
  name: string;
  description: string;
  ingredients: Ingredient[];
};

const getProductsFromFilter = (filter: string): Promise<Product[]> =>
  axios<Product[]>(DEFAULT_URL + `/products?filter=${filter}`).then(
    (p) => p.data
  );

const getProduct = async (id: number): Promise<Product | undefined> =>
  axios<Product>(DEFAULT_URL + `/product/${id}`).then((p) => {
    if (p) return p.data
  });

const addProduct = async (pr: Product): Promise<number | undefined> => {
  const { data, status } = await axios.post<Product>(
    DEFAULT_URL + "/dishes",
    pr
  );
  if (status == 400) {
    return;
  }
  return data.id;
};

const getDish = async (id: number): Promise<Dish | undefined> => {
  const { data, status } = await axios(DEFAULT_URL + `/dish/${id}`);
  if (status == 400) {
    return;
  }
  return data;
};

export {
  getProduct,
  getProductsFromFilter,
  getDish,
  addProduct,
  ProductCategory,
  DishCategory,
};
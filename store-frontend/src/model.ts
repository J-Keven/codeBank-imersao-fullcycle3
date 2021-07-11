export interface Products {
  id: string;
  name: string;
  description: string;
  image_url: string;
  slug: string;
  price: number;
  created_at: string;
}

export const products: Products[] = [
  {
    id: "1",
    name: "Product Test",
    description: "este é um produto teste",
    price: 100.73,
    image_url: "https://source.unsplash.com/random?product," + Math.random(),
    slug: "product_test",
    created_at: "2021-07-12T00:00:00"
  },
  {
    id: "1",
    name: "Product Test",
    description: "este é um produto teste",
    price: 100.73,
    image_url: "https://source.unsplash.com/random?product," + Math.random(),
    slug: "product_test",
    created_at: "2021-07-12T00:00:00"
  },
  {
    id: "2",
    name: "Product Test",
    description: "este é um produto teste",
    price: 100.73,
    image_url: "https://source.unsplash.com/random?product," + Math.random(),
    slug: "product_test",
    created_at: "2021-07-12T00:00:00"
  },
  {
    id: "3",
    name: "Product Test",
    description: "este é um produto teste",
    price: 100.73,
    image_url: "https://source.unsplash.com/random?product," + Math.random(),
    slug: "product_test",
    created_at: "2021-07-12T00:00:00"
  },
  {
    id: "4",
    name: "Product Test",
    description: "este é um produto teste",
    price: 100.73,
    image_url: "https://source.unsplash.com/random?product," + Math.random(),
    slug: "product_test",
    created_at: "2021-07-12T00:00:00"
  },
  {
    id: "5",
    name: "Product Test",
    description: "este é um produto teste",
    price: 100.73,
    image_url: "https://source.unsplash.com/random?product," + Math.random(),
    slug: "product_test",
    created_at: "2021-07-12T00:00:00"
  },

]
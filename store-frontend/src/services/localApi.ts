import axios, { AxiosInstance } from "axios"
import { Products } from "../model"


class LocalApi {
  private client: AxiosInstance
  constructor() {
    this.client = axios.create({
      baseURL: "http://localhost:3000/api"
    })
  }

  async getProducts(): Promise<Products[]> {
    const products = await this.client.get<Products[]>("/products")
    return products.data
  }
}

export default new LocalApi()
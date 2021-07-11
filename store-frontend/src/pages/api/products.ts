import type { NextApiRequest, NextApiResponse } from 'next'
import { Products, products } from '../../model'

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Products[]>
) {
  res.status(200).json(products)
}

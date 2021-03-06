import { Button, Card, CardActions, CardContent, CardHeader, CardMedia, Typography } from '@material-ui/core'
import { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import { Products, products } from '../../model'

interface ProductDetailPageProps {
  product: Products
}

// const ProductDetailPage: NextPage<ProductDetailPageProps> = ({ product }) => {
const ProductDetailPage = () => {
  const product = products[0]
  return (
    <div>
      <Head>
        <title>{product.name} - Detalhes do produto</title>
        {/* <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" /> */}
      </Head>
      <Card>
        <CardHeader title={product.name.toUpperCase()} subheader={`R$ ${product.price}`} />
        <CardActions>
          <Button size="small" component="a" color="primary">Comprar</Button>
        </CardActions>
        <CardMedia style={{ paddingTop: "56%" }} image={product.image_url} />
        <CardContent>
          <Typography component="p" variant="body2" gutterBottom color="textSecondary">
            {product.description}
          </Typography>
        </CardContent>
      </Card>

    </div>
  )
}

export default ProductDetailPage


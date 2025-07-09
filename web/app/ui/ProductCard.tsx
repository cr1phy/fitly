import { Card } from "@/components/ui/card"
import Image from "next/image"
import Link from "next/link"

type ProductCardProps = { id: number, title: string, description: string, imageUrl: string }

const ProductCard = ({ props }: { props: ProductCardProps }) => {
    return (
        <Link href={`/product/${props.id}`}>
            <Card className="flex">
                <Image src={props.imageUrl} alt={props.title} />
                <div>
                    <h3>{props.title}</h3>
                    <p>{props.description}</p>
                </div>
            </Card>
        </Link>
    )
}

export default ProductCard;
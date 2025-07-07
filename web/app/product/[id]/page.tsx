import { getProduct } from "@/app/lib/api"

export default async function Page({
    params,
}: {
    params: Promise<{ id: number }>
}) {
    const { id } = await params
    const product = await getProduct(id)

    return <h1>{product?.name && "Not found"}</h1>
}
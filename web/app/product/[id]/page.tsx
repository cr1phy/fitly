import { getProduct } from "@/lib/api"

export default async function Page({
    params,
}: {
    params: Promise<{ id: number }>
}) {
    const { id } = await params
    try {
        const product = await getProduct(id)
        if (!product) {
            return <h1>Not found</h1>
        }
        return <h1>{JSON.stringify(product)}</h1>
    } catch (error) {
        return <h1>Not found</h1>
    }
}
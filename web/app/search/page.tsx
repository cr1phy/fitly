"use client"

import { useEffect, useState } from "react";
import { getProductsFromFilter, Product } from "../lib/api";

export default function Page() {
    const [filter, setFilter] = useState<string>("")
    const [products, setProducts] = useState<Product[]>([])

    useEffect(() => {
        getProductsFromFilter(filter)
            .then(p => setProducts(p!!))
    }, [filter])

    return (
        <main className="flex flex-col items-center justify-center min-h-screen px-4 py-8 text-center">
            <h1 className="text-2xl md:text-4xl font-bold mb-6">Search something, what you want.</h1>
            <input
                onChange={(e) => setFilter(e.target.value)}
                className="w-full max-w-md rounded-md px-4 py-2 dark:bg-slate-700 bg-slate-300 focus:outline-none focus:ring-2 focus:ring-blue-400 text-base md:text-lg"
                placeholder="Search..."
                wfd-id="id0"
            />
            {filter !== "" && <div>
                {products.map((p) => (
                    <div key={p.id}>
                        <h1>{p.name}</h1>
                    </div>
                )) && "Not found"}
            </div>}
        </main>
    );
}
"use client"

import { useEffect, useState } from "react";
import { getProductsFromFilter, Product } from "../../lib/api";
import { Input } from "@/components/ui/input";
import ProductCard from "../ui/ProductCard";
import { useQuery } from "@tanstack/react-query";
import useDebounce from "../../lib/useDebounce";

export default function Page() {
    const [filter, setFilter] = useState<string>("")
    const debouncedFilter = useDebounce({ value: filter, delayMillis: 1000 })
    const { data, isSuccess } = useQuery({ queryKey: ["products"], queryFn: () => { if (filter !== "") return getProductsFromFilter(debouncedFilter) } })

    return (
        <main className="flex flex-col items-center justify-center min-h-screen px-4 py-8 text-center">
            <h1 className="text-2xl md:text-4xl font-bold mb-6">Search something, what you want.</h1>
            <Input
                onChange={(e) => setFilter(e.target.value.trim())}
                className="w-full max-w-md rounded-md px-4 py-2 dark:bg-slate-700 bg-slate-300 focus:outline-none focus:ring-2 focus:ring-blue-400 text-base md:text-lg"
                placeholder="Search..."
                wfd-id="id0"
            />
            {<div>
                {isSuccess ? data?.map((p) => (
                    <ProductCard key={p.id} props={{ id: p.id, title: p.name, description: p.description, imageUrl: "" }} />
                )) : "Not found."}
            </div>}
        </main>
    );
}
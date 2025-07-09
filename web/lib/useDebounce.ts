import { useEffect, useState } from "react";

interface UseDebounceOptions<T> {
    value: T;
    delayMillis: number;
}

const useDebounce = <T>({ value, delayMillis }: UseDebounceOptions<T>): T => {
    const [debouncedValue, setDebouncedValue] = useState<T>(value);

    useEffect(() => {
        const t: ReturnType<typeof setTimeout> = setTimeout(() => {
            setDebouncedValue(value);
        }, delayMillis);

        return () => clearTimeout(t);
    }, [value, delayMillis]);
    return debouncedValue;
};

export default useDebounce;

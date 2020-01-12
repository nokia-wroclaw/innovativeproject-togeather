import { Product } from './product';

export interface Restaurant {
    id: number;
    name: string;
    address: string;
    readonly menu: Product[];
}

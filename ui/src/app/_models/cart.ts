import { Product } from './product';

export interface Cart {
    products: Array<Partial<Product>>;
    cartTotal: number;
    deliveryCost: number;
    numberOfMembers: number;
}

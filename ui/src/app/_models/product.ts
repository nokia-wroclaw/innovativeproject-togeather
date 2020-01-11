/*
    Product is a model that represents cart item and restaurants menu element (dish).
*/
export interface Product {
    readonly id: number;
    readonly restaurantId: number;
    readonly name: string;
    readonly price: number;
    readonly description: string;
}

import { Injectable } from '@angular/core';
import { Product } from '../_models/product';

@Injectable({
    providedIn: 'root'
})
export class CartService {
    items: Product[] = [];

    // TODO: Should be replaced with websocket command "addToCart"
    addToCart(item: Product): void {
        this.items.push(item);
    }

    deleteFromCart(item: Product): void {
        this.items.splice(
            this.items.indexOf(item),
            1
        );
    }

    // TODO: Should be replaced with websocket command "getNewCart"(?)
    // or handle websocket response
    getItems(): Product[] {
        return this.items;
    }

    clearCart(): Product[] {
        this.items = [];
        return this.items;
    }
}

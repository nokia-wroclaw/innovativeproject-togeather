import { Injectable } from '@angular/core';
import { Product } from '../_models/product';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { ApiService } from './api.service';
import { Cart } from '../_models/cart';
import { CartDto } from '../_models/cart-dto';

@Injectable({
    providedIn: 'root'
})
export class CartService {

    readonly apiBaseUrl = environment.apiUrl;

    constructor(
        private http: HttpClient,
    ) { }

    static mapDtoToCart(cart: CartDto): Cart {
        return {
            products: cart.meals.map(meal => {
                return {
                    id: meal.id,
                    name: meal.name,
                    price: meal.price,
                };
            }),
            cartTotal: cart.cart_value,
            deliveryCost: cart.delivery_cost,
            numberOfMembers: cart.lobby_count,
        };
    }

    addToCart(item: Product, lobbyId: number): Observable<Cart> {
        return this.http.post<CartDto>(
            this.apiBaseUrl + `/lobbies/${lobbyId}/order`,
            { meal_id: item.id },
            { withCredentials: true },
        ).pipe(
            map(CartService.mapDtoToCart),
            catchError(ApiService.handleError)
        );
    }

    deleteFromCart(item: Partial<Product>, lobbyId: number): Observable<Cart> {
        return this.http.delete<CartDto>(
            this.apiBaseUrl + `/lobbies/${lobbyId}/order`,
            {
                withCredentials: true,
                params: { meal_id: item.id.toString() },
            }
        ).pipe(
            map(CartService.mapDtoToCart),
            catchError(ApiService.handleError)
        );
    }
}

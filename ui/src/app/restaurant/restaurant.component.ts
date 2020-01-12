import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from '../_services/api.service';
import { Restaurant } from '../_models/restaurant';
import { Product } from '../_models/product';
import { CartService } from '../_services/cart.service';

@Component({
    selector: 'app-restaurant',
    templateUrl: './restaurant.component.html',
    styleUrls: ['./restaurant.component.scss']
})
export class RestaurantComponent implements OnInit {

    restaurant: Restaurant;

    constructor(
        private route: ActivatedRoute,
        private apiService: ApiService,
        private cartService: CartService,
    ) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe(
            params => {
                this.apiService.getSingleRestaurant(params.get('restaurantId')).subscribe(
                    restaurant => {
                        this.restaurant = restaurant;
                    }
                );
            },
            error => {
                console.error(`Could not fetch single restaurant`, error);
            }
        );
    }

    addToCart(dish: Product) {
        this.cartService.addToCart(dish);
    }
}

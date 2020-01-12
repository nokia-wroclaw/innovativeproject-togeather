import { Component, Input, OnInit } from '@angular/core';
import { ApiService } from '../_services/api.service';
import { Restaurant } from '../_models/restaurant';
import { Product } from '../_models/product';
import { CartService } from '../_services/cart.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-restaurant',
  templateUrl: './restaurant.component.html',
  styleUrls: ['./restaurant.component.scss']
})
export class RestaurantComponent implements OnInit {

  @Input() restaurantId: string;
  restaurant$: Observable<Restaurant>;

  constructor(
    private apiService: ApiService,
    private cartService: CartService,
  ) {
  }

  ngOnInit() {
    this.restaurant$ = this.apiService.getRestaurant(this.restaurantId);
  }

  addToCart(dish: Product) {
    this.cartService.addToCart(dish);
  }
}

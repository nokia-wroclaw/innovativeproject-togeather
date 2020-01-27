import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ApiService } from '../_services/api.service';
import { Restaurant } from '../_models/restaurant';
import { Product } from '../_models/product';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-restaurant',
  templateUrl: './restaurant.component.html',
  styleUrls: ['./restaurant.component.scss']
})
export class RestaurantComponent implements OnInit {

  @Input() restaurantId: string;
  @Input() showAddToCart: boolean = false;
  @Output() orderProduct = new EventEmitter<Product>();
  restaurant$: Observable<Restaurant>;

  constructor(
    private apiService: ApiService,
  ) {
  }

  ngOnInit() {
    this.restaurant$ = this.apiService.getRestaurant(this.restaurantId);
  }

  addToCart(dish: Product) {
    this.orderProduct.emit(dish);
  }
}

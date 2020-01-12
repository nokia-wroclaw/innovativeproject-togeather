import { Component, OnInit } from '@angular/core';
import { CartService } from '../_services/cart.service';
import { Product } from '../_models/product';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.scss']
})
export class CartComponent implements OnInit {
  items: Product[];

  constructor(
      private cartService: CartService,
  ) { }

  ngOnInit() {
    this.items = this.cartService.getItems();
  }

  deleteItem(item: Product): void {
    this.cartService.deleteFromCart(item);
    this.items = this.cartService.getItems();
  }

  isCartEmpty(): boolean {
    return this.items ? this.items.length === 0 : true;
  }
}

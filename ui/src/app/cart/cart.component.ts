import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { Product } from '../_models/product';
import { Cart } from '../_models/cart';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.scss']
})
export class CartComponent implements OnChanges {

  @Input() cartState: Cart;
  @Output() delete = new EventEmitter<Partial<Product>>();

  ngOnChanges() {
    // TODO: Handle changes of the input
  }

  deleteItem(item: Partial<Product>): void {
    this.delete.emit(item);
  }

  isCartEmpty(): boolean {
    if (this.cartState) {
      return this.cartState.products ? this.cartState.products.length === 0 : true;
    }
    return true;
  }
}

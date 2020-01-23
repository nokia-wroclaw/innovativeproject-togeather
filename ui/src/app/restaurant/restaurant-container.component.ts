import { Component, OnInit } from '@angular/core';
import { pluck } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-restaurant-container',
  templateUrl: './restaurant-container.component.html'
})
export class RestaurantContainerComponent implements OnInit {

  restaurantId$: Observable<string>;

  constructor(
    private route: ActivatedRoute,
  ) { }

  ngOnInit() {
    this.restaurantId$ = this.route.params.pipe(pluck('restaurantId'));
  }

}

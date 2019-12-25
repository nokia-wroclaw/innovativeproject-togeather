import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { RestaurantsListComponent } from './restaurants-list/restaurants-list.component';
import { RestaurantComponent } from './restaurant/restaurant.component';
import { CreateLobbyComponent } from './create-lobby/create-lobby.component';


const routes: Routes = [
  { path: '', component: LandingPageComponent },
  { path: 'restaurants', component: RestaurantsListComponent },
  { path: 'restaurants/:restaurantId', component: RestaurantComponent },
  { path: 'create-lobby', component: CreateLobbyComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

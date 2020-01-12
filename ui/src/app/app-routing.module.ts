import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { RestaurantsListComponent } from './restaurants-list/restaurants-list.component';
import { CreateLobbyComponent } from './create-lobby/create-lobby.component';
import { LobbiesListComponent } from './lobbies-list/lobbies-list.component';
import { LobbyComponent } from './lobby/lobby.component';
import { RestaurantContainerComponent } from './restaurant/restaurant-container.component';


const routes: Routes = [
    { path: '', component: LandingPageComponent },
    { path: 'restaurants', component: RestaurantsListComponent },
    { path: 'restaurants/:restaurantId', component: RestaurantContainerComponent },
    { path: 'create-lobby', component: CreateLobbyComponent },
    { path: 'open-lobbies', component: LobbiesListComponent },
    { path: 'lobbies/:lobbyId', component: LobbyComponent },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }

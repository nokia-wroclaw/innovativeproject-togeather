import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Lobby } from '../_models/lobby';

@Injectable({
  providedIn: 'root'
})
export class RedirectionService {

  constructor(
    private router: Router,
  ) {
  }

  static handleError(viewName: string, error: any): void {
    console.error(`Error when redirecting to ${viewName} view:`, error);
  }

  redirectToSingleRestaurant(id: number): void {
    this.router.navigateByUrl(`/restaurants/${ id }`)
      .catch(error => {
        RedirectionService.handleError('single restaurant', error);
      });
  }

  redirectToHomePage(): void {
    this.router.navigateByUrl('/')
      .catch(error => {
        RedirectionService.handleError('home page', error);
      });
  }

  redirectToLobbyCreation(): void {
    this.router.navigateByUrl('/create-lobby')
      .catch(error => {
        RedirectionService.handleError('lobby creation', error);
      });
  }

  redirectToRestaurants(): void {
    this.router.navigateByUrl('/restaurants')
      .catch(error => {
        RedirectionService.handleError('restaurants', error);
      });
  }

  redirectToLobbies(): void {
    this.router.navigateByUrl('/open-lobbies')
      .catch(error => {
        RedirectionService.handleError('lobbies', error);
      });
  }

  redirectToSingleLobby(lobbyId: number): void {
    this.router.navigate([`/lobbies/${lobbyId}`], { state: { lobbyId: lobbyId } })
        .catch(error => {
          RedirectionService.handleError('lobby', error);
        });
  }
}

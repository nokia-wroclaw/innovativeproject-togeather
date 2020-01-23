import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Injectable({
  providedIn: 'root'
})
export class RedirectionService {

  constructor(
    private router: Router,
    private toaster: ToastrService,
  ) {
  }

  handleError(viewName: string, error: any): void {
    this.toaster.error(error, `Error when redirecting to ${viewName} view:`);
  }

  redirectToSingleRestaurant(id: number): void {
    this.router.navigateByUrl(`/restaurants/${ id }`)
      .catch(error => {
        this.handleError('single restaurant', error);
      });
  }

  redirectToHomePage(): void {
    this.router.navigateByUrl('/')
      .catch(error => {
        this.handleError('home page', error);
      });
  }

  redirectToLobbyCreation(): void {
    this.router.navigateByUrl('/create-lobby')
      .catch(error => {
        this.handleError('lobby creation', error);
      });
  }

  redirectToRestaurants(): void {
    this.router.navigateByUrl('/restaurants')
      .catch(error => {
        this.handleError('restaurants', error);
      });
  }

  redirectToLobbies(): void {
    this.router.navigateByUrl('/open-lobbies')
      .catch(error => {
        this.handleError('lobbies', error);
      });
  }

  redirectToSingleLobby(lobbyId: number): void {
    this.router.navigate([`/lobbies/${lobbyId}`], { state: { lobbyId: lobbyId } })
        .catch(error => {
          this.handleError('lobby', error);
        });
  }
}

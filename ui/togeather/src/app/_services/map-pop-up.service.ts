import { Injectable } from '@angular/core';
import { Lobby } from '../_models/lobby';

@Injectable({
  providedIn: 'root'
})
export class MapPopUpService {

  constructor() { }

  // TODO: Change this when the new lobby structure is finally given from api
  static makeLobbyPopup(data: Lobby): string {
    return '' + '<div>Restaurant: ' + data.restaurant.name + '</div>'
              + '<div>Address lobby: ' + '(no address given)' + '</div>'
              + '<div>Expiration Date: ' + data.expires + '</div>';

  }
}

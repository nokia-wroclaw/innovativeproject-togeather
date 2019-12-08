import { Injectable } from '@angular/core';
import { Lobby } from '../_models/lobby';

@Injectable({
  providedIn: 'root'
})
export class MapPopUpService {

  constructor() { }

  makeLobbyPopup(data: Lobby): string {
    return '' + '<div>Restaurant: ' + data.restaurant + '</div>'
              + '<div>Address lobby: ' + data.addressLobby + '</div>'
              + '<div>Expiration Date: ' + data.expirationDate + '</div>';
    
  }
}

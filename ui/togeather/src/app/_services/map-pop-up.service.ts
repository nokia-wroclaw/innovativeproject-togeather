import { Injectable } from '@angular/core';
import { Lobby } from '../_models/lobby';

@Injectable({
  providedIn: 'root'
})
export class MapPopUpService {

  constructor() { }

  static makeLobbyPopup(data: Lobby): string {
    const addressArray = data.address.replace(',', '').split(' ');
    const address = `${ addressArray[1] } ${ addressArray[0] }, ${ addressArray[2] }`;

    return '' + '<div>Restaurant: ' + data.restaurant.name + '</div>'
              + '<div>Address lobby: ' + address + '</div>'
              + '<div>Expiration Date: ' + data.expires + '</div>';

  }
}

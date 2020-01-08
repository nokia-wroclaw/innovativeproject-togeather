import { Injectable } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { BeautifyAddressPipe } from '../_pipes/beautify-address.pipe';

@Injectable({
  providedIn: 'root'
})
export class MapPopUpService {

  constructor() { }

  static makeLobbyPopup(data: Lobby): string {
    const beautifyAddressPipe = new BeautifyAddressPipe();

    return '' + '<div>Restaurant: ' + data.restaurant.name + '</div>'
              + '<div>Address lobby: ' + beautifyAddressPipe.transform(data.address) + '</div>'
              + '<div>Expiration Date: ' + data.expires + '</div>';

  }
}

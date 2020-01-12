import { Injectable } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { BeautifyAddressPipe } from '../_pipes/beautify-address.pipe';
import { BeautifyExpirationDatePipe } from '../_pipes/beautify-expiration-date.pipe';

@Injectable({
  providedIn: 'root'
})
export class MapPopUpService {

  constructor() { }

  static makeLobbyPopup(data: Lobby): string {
    const beautifyAddressPipe = new BeautifyAddressPipe();
    const expiresPipe = new BeautifyExpirationDatePipe();

    return '' + '<div>Restaurant: ' + data.restaurant.name + '</div>'
              + '<div>Address lobby: ' + beautifyAddressPipe.transform(data.address) + '</div>'
              + '<div>Expires in: ' + expiresPipe.transform(data.expires) + '</div>';

  }
}

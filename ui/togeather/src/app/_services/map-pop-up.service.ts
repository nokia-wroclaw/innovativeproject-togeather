import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MapPopUpService {

  constructor() { }

  makeLobbyPopup(data: any): string {
    return '' + '<div>Address: ' + data.address +'</div>';
    
  }
}

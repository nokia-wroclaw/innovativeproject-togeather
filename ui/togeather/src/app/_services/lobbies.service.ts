import { Injectable } from '@angular/core';
import * as LeafLet from 'leaflet';

@Injectable({
  providedIn: 'root'
})
export class LobbiesService {

  lobbies = [
    {lat:51.101417, lon: 17.036716},
    {lat:51.100919, lon: 17.032328},
  ]
  constructor() { }

  makeLobbiesMarkers(map: LeafLet.map): void {
    for (const c of this.lobbies) {
      const lat_coord = c.lat;
      const lon_coord = c.lon;
      const maker = LeafLet.marker([lat_coord, lon_coord]).addTo(map);
    }
  };

}

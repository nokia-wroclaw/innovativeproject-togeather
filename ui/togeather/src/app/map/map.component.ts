import { Component, Input, OnInit } from '@angular/core';
import * as LeafLet from 'leaflet';
import { LocationService } from '../_services/location.service';
import { LobbiesService } from '../_services/lobbies.service';

export interface Coordinates {
  lat: number;
  lon: number;
}

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss']
})
export class MapComponent implements OnInit {

  @Input() coords: Coordinates;
  private map;

  constructor(
      private locationService: LocationService,
      private lobbiesService: LobbiesService,
  ) { }

  ngOnInit() {
    this.initMap();

    this.locationService.getLocation()
        .then((location: Coordinates) => {
          this.map.panTo(new LeafLet.LatLng(location.lat, location.lon));
        })
        .catch(error => {
          alert(error);
        });
  }

  private initMap(): void {
    this.map = LeafLet.map('map', {
      center: [ 51.12584, 16.97778 ],
      zoom: 15,
    });

    const tiles = LeafLet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      maxZoom: 20,
      attribution: '&copy;<a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    });
    tiles.addTo(this.map);

    this.lobbiesService.makeLobbiesMarkers(this.map);
  }

}

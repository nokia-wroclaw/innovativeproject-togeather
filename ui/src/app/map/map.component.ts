import { AfterViewInit, Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import * as LeafLet from 'leaflet';

export interface Coordinates {
  lat: number;
  lon: number;
}

let id = 0;

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss']
})
export class MapComponent implements AfterViewInit, OnChanges {

  @Input() id: string = `map-${id++}`;
  @Input() coords: Coordinates;
  private map;

  constructor( ) { }

  ngAfterViewInit() {
    this.initMap();
  }

  ngOnChanges(changes: SimpleChanges) {
    if (this.map && changes.coords) {
      const leafletCoords = new LeafLet.LatLng(changes.coords.currentValue.lat, changes.coords.currentValue.lon);
      this.map.panTo(leafletCoords);
    }
  }

  getMap(): ReturnType<LeafLet.map> {
    return this.map;
  }

  private initMap(): void {
    this.map = LeafLet.map(this.id, {
      center: [ 51.12584, 16.97778 ],
      zoom: 12.5,
    });

    if (this.coords) {
      this.map.panTo(new LeafLet.LatLng(this.coords.lat, this.coords.lon));
    }

    const tiles = LeafLet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      maxZoom: 20,
      attribution: '&copy;<a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    });
    tiles.addTo(this.map);
  }

}

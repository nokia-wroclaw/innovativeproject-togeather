import { AfterViewInit, Component, Host, Input, OnChanges, OnDestroy, OnInit, SimpleChanges } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { LobbiesService } from '../_services/lobbies.service';
import { MapComponent } from '../map/map.component';

@Component({
  selector: 'app-map-lobby-marker',
  template: '',
})
export class MapLobbyMarkerComponent implements OnInit, AfterViewInit, OnChanges, OnDestroy {

  @Input() lobby: Lobby;
  private marker;

  constructor(
      private lobbiesService: LobbiesService,
      @Host() private mapComponent: MapComponent,
  ) { }

  ngOnInit() {
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.lobby && this.mapComponent.getMap()) {
      if (this.marker) {
        LobbiesService.removeMarker(this.mapComponent.getMap(), this.marker);
      }
      this.marker = this.lobbiesService.makeLobbyMarker(this.mapComponent.getMap(), this.lobby);
    }
  }

  ngAfterViewInit(): void {
    setTimeout(() => {
      if (this.mapComponent.getMap()) {
        this.marker = this.lobbiesService.makeLobbyMarker(this.mapComponent.getMap(), this.lobby);
      }
    }, 0);
  }

  ngOnDestroy() {
    if (this.marker) {
      LobbiesService.removeMarker(this.mapComponent.getMap(), this.marker);
    }
  }

}

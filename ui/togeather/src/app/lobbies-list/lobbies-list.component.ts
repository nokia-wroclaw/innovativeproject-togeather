import { Component, OnInit } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { ApiService } from '../_services/api.service';
import { Observable } from 'rxjs';
import { RedirectionService } from '../_services/redirection.service';

@Component({
  selector: 'app-lobbies-list',
  templateUrl: './lobbies-list.component.html',
  styleUrls: ['./lobbies-list.component.scss']
})
export class LobbiesListComponent implements OnInit {

  lobbies$: Observable<Lobby[]>;

  constructor(
    private apiService: ApiService,
    private redirectionService: RedirectionService,
  ) { }

  ngOnInit() {
    this.lobbies$ = this.apiService.getLobbies();
  }

  redirectToRestaurant(id: number): void {
    this.redirectionService.redirectToSingleRestaurant(id);
  }

}

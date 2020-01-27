import { Component, OnInit } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { ApiService } from '../_services/api.service';
import { Observable } from 'rxjs';
import { RedirectionService } from '../_services/redirection.service';
import { MatDialog } from '@angular/material/dialog';
import { JoinLobbyComponent } from '../join-lobby/join-lobby.component';
import { ToastrService } from 'ngx-toastr';

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
    private api: ApiService,
    private toaster: ToastrService,
    public dialog: MatDialog,
  ) { }

  ngOnInit() {
    this.lobbies$ = this.apiService.getLobbies();
  }

  redirectToLobby(id: number): void {
    this.redirectionService.redirectToSingleLobby(id);
  }

  openPopup(lobby: Lobby) {
    const dialogRef = this.dialog.open(JoinLobbyComponent, {
      width: '300px',
      data: { lobbyId: lobby.id },
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('Dialog closed');
    });
  }

  onJoinClick(lobbyId: number): void {
    this.api.joinLobby(lobbyId).subscribe(
        id => {
          this.redirectionService.redirectToSingleLobby(id);
        },
        error => {
          this.toaster.error(error);
        },
    );
  }

}

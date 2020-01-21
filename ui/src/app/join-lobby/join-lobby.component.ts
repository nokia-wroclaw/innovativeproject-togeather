import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ApiService } from '../_services/api.service';
import { ToastrService } from 'ngx-toastr';
import { RedirectionService } from '../_services/redirection.service';

@Component({
  selector: 'app-join-lobby',
  templateUrl: './join-lobby.component.html',
})
export class JoinLobbyComponent implements OnInit {

  userName: string;

  constructor(
    private api: ApiService,
    private toaster: ToastrService,
    private redirectionService: RedirectionService,
    public dialogRef: MatDialogRef<JoinLobbyComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { lobbyId: number },
  ) { }

  ngOnInit() {
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  onJoinClick(): void {
    this.api.joinLobby(this.data.lobbyId, this.userName).subscribe(
      () => {
        this.dialogRef.close();
        this.redirectionService.redirectToSingleLobby(this.data.lobbyId);
      },
      error => {
        this.toaster.error(error);
      },
    );
  }

}

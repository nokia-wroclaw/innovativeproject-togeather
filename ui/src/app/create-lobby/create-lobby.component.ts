import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Restaurant } from '../_models/restaurant';
import { ApiService } from '../_services/api.service';
import { Observable, of } from 'rxjs';
import { PostLobbyDto } from '../_models/post-lobby-dto';

@Component({
  selector: 'app-create-lobby',
  templateUrl: './create-lobby.component.html',
  styleUrls: ['./create-lobby.component.scss']
})
export class CreateLobbyComponent implements OnInit {

  disableCreateButton = false;
  expirationTime: { hour: number, minute: number, meriden: 'PM' | 'AM', format: 24 | 12 };
  restaurants$: Observable<Restaurant[]> = of([]);

  lobbyForm = this.fb.group({
    ownerName: ['', Validators.required ],
    street: [ '', Validators.required ],
    nr: [ '', Validators.required ],
    city: [ '', Validators.required ],
    restaurantId: [ null, Validators.required ],
  });

  static sanitize (field: string) {
    return field.trim().replace(',', ' ');
  }

  constructor(
      private api: ApiService,
      private fb: FormBuilder,
  ) { }

  ngOnInit() {
    const today = new Date();
    this.expirationTime = {
      hour: today.getHours(),
      minute: today.getMinutes(),
      meriden: 'PM',
      format: 24
    };

    this.restaurants$ = this.api.getRestaurants();
  }

  createNewLobby(): void {
    if (this.lobbyForm.valid) {
      this.disableCreateButton = true;

      const expirationDate = new Date();
      expirationDate.setUTCHours(this.expirationTime.hour, this.expirationTime.minute);

      const control = this.lobbyForm.controls;
      const address = CreateLobbyComponent.sanitize(control.nr.value)
          + ','
          + CreateLobbyComponent.sanitize(control.street.value)
          + ','
          + CreateLobbyComponent.sanitize(control.city.value);

      const newLobby: PostLobbyDto = {
        restaurant_id: this.lobbyForm.controls['restaurantId'].value,
        owner_name: this.lobbyForm.controls['ownerName'].value,
        expires: expirationDate.toISOString(),
        address: address
      };


      this.api.postLobby(newLobby).subscribe(lobby => {
        this.disableCreateButton = false;
        console.log('Newly created lobby: ');
        console.log(lobby);
      }, () => {
        this.disableCreateButton = false;
      });
    } else {
      console.error('Form is not valid! I won\'t send the request >.<');
    }
  }
}

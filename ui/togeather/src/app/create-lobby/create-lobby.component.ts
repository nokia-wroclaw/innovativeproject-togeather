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

  constructor(
      private api: ApiService,
      private fb: FormBuilder,
  ) { }
  today = new Date();
  expirationTime: { hour: number, minute: number, meriden: 'PM' | 'AM', format: 24 | 12 };
  restaurants$: Observable<Restaurant[]> = of([]);

  lobbyForm = this.fb.group({
    street: [ '', Validators.required ],
    nr: [ '', Validators.required ],
    city: [ '', Validators.required ],
    restaurantId: [ null, Validators.required ],
  });

  static sanitize (field: string) {
    return field.trim().replace(',', ' ');
  }

  ngOnInit() {
    this.expirationTime = {
      hour: this.today.getHours(),
      minute: this.today.getMinutes(),
      meriden: 'PM',
      format: 24
    };

    this.restaurants$ = this.api.getRestaurants();
  }

  createNewLobby() {
    if (this.lobbyForm.valid) {
      const expirationDate = new Date(
          this.today.getFullYear(),
          this.today.getMonth(),
          this.today.getDay(),
          this.expirationTime.hour,
          this.expirationTime.minute
      );

      const control = this.lobbyForm.controls;
      const address = control.nr.value.sanitize + ',' + control.street.value.sanitize + ',' + control.city.value.sanitize;

      const newLobby: PostLobbyDto = {
        restaurant_id: this.lobbyForm.controls['restaurantId'].value,
        owner: 1,
        expires: expirationDate.toISOString(),
        address: address
      };

      this.api.postLobby(newLobby).subscribe(lobby => {
        console.log('Newly created lobby: ');
        console.log(lobby);
      });
    } else {
      console.error('Form is not valid! I won\'t send the request >.<');
    }
  }
}

import { Component, OnInit } from '@angular/core';
import { ApiService } from './_services/api.service';
import { UserService } from './_services/user.service';
import { take } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
  title = 'togeather';

  constructor(
      private api: ApiService,
      private user: UserService,
  ) {
  }

  ngOnInit(): void {
    this.api.checkUserLogin()
        .pipe(take(1))
        .subscribe(
            () => {
              // ignore
            },
            () => {
              this.user.clearLoggedInUser();
            }
        );
  }
}

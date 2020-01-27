import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class UserService {
    private USER_LOCALSTORAGE_KEY = 'loggedInUser';

    get loggedInUser(): string {
        return localStorage.getItem(this.USER_LOCALSTORAGE_KEY);
    }


    set loggedInUser(userName: string) {
        this.clearLoggedInUser();
        localStorage.setItem(this.USER_LOCALSTORAGE_KEY, userName);
    }

    clearLoggedInUser() {
        if (localStorage.getItem(this.USER_LOCALSTORAGE_KEY)) {
            localStorage.removeItem(this.USER_LOCALSTORAGE_KEY);
        }
    }
}

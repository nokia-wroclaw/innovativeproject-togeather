import { Component, OnInit } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { ActivatedRoute } from '@angular/router';
import { catchError, pluck, switchMap } from 'rxjs/operators';
import { ApiService } from '../_services/api.service';
import { throwError } from 'rxjs';
import { ToastrService } from 'ngx-toastr';
import { RedirectionService } from '../_services/redirection.service';
import { Product } from '../_models/product';
import { CartService } from '../_services/cart.service';
import { Cart } from '../_models/cart';

@Component({
    selector: 'app-lobby',
    templateUrl: './lobby.component.html',
    styleUrls: ['./lobby.component.scss']
})
export class LobbyComponent implements OnInit {

    lobby: Lobby;
    cartState: Cart;

    constructor(
        private route: ActivatedRoute,
        private api: ApiService,
        private toaster: ToastrService,
        private redirectionService: RedirectionService,
        private cart: CartService,
    ) { }

    ngOnInit() {
        this.route.params.pipe(
            pluck('lobbyId'),
            switchMap(lobbyId => this.api.getLobby(lobbyId)),
            catchError(error => {
                this.redirectionService.redirectToLobbies();
                this.toaster.error(error, 'Could not load this lobby');
                return throwError(error);
            }),
        ).subscribe(lobby => {
            this.lobby = lobby;
        });
    }

    addProductToCart(item: Product) {
        this.cart.addToCart(item, this.lobby.id).subscribe(
            cart => this.cartState = cart,
            error => this.toaster.error(error, 'Error when adding to cart')
        );
    }

    deleteProductFromCart(item: Partial<Product>) {
        this.cart.deleteFromCart(item, this.lobby.id).subscribe(
            cart => this.cartState = cart,
            error => this.toaster.error(error, 'Error when deleting from cart')
        );
    }
}

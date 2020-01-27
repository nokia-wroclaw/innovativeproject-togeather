import { Component, OnDestroy, OnInit } from '@angular/core';
import { Lobby } from '../_models/lobby';
import { ActivatedRoute } from '@angular/router';
import { catchError, pluck, share, switchMap, takeUntil, tap } from 'rxjs/operators';
import { ApiService } from '../_services/api.service';
import { Subject, throwError } from 'rxjs';
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
export class LobbyComponent implements OnInit, OnDestroy {

    lobby: Lobby;
    cartState: Cart;
    private _unsubscribe$ = new Subject<void>();

    constructor(
        private route: ActivatedRoute,
        private api: ApiService,
        private toaster: ToastrService,
        private redirectionService: RedirectionService,
        private cart: CartService,
    ) { }

    ngOnInit() {
        const lobbyId$ = this.route.params.pipe(
            takeUntil(this._unsubscribe$),
            pluck('lobbyId'),
            // share(),
        );

        lobbyId$.pipe(
            switchMap(id => this.api.getLobby(id)),
            tap(() => console.log('test')),
            catchError(error => {
                this.redirectionService.redirectToLobbies();
                this.toaster.error(error, 'Could not load this lobby');
                return throwError(error);
            }),
        ).subscribe(lobby => {
            this.lobby = lobby;
        });
        //
        // lobbyId$.pipe(
        //     switchMap(id => this.cart.getCart(id)),
        //     tap(() => console.log('test2')),
        //     catchError(error => {
        //         console.log('No cart found');
        //         return throwError(error);
        //     })
        // ).subscribe(cart => {
        //     this.cartState = cart;
        //     console.log(cart);
        // });
    }

    ngOnDestroy(): void {
        this._unsubscribe$.next();
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

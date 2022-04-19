package controllers

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("cannot find product")
	ErrCantDecodeProduct  = errors.New("cannot find product")
	ErrUserIdIsNotValid   = errors.New("user id is invalid")
	ErrCantUpdateUser     = errors.New("cannot add this product to cart")
	ErrCantRemoveItemCart = errors.New("cannot this remove item from cart")
	ErrCantGetItem        = errors.New("cannot this get items from cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveItemFromCart() {

}
func BuyItemFromCart() {

}

func InstantBuy() {

}

// +build amd64,!noasm

package fourq

//go:noescape
func feMul(c, a, b *gfP2)

//go:noescape
func feSquare(c, a *gfP2)

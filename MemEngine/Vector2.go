package MemEngine

type Vector2 struct {
	X, Y float32
}

func NewVector2(X, Y float32) Vector2 {
	return Vector2{
		X: X,
		Y: Y,
	}
}
func NewVector2One() Vector2 {
	return NewVector2(1, 1)
}
func NewVector2Zero() Vector2 {
	return NewVector2(0, 0)
}
func NewVector2UnitX() Vector2 {
	return NewVector2(1, 0)
}
func NewVector2UnitY() Vector2 {
	return NewVector2(0, 1)
}

func (vec Vector2) Add(Other Vector2) Vector2 {
	return NewVector2(vec.X+Other.X, vec.Y+Other.Y)
}
func (vec Vector2) Sub(Other Vector2) Vector2 {
	return NewVector2(vec.X-Other.X, vec.Y-Other.Y)
}
func (vec Vector2) Mul(Other Vector2) Vector2 {
	return NewVector2(vec.X*Other.X, vec.Y*Other.Y)
}
func (vec Vector2) Div(Other Vector2) Vector2 {
	return NewVector2(vec.X/Other.X, vec.Y/Other.Y)
}
func (vec Vector2) Scale(Other float32) Vector2 {
	return NewVector2(vec.X*Other, vec.Y*Other)
}
func (vec Vector2) ConvertToBytes() []byte {
	var b []byte
	b = append(b, Float32ToBytes(vec.X)...)
	b = append(b, Float32ToBytes(vec.Y)...)
	return b
}

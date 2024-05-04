package MemEngine

type Vector3 struct {
	X, Y, Z float32
}

func NewVector3(X, Y, Z float32) Vector3 {
	return Vector3{
		X: X,
		Y: Y,
		Z: Z,
	}
}
func NewVector3One() Vector3 {
	return NewVector3(1, 1, 1)
}
func NewVector3Zero() Vector3 {
	return NewVector3(0, 0, 0)
}
func NewVector3UnitX() Vector3 {
	return NewVector3(1, 0, 0)
}
func NewVector3UnitY() Vector3 {
	return NewVector3(0, 1, 0)
}
func NewVector3UnitZ() Vector3 {
	return NewVector3(0, 0, 1)
}

func (vec Vector3) Add(Other Vector3) Vector3 {
	return NewVector3(vec.X+Other.X, vec.Y+Other.Y, vec.Z+Other.Z)
}
func (vec Vector3) Sub(Other Vector3) Vector3 {
	return NewVector3(vec.X-Other.X, vec.Y-Other.Y, vec.Z-Other.Z)
}
func (vec Vector3) Mul(Other Vector3) Vector3 {
	return NewVector3(vec.X*Other.X, vec.Y*Other.Y, vec.Z*Other.Z)
}
func (vec Vector3) Div(Other Vector3) Vector3 {
	return NewVector3(vec.X/Other.X, vec.Y/Other.Y, vec.Z/Other.Z)
}
func (vec Vector3) Scale(Other float32) Vector3 {
	return NewVector3(vec.X*Other, vec.Y*Other, vec.Z*Other)
}
func (vec Vector3) ConvertToBytes() []byte {
	var b []byte
	b = append(b, Float32ToBytes(vec.X)...)
	b = append(b, Float32ToBytes(vec.Y)...)
	b = append(b, Float32ToBytes(vec.Z)...)
	return b
}

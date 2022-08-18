package entity

import (
	"github.com/shpeliving/go-dxf/format"
)

type Arc struct {
	*Circle
	Angle []float64 // 50, 51 (Degree)
	XData map[string]string
}

// IsEntity is for Entity interface.
func (a *Arc) IsEntity() bool {
	return true
}

func (a *Arc) AppendXData(key, val string) {
	a.XData[key] = val
}

// NewArc creates a new Arc.
func NewArc(c *Circle) *Arc {
	if c == nil {
		c = NewCircle()
	}
	a := &Arc{
		Circle: c,
		Angle:  []float64{0.0, 180.0},
		XData:  make(map[string]string),
	}
	a.SetEntityType(ARC)
	return a
}

// Format writes data to formatter.
func (a *Arc) Format(f format.Formatter) {
	a.Circle.Format(f)
	f.WriteString(100, "AcDbArc")
	for i := 0; i < 2; i++ {
		f.WriteFloat(50+i, a.Angle[i])
	}
	f.WriteXData(format.RhinoAppID, a.XData)
}

// String write out the String representation
func (a *Arc) String() string {
	f := format.NewASCII()
	a.Format(f)
	return f.Output()
}

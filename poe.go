package jsonport

import "github.com/qbox-io/jgh"

func (j Json)  PoeArray() (ret []Json, err error) {
	ret, err = j.Array()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeBool() (ret bool, err error) {
	ret, err = j.Bool()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeBoolArray() (ret []bool, err error) {
	ret, err = j.BoolArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeFloat() (ret float64, err error) {
	ret, err = j.Float()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeFloatArray() (ret []float64, err error) {
	ret, err = j.FloatArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetBool(keys ...interface{}) (ret bool, err error) {
	ret, err = j.GetBool(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetFloat(keys ...interface{}) (ret float64, err error) {
	ret, err = j.GetFloat(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetInt(keys ...interface{}) (ret int64, err error) {
	ret, err = j.GetInt(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetString(keys ...interface{}) (ret string, err error) {
	ret, err = j.GetString(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeInt() (ret int64, err error) {
	ret, err = j.Int()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeIntArray() (ret []int64, err error) {
	ret, err = j.IntArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeKeys() (ret []string, err error) {
	ret, err = j.Keys()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeLen() (ret int, err error) {
	ret, err = j.Len()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeString() (ret string, err error) {
	ret, err = j.String()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeStringArray() (ret []string, err error) {
	ret, err = j.StringArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeValues() (ret []Json, err error) {
	ret, err = j.Values()
	jgh.PanicOnErr(err)
	return
}

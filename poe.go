package jsonport

import "github.com/qbox-io/jgh"

func (j Json)  PoeArray() (ret []Json) {
	ret, err := j.Array()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeBool() (ret bool) {
	ret, err := j.Bool()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeBoolArray() (ret []bool) {
	ret, err := j.BoolArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeFloat() (ret float64) {
	ret, err := j.Float()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeFloatArray() (ret []float64) {
	ret, err := j.FloatArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetBool(keys ...interface{}) (ret bool) {
	ret, err := j.GetBool(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetFloat(keys ...interface{}) (ret float64) {
	ret, err := j.GetFloat(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetInt(keys ...interface{}) (ret int64) {
	ret, err := j.GetInt(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeGetString(keys ...interface{}) (ret string) {
	ret, err := j.GetString(keys...)
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeInt() (ret int64) {
	ret, err := j.Int()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeIntArray() (ret []int64) {
	ret, err := j.IntArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeKeys() (ret []string) {
	ret, err := j.Keys()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeLen() (ret int) {
	ret, err := j.Len()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeString() (ret string) {
	ret, err := j.String()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeStringArray() (ret []string) {
	ret, err := j.StringArray()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  PoeValues() (ret []Json) {
	ret, err := j.Values()
	jgh.PanicOnErr(err)
	return
}

func (j Json)  Exists(keys ...interface{}) (bool) {
	t := j.Get(keys...).Type()
	return (t != NULL) && (t != INVALID)
}

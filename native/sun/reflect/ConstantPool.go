package reflect

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_cp(getLongAt0, "getLongAt0", "(Ljava/lang/Object;I)J")
	_cp(getUTF8At0, "getUTF8At0", "(Ljava/lang/Object;I)Ljava/lang/String;")
}

func _cp(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/reflect/ConstantPool", name, desc, method)
}

// private native long getLongAt0(Object o, int i);
// (Ljava/lang/Object;I)J
func getLongAt0(frame *rtda.Frame) {
	cp, index := _getPop(frame)
	val := cp.GetConstant(index).(int64)
	frame.PushLong(val)
}

// private native String getUTF8At0(Object o, int i);
// (Ljava/lang/Object;I)Ljava/lang/String;
func getUTF8At0(frame *rtda.Frame) {
	cp, index := _getPop(frame)
	kUtf8 := cp.GetConstant(index).(*heap.ConstantUtf8)
	goStr := kUtf8.Str()
	jStr := heap.JString(goStr)
	frame.PushRef(jStr)
}

func _getPop(frame *rtda.Frame) (cp *heap.ConstantPool, index uint) {
	this := frame.GetThis()
	index = uint(frame.GetIntVar(2))
	cp = this.Extra().(*heap.ConstantPool)
	return
}
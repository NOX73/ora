// Copyright 2014 Rana Ian. All rights reserved.
// Use of this source code is governed by The MIT License
// found in the accompanying LICENSE file.

package ora

/*
#include <oci.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	//	"fmt"
	"unsafe"
)

type oraInt32Define struct {
	environment *Environment
	ocidef      *C.OCIDefine
	ociNumber   C.OCINumber
	isNull      C.sb2
}

func (oraInt32Define *oraInt32Define) define(position int, ocistmt *C.OCIStmt) error {
	r := C.OCIDefineByPos2(
		ocistmt,                                   //OCIStmt     *stmtp,
		&oraInt32Define.ocidef,                    //OCIDefine   **defnpp,
		oraInt32Define.environment.ocierr,         //OCIError    *errhp,
		C.ub4(position),                           //ub4         position,
		unsafe.Pointer(&oraInt32Define.ociNumber), //void        *valuep,
		C.sb8(C.sizeof_OCINumber),                 //sb8         value_sz,
		C.SQLT_VNU,                                //ub2         dty,
		unsafe.Pointer(&oraInt32Define.isNull),    //void        *indp,
		nil,           //ub2         *rlenp,
		nil,           //ub2         *rcodep,
		C.OCI_DEFAULT) //ub4         mode );
	if r == C.OCI_ERROR {
		return oraInt32Define.environment.ociError()
	}
	return nil
}

func (oraInt32Define *oraInt32Define) value() (value interface{}, err error) {
	int32Value := Int32{IsNull: oraInt32Define.isNull < 0}
	if !int32Value.IsNull {
		r := C.OCINumberToInt(
			oraInt32Define.environment.ocierr, //OCIError              *err,
			&oraInt32Define.ociNumber,         //const OCINumber       *number,
			C.uword(4),                        //uword                 rsl_length,
			C.OCI_NUMBER_SIGNED,               //uword                 rsl_flag,
			unsafe.Pointer(&int32Value.Value)) //void                  *rsl );
		if r == C.OCI_ERROR {
			err = oraInt32Define.environment.ociError()
		}
	}
	value = int32Value
	return value, err
}

func (oraInt32Define *oraInt32Define) alloc() error {
	return nil
}

func (oraInt32Define *oraInt32Define) free() {

}

func (oraInt32Define *oraInt32Define) close() {
	defer func() {
		recover()
	}()
	oraInt32Define.ocidef = nil
	oraInt32Define.isNull = C.sb2(0)
	oraInt32Define.environment.oraInt32DefinePool.Put(oraInt32Define)
}
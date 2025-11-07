//go:build race2

package runtime

import "unsafe"

const race2enabled = true


func race2funcenter(callpc uintptr) {}
func race2funcexit() {}
func race2read(addr unsafe.Pointer) {}
func race2write(addr unsafe.Pointer) {}
func race2readrange(addr unsafe.Pointer, size uintptr) {}
func race2writerange(addr unsafe.Pointer, size uintptr) {}


//go:nosplit
func race2ReadObjectPC(t *_type, addr unsafe.Pointer, callerpc, pc uintptr)  {}
//go:nosplit
func race2WriteObjectPC(t *_type, addr unsafe.Pointer, callerpc, pc uintptr) {}
//go:nosplit
func race2init() (uintptr, uintptr) { return 0,0 }
//go:nosplit
func race2fini() {}
//go:nosplit
func race2proccreate() uintptr { return 0 }
//go:nosplit
func race2procdestroy(ctx uintptr) {}
//go:nosplit
func race2mapshadow(addr unsafe.Pointer, size uintptr) {}
//go:nosplit
func race2writepc(addr unsafe.Pointer, callerpc, pc uintptr) {}
//go:nosplit
func race2readpc(addr unsafe.Pointer, callerpc, pc uintptr) {}
//go:nosplit
func race2readrangepc(addr unsafe.Pointer, sz, callerpc, pc uintptr)  {}
//go:nosplit
func race2writerangepc(addr unsafe.Pointer, sz, callerpc, pc uintptr) {}
//go:nosplit
func race2acquire(addr unsafe.Pointer) {}
//go:nosplit
func race2acquireg(gp *g, addr unsafe.Pointer) {}
//go:nosplit
func race2acquirectx(racectx uintptr, addr unsafe.Pointer) {}
//go:nosplit
func race2release(addr unsafe.Pointer) {}
//go:nosplit
func race2releaseg(gp *g, addr unsafe.Pointer) {}
//go:nosplit
func race2releaseacquire(addr unsafe.Pointer) {}
//go:nosplit
func race2releaseacquireg(gp *g, addr unsafe.Pointer) {}
//go:nosplit
func race2releasemerge(addr unsafe.Pointer) {}
//go:nosplit
func race2releasemergeg(gp *g, addr unsafe.Pointer) {}
//go:nosplit
func race2fingo() {}
//go:nosplit
func race2malloc(p unsafe.Pointer, sz uintptr) {}
//go:nosplit
func race2free(p unsafe.Pointer, sz uintptr) {}
//go:nosplit
func race2gostart(pc uintptr) uintptr { return 0 }
//go:nosplit
func race2goend() {}
//go:nosplit
func race2ctxstart(spawnctx, racectx uintptr) uintptr { return 0 }
//go:nosplit
func race2ctxend(racectx uintptr){}
//go:nosplit
func race2notify(c *hchan, idx uint, sg *sudog) {}
//go:nosplit
func race2sync(c *hchan, sg *sudog) {}
//go:nosplit
func race2EnterNewCtx() uintptr { return 0 }
//go:nosplit
func race2RestoreCtx(ctx uintptr) {}

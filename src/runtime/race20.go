//go:build !race2

// Dummy race detection API, used when not built with -race2.

package runtime

import (
	"unsafe"
)

const race2enabled = false

// Because raceenabled is false, none of these functions should be called.

func race2ReadObjectPC(t *_type, addr unsafe.Pointer, callerpc, pc uintptr)  { throw("race") }
func race2WriteObjectPC(t *_type, addr unsafe.Pointer, callerpc, pc uintptr) { throw("race") }
func race2init() (uintptr, uintptr)                                          { throw("race"); return 0, 0 }
func race2fini()                                                             { throw("race") }
func race2proccreate() uintptr                                               { throw("race"); return 0 }
func race2procdestroy(ctx uintptr)                                           { throw("race") }
func race2mapshadow(addr unsafe.Pointer, size uintptr)                       { throw("race") }
func race2writepc(addr unsafe.Pointer, callerpc, pc uintptr)                 { throw("race") }
func race2readpc(addr unsafe.Pointer, callerpc, pc uintptr)                  { throw("race") }
func race2readrangepc(addr unsafe.Pointer, sz, callerpc, pc uintptr)         { throw("race") }
func race2writerangepc(addr unsafe.Pointer, sz, callerpc, pc uintptr)        { throw("race") }
func race2acquire(addr unsafe.Pointer)                                       { throw("race") }
func race2acquireg(gp *g, addr unsafe.Pointer)                               { throw("race") }
func race2acquirectx(racectx uintptr, addr unsafe.Pointer)                   { throw("race") }
func race2release(addr unsafe.Pointer)                                       { throw("race") }
func race2releaseg(gp *g, addr unsafe.Pointer)                               { throw("race") }
func race2releaseacquire(addr unsafe.Pointer)                                { throw("race") }
func race2releaseacquireg(gp *g, addr unsafe.Pointer)                        { throw("race") }
func race2releasemerge(addr unsafe.Pointer)                                  { throw("race") }
func race2releasemergeg(gp *g, addr unsafe.Pointer)                          { throw("race") }
func race2fingo()                                                            { throw("race") }
func race2malloc(p unsafe.Pointer, sz uintptr)                               { throw("race") }
func race2free(p unsafe.Pointer, sz uintptr)                                 { throw("race") }
func race2gostart(pc uintptr) uintptr                                        { throw("race"); return 0 }
func race2goend()                                                            { throw("race") }
func race2ctxstart(spawnctx, racectx uintptr) uintptr                        { throw("race"); return 0 }
func race2ctxend(racectx uintptr)                                            { throw("race") }

func race2notify(c *hchan, idx uint, sg *sudog) { throw("race") }

func race2sync(c *hchan, sg *sudog) { throw("race") }

func race2EnterNewCtx() uintptr { throw("race"); return 0 }
func race2RestoreCtx(ctx uintptr) { throw("race") }

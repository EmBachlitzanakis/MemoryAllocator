package main

import (
	"fmt"
	"unsafe"
)

// MemoryPool represents a simple memory pool for fixed-size allocation
type MemoryPool struct {
	pool     []byte
	poolSize int
	freeList *FreeBlock
}

// FreeBlock is used to manage the free list of memory blocks
type FreeBlock struct {
	next *FreeBlock
}

// NewMemoryPool creates a new memory pool of the given size
func NewMemoryPool(size int) *MemoryPool {
	pool := make([]byte, size)
	return &MemoryPool{
		pool:     pool,
		poolSize: size,
		freeList: (*FreeBlock)(unsafe.Pointer(&pool[0])),
	}
}

// AlignPointer aligns a pointer to the given alignment
func AlignPointer(ptr uintptr, alignment int) uintptr {
	offset := ptr % uintptr(alignment)
	if offset != 0 {
		ptr += uintptr(alignment) - offset
	}
	return ptr
}

// Allocate memory from the pool
func (mp *MemoryPool) Allocate(size int) unsafe.Pointer {
	// Align the size to the size of FreeBlock
	alignedSize := (size + int(unsafe.Sizeof(FreeBlock{})) - 1) & ^(int(unsafe.Sizeof(FreeBlock{})) - 1)

	// Find a suitable free block
	prev := (*FreeBlock)(nil)
	curr := mp.freeList

	for curr != nil {
		currPtr := uintptr(unsafe.Pointer(curr))
		alignedPtr := AlignPointer(currPtr, int(unsafe.Sizeof(FreeBlock{})))
		if alignedPtr+uintptr(alignedSize) <= uintptr(unsafe.Pointer(&mp.pool[0]))+uintptr(mp.poolSize) {
			if prev != nil {
				prev.next = curr.next
			} else {
				mp.freeList = curr.next
			}
			return unsafe.Pointer(alignedPtr)
		}
		prev = curr
		curr = curr.next
	}

	return nil // No suitable block found
}

// Deallocate memory back to the pool
func (mp *MemoryPool) Deallocate(ptr unsafe.Pointer) {
	block := (*FreeBlock)(ptr)
	block.next = mp.freeList
	mp.freeList = block
}

func main() {
	const poolSize = 1024
	allocator := NewMemoryPool(poolSize)

	// Allocate 128 bytes
	p1 := allocator.Allocate(128)
	fmt.Printf("Allocated 128 bytes at %v\n", p1)

	// Allocate 256 bytes
	p2 := allocator.Allocate(256)
	fmt.Printf("Allocated 256 bytes at %v\n", p2)

	// Deallocate memory at p1
	allocator.Deallocate(p1)
	fmt.Printf("Deallocated memory at %v\n", p1)

	// Allocate 64 bytes
	p3 := allocator.Allocate(64)
	fmt.Printf("Allocated 64 bytes at %v\n", p3)
}

# Memory Pool Allocator in Go

This project implements a simple memory pool allocator in Go, which can be used to efficiently allocate and deallocate fixed-size memory blocks. The goal of this project is to minimize memory fragmentation and improve allocation performance by managing memory manually using a memory pool.

## Features

- Custom memory pool with a fixed size
- Efficient allocation and deallocation of memory blocks
- Alignment of memory pointers for proper memory access
- Suitable for fixed-size memory allocation scenarios

## Overview

We define a `MemoryPool` that contains a slice of bytes (`pool`) representing the memory pool and a free list (`freeList`) that tracks free memory blocks. The `Allocate` function allocates memory blocks from the pool, and the `Deallocate` function returns the blocks to the pool.

The memory allocator ensures that memory pointers are aligned according to the size of the `FreeBlock` structure. This is important to prevent misaligned memory access, which can cause crashes or performance degradation on certain architectures.

### Key Components

- **MemoryPool**: Manages the memory pool and the free list of available memory blocks.
- **FreeBlock**: Represents a free block of memory in the pool. It contains a pointer to the next available block.
- **NewMemoryPool**: Initializes a new memory pool of a given size.
- **Allocate**: Allocates memory from the pool and returns an aligned pointer to the allocated memory block.
- **Deallocate**: Returns a previously allocated memory block back to the free list for reuse.
- **AlignPointer**: Aligns a pointer to the specified alignment to ensure proper memory access.

### Memory Alignment

Memory alignment is an important aspect of memory management. In this project, memory blocks are aligned based on the size of the `FreeBlock` structure. The `AlignPointer` function ensures that the allocated pointers are correctly aligned, preventing potential access issues.

## How to Run

   ```bash
   git clone https://github.com/EmBachlitzanakis/MemoryAllocator.git

   cd MemoryPool

   go run ./


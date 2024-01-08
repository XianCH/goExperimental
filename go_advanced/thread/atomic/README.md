## 原子操作

    Load 和 Store 操作：
        atomic.LoadUint32, atomic.LoadUint64, atomic.LoadInt32, atomic.LoadInt64: 用于原子地加载一个整数的值。
        atomic.StoreUint32, atomic.StoreUint64, atomic.StoreInt32, atomic.StoreInt64: 用于原子地存储一个整数的值。

    Add 和 Subtract 操作：
        atomic.AddUint32, atomic.AddUint64, atomic.AddInt32, atomic.AddInt64: 用于原子地将一个整数和另一个整数相加，并返回相加后的结果。
        atomic.SubUint32, atomic.SubUint64, atomic.SubInt32, atomic.SubInt64: 用于原子地将一个整数减去另一个整数，并返回减去后的结果。

    CompareAndSwap 操作：
        atomic.CompareAndSwapUint32, atomic.CompareAndSwapUint64, atomic.CompareAndSwapInt32, atomic.CompareAndSwapInt64: 用于比较并原子地交换一个整数的值，如果当前值等于旧值，则更新为新值，返回是否执行了交换。

    Swap 操作：
        atomic.SwapUint32, atomic.SwapUint64, atomic.SwapInt32, atomic.SwapInt64: 用于原子地交换两个整数的值，并返回原来的值。

    AddWithOverflow 和 SubtractWithOverflow 操作：
        atomic.AddUint32WithOverflow, atomic.AddUint64WithOverflow, atomic.AddInt32WithOverflow, atomic.AddInt64WithOverflow: 用于原子地将一个整数和另一个整数相加，并返回相加后的结果和是否发生溢出。
        atomic.SubUint32WithOverflow, atomic.SubUint64WithOverflow, atomic.SubInt32WithOverflow, atomic.SubInt64WithOverflow: 用于原子地将一个整数减去另一个整数，并返回减去后的结果和是否发生溢出。


Go 语言的 sync/atomic 包提供了一组原子操作，用于在多个 goroutine 之间进行原子操作，以确保操作的原子性和避免竞态条件。以下是 sync/atomic 包中常用的几个原子操作：
    AddInt32、AddInt64：
    用于原子地将一个整数和另一个整数相加，并返回相加后的结果。
```go
func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
```
SwapInt32、SwapInt64：
用于原子地交换两个整数的值，并返回原来的值。

```go
func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)

```
CompareAndSwapInt32、CompareAndSwapInt64：
用于比较并原子地交换一个整数的值，如果当前值等于旧值，则更新为新值，返回是否执行了交换。

```go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
```


LoadInt32、LoadInt64：
用于原子地加载一个整数的值。

```go
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
```
StoreInt32、StoreInt64：
用于原子地存储一个整数的值。

```go
    func StoreInt32(addr *int32, val int32)
    func StoreInt64(addr *int64, val int64)
```
    AddUint32、AddUint64、SwapUint32、SwapUint64、CompareAndSwapUint32、CompareAndSwapUint64、LoadUint32、LoadUint64、StoreUint32、StoreUint64：
    类似于对应的带有 Uint 前缀的整数类型的原子操作。

    AddPointer、SwapPointer、CompareAndSwapPointer、LoadPointer、StorePointer：
    用于对指针类型进行原子操作。

这些原子操作提供了一种在多个 goroutine 之间进行原子操作的手段，以确保对共享资源的访问是线程安全的。在并发编程中，使用原子操作是一种有效的方式来避免竞态条件。

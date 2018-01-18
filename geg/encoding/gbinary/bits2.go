package main

import (
    "fmt"
    "gitee.com/johng/gf/g/encoding/gbinary"
)

func main() {
    // Meta元数据文件数据结构：[键名哈希64(64bit,8byte) 键名长度(8bit,1byte) 键值长度(24bit,3byte) 数据文件偏移量(40bit,5byte)](变长)
    hash   := 521369841259754125
    klen   := 12
    vlen   := 35535
    offset := 80000000

    // 编码
    bits   := make([]gbinary.Bit, 0)
    bits    = gbinary.EncodeBits(bits, hash,   64)
    bits    = gbinary.EncodeBits(bits, klen,    8)
    bits    = gbinary.EncodeBits(bits, vlen,   24)
    bits    = gbinary.EncodeBits(bits, offset, 40)
    buffer := gbinary.EncodeBitsToBytes(bits)
    fmt.Println("meta length:", len(buffer))

    /* 然后将二进制数据存储到元数据文件中，查询数据时涉及到的元数据解码操作如下 */

    // 解码
    metabits := gbinary.DecodeBytesToBits(buffer)
    fmt.Println("hash  :", gbinary.DecodeBits(metabits[0  : 64]))
    fmt.Println("klen  :", gbinary.DecodeBits(metabits[64 : 72]))
    fmt.Println("vlen  :", gbinary.DecodeBits(metabits[72 : 96]))
    fmt.Println("offset:", gbinary.DecodeBits(metabits[96 : 136]))
}
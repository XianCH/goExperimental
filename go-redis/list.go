package goredis

import "fmt"

func ListDemo() {
	err := rdc.LPush(ctx, "mylist", "value1", "value2").Err()
	if err != nil {
		fmt.Printf("redis Lpush error :%v", err)
		return
	}
	err = rdc.RPush(ctx, "mylist", "value3", "value4").Err()
	if err != nil {
		fmt.Println("Error inserting element with RPUSH:", err)
		return
	}

	fmt.Println("Element insert successful")
}

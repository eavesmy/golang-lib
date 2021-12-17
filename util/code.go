/*
# File Name: code.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: 五 12/17 10:20:46 2021
*/

package code

/**
 * 随机字符串
 */
var CHARS = []string{"F", "L", "G", "W", "5", "X", "C", "3", "9", "Z", "M", "6", "7", "Y", "R", "T", "2", "H", "S", "8", "D", "V", "E", "J", "4", "K", "Q", "P", "U", "A", "N", "B"}

// var CHARS = []string{"F", "L", "G", "W", "5", "X", "C", "3", "9", "Z", "M", "6", "7", "Y", "R", "T", "2", "H", "S", "8", "D", "V", "E", "J", "4", "K", "Q", "P", "U", "A", "N", "B"}

const CHARS_LENGTH = 32

/**
 * 邀请码长度
 */
const CODE_LENGTH = 8

/**
 * 随机数据
 */
const SLAT int64 = 1234561

/**
 * PRIME1 与 CHARS 的长度 L互质，可保证 ( id * PRIME1) % L 在 [0,L)上均匀分布
 */
const PRIME1 = 3

/**
 * PRIME2 与 CODE_LENGTH 互质，可保证 ( index * PRIME2) % CODE_LENGTH  在 [0，CODE_LENGTH）上均匀分布
 */
const PRIME2 = 11

/**
 * 生成邀请码, 默认为 8位邀请码
 *
 * @param id 唯一的id. 支持的最大值为: (32^7 - {@link #SLAT})/{@link #PRIME1} = 11452834602
 * @return
 */
func Gen(id int64,length int) string {
	return _gen(id, length)
}

/**
 * 生成邀请码
 *
 * @param id 唯一的id主键. 支持的最大值为: (32^7 - {@link #SLAT})/{@link #PRIME1}
 * @return code
 */
func _gen(id int64, length int) string {

	// 补位
	id = id*PRIME1 + SLAT
	//将 id 转换成32进制的值
	b := make([]int64,length)
	// 32进制数
	b[0] = id
	for i := 0; i < length-1; i++ {
		b[i+1] = b[i] / CHARS_LENGTH
		// 按位扩散
		b[i] = (b[i] + int64(i)*b[0]) % CHARS_LENGTH
	}

	var tmp int64 = 0
	for i := 0; i < length-2; i++ {
		tmp += b[i]
	}
	b[length-1] = tmp * PRIME1 % CHARS_LENGTH

	// 进行混淆
	codeIndexArray := make([]int64,length)
	str := []rune{}
	for i := 0; i < length; i++ {
		codeIndexArray[i] = b[i*PRIME2%length]
		str = append(str, []rune(CHARS[codeIndexArray[i]])...)
	}

	return string(str)

	// StringBuilder buffer = new StringBuilder();
	// Arrays.stream(codeIndexArray).boxed().map(Long::intValue).map(t -> CHARS[t]).forEach(buffer::append);
	// return buffer.toString();
}

/**
 * 将邀请码解密成原来的id
 *
 * @param code 邀请码
 * @return id
 */
func Decode(code string,length int) int64 {
	if len(code) != length {
		return -1
	}
	// 将字符还原成对应数字
	a := make([]int64,length)
	for i := 0; i < length; i++ {
		c := rune(code[i])
		index := findIndex(c)
		if index == -1 {
			// 异常字符串
			return -2
		}
		a[i*PRIME2%length] = int64(index)
	}
	b := make([]int64,length)
	for i := length - 2; i >= 0; i-- {
		b[i] = (a[i] - a[0]*int64(i) + CHARS_LENGTH*int64(i)) % CHARS_LENGTH
	}

	var res int64 = 0

	for i := length - 2; i >= 0; i-- {
		res += b[i]
		if i > 0 {
			res *= CHARS_LENGTH
		}
	}
	return (res - SLAT) / PRIME1
}

/*
 * 查找对应字符的index
 *
 * @param c 字符
 * @return index
 */
func findIndex(c rune) int {
	for i := 0; i < CHARS_LENGTH; i++ {
		if CHARS[i] == string(c) {
			return i
		}
	}
	return -1
}

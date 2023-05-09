package array
/*
旋转图像
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

*/

func rotate(matrix [][]int)  {
    //todo 上下交换
    length := len(matrix)
    for i:=0;i<length/2;i++{
        temp := matrix[i]
        matrix[i] = matrix[length-1-i]
        matrix[length-1-i] = temp
    }
    //todo对角线交换
    for i:=0;i<length;i++{
        for j:=i+1;j<length;j++{
            temp := matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = temp
        }
    }
}
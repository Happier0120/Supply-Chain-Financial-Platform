// 全局变量
const globalObj = {};

// 定义公共变量
globalObj.name = '小明';

// 定义公共方法
globalObj.getAge = function (birthYear) {
    return new Date().getFullYear() - birthYear;
};

export default globalObj
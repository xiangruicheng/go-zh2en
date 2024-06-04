# go-zh2en

## 项目背景
由于项目是做海外业务的。为了应对当地的监管，需要将项目中的所有中文去掉或者转换成英文

## 思路
* 遍历文件目录。
* 逐行读取文件内容。
* 识别出汉字
* 调用翻译API：
  * 我这里用的是youdao翻译，用的不是正规途径API，是模拟调用youdao网页的api实现的，优势在于免费且没有频率限制，相当的舒爽
* 替换文件中的原始内容。
* 保存修改后的文件。


## 用法
* 检查某个目录下包含中文的行信息
```
go run main.go check 目录绝对路径
```

* 翻译并替换某个目录下的中文
```
go run main.go replace 目录绝对路径
```

## 示例
* test/下有一个route.php;原文如下
```
<?php
aaaaa;//信息列表
echo "人们的关怀{$uid}";
```
* 检查
```
go run main.go check test/  
File test/route.php contains Chinese characters on line 2: aaaaa;//信息列表
File test/route.php contains Chinese characters on line 3: echo "人们的关怀{$uid}";
total rows:2
```
* 替换后
```
<?php
aaaaa;//information list
echo "People care{$uid}";
```

## 腾讯云函数 - 字符串相似度比较

### 依赖

- 云函数平台使用腾讯云 SCF
- 余弦相似度算法
- 分词 github.com/go-ego/gse

### 使用说明

- 测试接口：https://service-ij14p6y2-1254302252.gz.apigw.tencentcs.com/release/similarity

- **请求接口**

  ```
   POST SCF_BASE_URL HTTP/1.1
   Content-Type: application/json; charset=utf-8

  ```

  `SCF_BASE_URL` 为腾讯云函数相应 API 网关的访问路径，如 `https://service-dlxxjcx0-1xxx02252.gz.apigw.tencentcs.com/release/xxxxxxx`

* **请求参数**

  | 参数名 | 类型     | 必填 | 值  | 说明                             |
  | ------ | -------- | ---- | --- | -------------------------------- |
  | data   | []object | 是   |     | 传入对象数组可一次校验多对字符串 |

  `object`
  | 参数名 | 类型 | 必填 | 值 | 说明 |
  | ------ | ------ | ---- | --- | ---- |
  | id | int | 是 | | 该对字符串的 id |
  | first | string | 是 | | 第一个字符串 |
  | second | string | 是 | | 第二个字符串 |

- **响应参数**

  | 参数名 | 类型     | 必填 | 值            | 说明                                   |
  | ------ | -------- | ---- | ------------- | -------------------------------------- |
  | code   | int      | 是   | `2000`,`2901` | 状态码，`2000`成功，`2901`失败         |
  | msg    | string   | 是   |               | 状态码提示                             |
  | data   | []object | 是   |               | 传入对象数组可一次校验一对或多对字符串 |

  `object`
  | 参数名 | 类型 | 必填 | 值 | 说明 |
  | ------ | ------ | ---- | --- | ---- |
  | id | int | 是 | | 相应请求该对字符串的 id |
  | probability | float | 是 | `0.00` ~ `1.00`| 该对字符串的相似度，精确到小数点后两位 |

### 部署到云函数

- [编译打包](https://cloud.tencent.com/document/product/583/18032#.E7.BC.96.E8.AF.91.E6.89.93.E5.8C.85)

- 将打包后的`main`和`dictionary.txt`文件压缩到`xxx.zip`

- 云函数部署，参考[文档](https://cloud.tencent.com/document/product/583/19806)

- 触发器部署，参考[文档](https://cloud.tencent.com/document/product/583/30230)，部署为`API网关触发器`

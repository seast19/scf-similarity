## 腾讯云函数 - 字符串相似度比较

### 依赖

- 云函数平台使用腾讯云 SCF

### 使用说明

- **请求接口**

  ```
   POST SCF_BASE_URL HTTP/1.1
   Content-Type: application/json; charset=utf-8

  ```

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

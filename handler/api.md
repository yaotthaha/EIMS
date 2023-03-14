## API

#

```
POST /api/employee // 添加Employee

Request==>
{
    "token": "",
    "name": "",
    "sex": "" // 性别: 男/女
    "phone": "",
    "email": "",
    "position": "",
    "marry": "" // 婚姻情况: 已婚/未婚/离婚
    "education": "",
    "join_time": 0, // 时间戳
    "department_id": 0
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": 1000000,
        "name": "",
        "sex": "" // 性别: 男/女
        "phone": "",
        "email": "",
        "position": "",
        "marry": "" // 婚姻情况: 已婚/未婚/离婚
        "education": "",
        "join_time": 0, // 时间戳
        "department_id": 0
    }
}

{"code": 100, "message": "field duplication: ${field}"}
{"code": 101, "message": "field invalid: ${field}"}
```

```
PATCH /api/employee // 修改Employee

Request==>
{
    "token": "",
    "user_id": ""
    "new_name": "",
    "new_sex": "" // 性别: 男/女
    "new_phone": "",
    "new_email": "",
    "new_position": "",
    "new_marry": "" // 婚姻情况: 已婚/未婚/离婚
    "new_education": "",
    "new_join_time": 0, // 时间戳
    "new_department_id": 0
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "user_id": 1000000,
        "name": "",
        "sex": "" // 性别: 男/女
        "phone": "",
        "email": "",
        "position": "",
        "marry": "" // 婚姻情况: 已婚/未婚/离婚
        "education": "",
        "join_time": 0, // 时间戳
        "department_id": 0
    }
}

{"code": 101, "message": "field invalid: ${field}"}
```

```
GET /api/employee // 获取Employee

Request==>
{
    "token": "",
    "user_id": ""
    // ...
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "count": 1,
        "list": [
            {
                "user_id": ""
                ...
            }
        ]
    }
}
```

```
DELETE /api/employee // 删除Employee

Request==>
{
    "token": "",
    "user_id": ""
}

Response<==
{
    "code": 0,
    "message": "success"
}
```

```
GET /api/employee_list // 获取全部Employee

Request==>
{
    "token": ""
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "count": 1,
        "list": [
            {
                "user_id": ""
                ...
            }
        ]
    }
}
```

```
GET /api/department // 获取Department

Request==>
{
    "token": ""，
    "department_id": ""
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "department_id": "",
        "name": ""
    }
}
```

```
GET /api/department_list // 获取全部Department

Request==>
{
    "token": ""
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "count": 1,
        "list": [
            {
                "department_id": "",
                "name": ""
            }
        ]
    }
}
```

```
POST /api/department // 添加Department

Request==>
{
    "token": "",
    "name": ""
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "department_id": "",
        "name": ""
    }
}
```

```
DELETE /api/department // 删除Department

Request==>
{
    "token": "",
    "department_id": ""
}

Response<==
{
    "code": 0,
    "message": "success"
}
```

```
PATCH /api/department // 修改Department

Request==>
{
    "token": "",
    "department_id": "",
    "new_name": ""
}

Response<==
{
    "code": 0,
    "message": "success",
    "data": {
        "department_id": "",
        "name": ""
    }
}
```


# bsn_backend API Reference

🔗 BSN区块链比赛,对参赛合约进行包装,提供HTTP API,以方便测试调用,本文主要描述提供的基本接口信息.

## ☛ <font color=#4CAF50>POST</font> <font color=#8E24AA>RegisterProduct</font> <font color=#26C6DA>注册产品</font>

**Description**

注册产品

>  产品生产厂家生产出实体产品后调用该接口以将产品信息上链

**Path**

```
/products/register
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "type": "TV",
    "product_id": "0x0001",
    "batch_id": "0x1001",
    "material_batches": [
        "0x2001",
        "0x2002"
    ]
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "is_succeed": true
    }
}
```

**cURL Example**

```
curl --location --request POST 'http://127.0.0.1:8000/products/register' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "type": "TV",
    "product_id": "0x0001",
    "batch_id": "0x1001",
    "material_batches": [
        "0x2001",
        "0x2002"
    ]
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>ProductDetails</font> <font color=#26C6DA>获取产品详情</font>

**Description**

获取产品详情

> 通过产品ID获取产品详情,批次号,组成产品的零部件批次号等

**Path**

```
/products/info
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "product_id": "0x0001"
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "owner": "0x391D8f94787275F62727442a59e815fD2358b309",
        "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
        "created_at": 1590041799456,
        "batch_id": "0x1001",
        "material_batches": [
            "0x2001",
            "0x2002"
        ],
        "is_sold": false
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/products/info' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "product_id": "0x0001"
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>GetProducts</font> <font color=#26C6DA>获取产品数量</font>

**Description**

获取产品数量

> 通过产品型号获取产品数量

**Path**

```
/products
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "product_type": "TV"
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "products": {
            "0x0001": {
                "owner": "0x391D8f94787275F62727442a59e815fD2358b309",
                "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
                "created_at": 1590041799456,
                "batch_id": "0x1001",
                "material_batches": [
                    "0x2001",
                    "0x2002"
                ],
                "is_sold": false
            },
            "0x0002": {
                "owner": "0x391D8f94787275F62727442a59e815fD2358b309",
                "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
                "created_at": 1590129698531,
                "batch_id": "0x1001",
                "material_batches": [
                    "0x2001",
                    "0x2002"
                ],
                "is_sold": false
            }
        }
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/products' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "product_type": "TV"
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>RegisterMaterial</font> 注册零部件

**Description**

注册零部件

> 零部件生产厂商生产实体零件后调用该接口注册零部件

**Path**

```
/material/register
```

**Request Body **(json)

```json
{
    "auth_name": "LCD_PRODUCER",
    "material_type": "LCD",
    "batch_id": "0x2001",
    "total_num": 100
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "material_type": "LCD",
        "batch_id": "0x2001",
        "producer": "0xd862a00C790423811e901a1fAe8920c5ec0aD5e1",
        "created_at": "2020-05-22 14:48:36"
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/material/register' \
--data-raw '{
    "auth_name": "LCD_PRODUCER",
    "material_type": "LCD",
    "batch_id": "0x2001",
    "total_num": 100
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>GetMaterial</font> 获取零部件数量

**Description**

获取零部件数量

> 通过零部件生产厂商生产实体零件后调用该接口注册零部件

**Path**

```
/material
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "material_type": "LCD"
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "num": 0
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/material' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "material_type": "LCD"
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>ConsumeMaterial</font> 消耗零部件

**Description**

消耗零部件

> 生产产品过程消耗零部件则调用该接口消耗掉注册的零部件

**Path**

```
/material/consume
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "material_type": "LCD",
    "count": 100
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "material_type": "LCD",
        "batch_id": "0x2001",
        "producer": "0xd862a00C790423811e901a1fAe8920c5ec0aD5e1",
        "created_at": "2020-05-22 14:48:36"
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/material/consume' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "material_type": "LCD",
    "count": 100
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>MakeOrder</font> 创建订单

**Description**

创建订单

> 创建订单接口,可由Custom下产品制造商下产品订单或由产品制造商向零部件制造商下零部件订单

**Path**

```
/order
```

**Request Body **(json)

```json
{
    "payer": "Custom",
    "producer": "TV_PRODUCER",
    "is_material": false,
    "type": "TV",
    "price": 999,
    "count": 10
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "order_id": "1"
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/order' \
--data-raw '{
    "payer": "Custom",
    "producer": "TV_PRODUCER",
    "is_material": false,
    "type": "TV",
    "price": 999,
    "count": 10
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>GetOrder</font> 获取订单详情

**Description**

获取订单详情

> 通过订单ID查询订单详情

**Path**

```
/order/info
```

**Request Body **(json)

```json
{
    "order_id": 1
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "payer": "0x5248840997066A50D6630869A42d2909591a6fEA",
        "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
        "amount": 30000,
        "count": 10,
        "type": "TV"
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/order/info' \
--data-raw '{
    "order_id": 1
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>GetOrders</font> 批量获取订单详情

**Description**

批量获取订单详情

> 通过订单发起方查询订单详情

**Path**

```
/order/all
```

**Request Body **(json)

```json
{
    "from": "Custom"
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "1": {
            "payer": "0x5248840997066A50D6630869A42d2909591a6fEA",
            "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
            "amount": 30000,
            "count": 10,
            "type": "TV"
        }
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/order/all' \
--data-raw '{
    "from": "Custom"
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>ConfirmOrders</font> 确认订单

**Description**

确认订单

> 线下交付订单产品后调用该接口确认订单

**Path**

```
/order/confirm
```

**Request Body **(json)

```json
{
    "auth_name": "Custom",
    "order_id": 1
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "is_succeed": true
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/order/confirm' \
--data-raw '{
    "auth_name": "Custom",
    "order_id": 1
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>CancelOrders</font> 取消订单

**Description**

取消订单

> 取消订单

**Path**

```
/order/cancel
```

**Request Body **(json)

```json
{
    "auth_name": "Custom",
    "order_id": 1
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "is_succeed": true
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/order/cancel' \
--data-raw '{
    "auth_name": "Custom",
    "order_id": 1
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>Mint</font> 铸币

**Description**

铸币

> 给测试角色充值代币,用于订单交易

**Path**

```
/funds/mint
```

**Request Body **(json)

```json
{
    "auth_name": "RootAdmin",
    "account": "Custom",
    "amount": 10000
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "is_succeed": true
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/funds/mint' \
--data-raw '{
    "auth_name": "RootAdmin",
    "account": "Custom",
    "amount": 10000
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>BalanceOf</font> 查询账户余额

**Description**

查询账户余额

> 查询账户余额

**Path**

```
/funds/balance
```

**Request Body **(json)

```json
{
    "account": "Custom"
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "account": "Custom",
        "balance": 89000
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/funds/balance' \
--data-raw '{
    "account": "Custom"
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>Trace</font> 物料溯源

**Description**

物料溯源

> 根据物料溯源

**Path**

```
/funds/balance
```

**Request Body **(json)

```json
{
    "material_batch": "0x2001"
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "product_ids": [
            "0x0001",
            "0x0002",
            "0x0003",
            "0x0004",
            "0x0005"
        ]
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/trace' \
--data-raw '{
    "material_batch": "0x2001"
}'
```

## ☛ <font color=#008000>POST</font> <font color=#8E24AA>SetPrice</font> 设置价格

**Description**

设置价格

> 设置产品价格

**Path**

```
/price
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "type": "TV",
    "price": 100,
    "is_material": false
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "price": [
            {
                "name": "TV",
                "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
                "price": 100,
                "is_material": false
            }
        ]
    }
}
```

**Example**

```
curl --location --request POST 'http://127.0.0.1:8000/price' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "type": "TV",
    "price": 100,
    "is_material": false
}'
```

## ☛ <font color=#008000>GET</font> <font color=#8E24AA>GetPrice</font> 获取产品价格

**Description**

获取产品价格

> 获取产品价格

**Path**

```
/price
```

**Request Body **(json)

```json
{
    "auth_name": "TV_PRODUCER",
    "type": "TV",
    "is_material": false
}
```

**Response**

```json
{
    "code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "price": [
            {
                "name": "TV",
                "producer": "0x391D8f94787275F62727442a59e815fD2358b309",
                "price": 100,
                "is_material": false
            }
        ]
    }
}
```

**Example**

```
curl --location --request GET 'http://127.0.0.1:8000/price' \
--data-raw '{
    "auth_name": "TV_PRODUCER",
    "type": "TV",
    "is_material": false
}'
```


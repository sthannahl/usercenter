# 开发日志
## oauth客户端
### 鉴权
```
curl --location --request POST 'http://localhost:8888/token' \
--header 'Authorization: Basic eHRiOjEyMzQ1Ng==' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--header 'Cookie: go_session_id=ZDU3ZDllZTAtNmIwMC00NmQ0LWJiZjgtZGI5NmIwMGNjNWQ5.361756f3aad2f34bd3e21c85916d81ef51700e90' \
--data-urlencode 'grant_type=client_credentials'
```
基于MongoDB
- 表: `oauth2_clients`
1. 客户端秘钥保存到数据库（原为配置文件，考虑不停机新增客户端做此变化）
2. 新增客户端
    a. 秘钥bcrypt加密与检验
3. 客户端鉴权并返回token

## 用户
### 用户注册
```
curl --location --request POST 'http://localhost:8888/signUp' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ4dGIiLCJleHAiOjE1OTgyNjEzODl9.WCGsLYbUXqHCKoe1CZnUMMDU38p6Meui1I5UIwGDu1B_Uc5P4er0uigCqgNC1zz5zqNgnw4UDzB5f9ACRV1m_w' \
--header 'Content-Type: application/json' \
--header 'Cookie: go_session_id=ZDU3ZDllZTAtNmIwMC00NmQ0LWJiZjgtZGI5NmIwMGNjNWQ5.361756f3aad2f34bd3e21c85916d81ef51700e90' \
--data-raw '{"user_id":"sth"}'
```
### 用户查询
```
curl --location --request GET 'http://localhost:8888/user?user_id=sth' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ4dGIiLCJleHAiOjE1OTgyNjIxNzN9._WLeBvsffLawWr3vkteYSoR3PaE0Xxgy6SV-TYiH9NrrejTUflptRNMfsDjRdwzfvtM4sWRGRRi4cAsKsdZBtw' \
--header 'Cookie: go_session_id=ZDU3ZDllZTAtNmIwMC00NmQ0LWJiZjgtZGI5NmIwMGNjNWQ5.361756f3aad2f34bd3e21c85916d81ef51700e90'
```
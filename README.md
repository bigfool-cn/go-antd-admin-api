# Go-Antd-Admin-Api
[react-antd-admin](https://github.com/bigfool-cn/react-antd-admin) 服务端接口项目(Gin + Gorm + Jwt)

[预览](https://react.bigfool.cn)

## 使用
### 环境
 - go：1.14.10
 - mysql：5.7.28

### 数据库
根目录下的react-antd-admin.sql文件导入mysql中

### 配置
configs/temp.configs.yaml 修改为 configs/configs.yaml

替换configs/configs.yaml配置值为自己环境的配置

### 启动
项目使用go module管理依赖
 
GO111MODULE 请设置为 on

建议使用bee启动项目：bee run


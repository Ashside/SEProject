# 2024年春季实践项目

## Requirements

### 后端

后端使用 Go 语言编写，版本为 1.20，适用于 Windows / amd64 架构。Web 框架采用 gin-gonic / gin v1.7.4。你需要安装 Go 和相关的库来保证程序的正常运行。

### 前端

前端使用 Vue3 进行编写，你需要安装 Node.js 和 Vue 来保证程序的正常运行。

### 数据库

数据库使用 MySQL 数据库。

## 实现功能

## 如何部署和配置

### 前端部署

#### `npm` 部署

前端你可以直接使用 `npm` 进行部署。要确保你已经安装了 Node.js，如果没有安装，请前往 https://nodejs.org/ 进行安装。安装完成后，使用：

```
npm install
```

`npm` 会自动完成依赖配置，之后使用命令：

```
npm run serve
```

运行前端服务，并在浏览器中查看。

#### Docker 部署

使用 Docker 运行前端界面，切换到 `\frontend` 文件夹下，执行如下指令：

```
docker build -t frontend .
docker run -p 8080:3000 frontend
```

之后在本机的 `8080` 端口即可完成访问。

#### 运行 IP 配置

对于后端不同的服务器 IP 地址，可以在 `src/frontend/.env` 中进行配置。后端 API 访问的 IP 地址需修改一下字段：

```
VUE_APP_API_URL = http://localhost:8081
```

重新启动并运行。

## OpenAPI

我们使用 Apifox 组织 API，你可以在[此网站](https://apifox.com/apidoc/shared-6bd451e3-8d10-40a4-bb52-5ce49f6262de)中查到目前开放的API接口。

## Acknowledgement

1. 感谢 OpenAI [ChatGPT](https://chatgpt.com/#) 的大力支持。

## 策划案

### 项目名称

题库管理与在线考试

### 基本需求
- 题库管理：能够录入题目、答案、难易度值等
- 能够处理多种题型：选择题（单选、多选）、判断题、问答题
- 组卷：选择题目构成试卷、重复判断、题目构成分析等
- 权限管理：老师之能维护各自题目，管理员有权组卷等
- 附加功能：在线答题、自动评判（选择题、判断题）

### 前期规划

- 前端
  - 录入题目（编辑器）
    - markdown文本
    - 图片文件（包括markdown中引用的）
    - 设定难易度（数值）
    - 设定题目类型（选择题、判断题、问答题）
    - 设定答案（文本及图片，与题目类型对应）
    - 上传题目
    
  - 组卷
    - 从后端获取题目（按难易度、类型等筛选）
    - 展示筛选后的题目
    - 添加可变数量的题目（按链表处理？）
    - 生成试卷（pdf格式）
    - 题目查重
    - 试卷分析（题目类型、难易度分布，或者其他tag，需要定义）
    - 上传试卷

  - 判卷
    - 自动判卷（选择题、判断题）
  - 在线答题
    - 从后端获取试卷
    - 自动判卷（选择题、判断题）
- 后端
  - 题库管理
    - 题目增删改查
      - 统一管理上传的图片文件和markdown文本
      - 需要做到能够定位图片文件
    - 题目类型增删改查
    - 题目难易度增删改查
    - （题目标签增删改查）
  - 组卷
    - 题目筛选
    - 试卷生成
    - 试卷上传
  - 判卷
    - 自动判卷
  - 在线答题
    - 试卷获取
    - 自动判卷

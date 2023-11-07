# 阶段1：构建阶段
FROM node:16 as builder

# 设置工作目录
WORKDIR /usr/src/app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装所有依赖项（包括开发依赖项）
RUN npm install

# 复制源代码到工作目录
COPY . .

# 构建应用程序（如果有构建脚本的话）
RUN npm run build

# 阶段2：生产阶段
FROM node:16-slim

# 设置工作目录
WORKDIR /usr/src/app

# 仅复制阶段1中构建的生产依赖项
COPY --from=builder /usr/src/app/package*.json ./
RUN npm install --only=production

# 从构建阶段复制构建后的代码到当前工作目录
COPY --from=builder /usr/src/app/dist ./dist

# 如果你的应用程序绑定在其他端口，你需要修改这个值
EXPOSE 3000

# 定义运行时的环境变量
ENV NODE_ENV production

# 运行应用程序
CMD ["node", "dist/app.js"]

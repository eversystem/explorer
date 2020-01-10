FROM nginx:lts-alpine

COPY . /work_dir/frontend
WORKDIR /work_dir/frontend
RUN apk add --no-cache nginx nodes
RUN npm install
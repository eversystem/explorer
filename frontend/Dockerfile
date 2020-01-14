FROM nginx

COPY . /work_dir/frontend
WORKDIR /work_dir/frontend
RUN apk add --no-cache nginx 
RUN npm install
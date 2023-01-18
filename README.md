以github Action的方式，通过腾讯云API的方式，触发CDN更新
```
name: refresh-cdn
run-name: ${{ github.actor }} is refreshing cdn
on:
  push:
    branches: [ main ]
jobs:
  refresh-cdn:
    runs-on: ubuntu-latest
    container:
      image: registry.cn-shanghai.aliyuncs.com/codev/tencent-cdn-refresh:latest
    env:
      SECRET_ID: ${{ secrets.SECRET_ID }}
      SECRET_KEY: ${{ secrets.SECRET_KEY }}

    steps:
      - name: sleep 3min
        run: sleep 180
      - name: refresh fluffy.vinf.top
        run: tencent-cdn-refresh purge-path --secret-id=$SECRET_ID --secret-key=$SECRET_KEY https://fluffy.vinf.top
```

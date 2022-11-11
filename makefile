#TAG = test-$(shell git log -1 --format=%h)
TAG = latest
WORK_DIR = .
REGISTRY = registry.cn-shanghai.aliyuncs.com/codev

image:
	docker build -t $(REGISTRY)/tencent-cdn-refresh:$(TAG) -f ./Dockerfile $(WORK_DIR)

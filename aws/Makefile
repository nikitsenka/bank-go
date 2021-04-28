PROFILE=dev
STACK=benchmark

delete:
	AWS_PROFILE=$(PROFILE) aws cloudformation delete-stack --stack-name $(STACK)

deploy:
	AWS_PROFILE=$(PROFILE) aws cloudformation deploy --stack-name $(STACK) --template-file cloudformation.yml

info:
	AWS_PROFILE=$(PROFILE) aws cloudformation describe-stacks --stack-name "$(STACK)"

PROJECT_PATH := ${PWD}
DEVKIT_DIR := $(PROJECT_PATH)/devkit
KIND_DIR := $(DEVKIT_DIR)/kind
KIND := ${shell which kind}
KIND_VERSION := latest
KIND_NODE_TAG := v1.24.3
.PHONY: bits
bits: build manifests docker-build docker-push bundle bundle-build bundle-push

.PHONY: sdk-run
sdk-run:
	@echo ""
	@echo "Using VERSION: $(VERSION)"
	operator-sdk run bundle $(BUNDLE_IMG)

.PHONY: sample
sample:
	@echo ""
	@echo "Creating a DexConfig instance ..."
	oc apply -f ./config/samples/dexconfig.yaml

.PHONY: check
check:
	@echo ""
	@echo "Verify openid configuration"
	@curl -k https://$(shell oc get route -l owner=dex-operator -ojsonpath='{.items[].spec.host}')/.well-known/openid-configuration
	@curl -k -I https://$(shell oc get route -l owner=dex-operator -ojsonpath='{.items[].spec.host}')/.well-known/openid-configuration

.PHONY: cleanup
cleanup:
	operator-sdk cleanup dex-operator

.PHONY: wait
wait:
	sleep 20

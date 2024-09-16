
install_dependencies:
	@chmod 755 scripts/install_dependencies.sh
	@./scripts/install_dependencies.sh

unit_test:
	@chmod 755 scripts/unit_test.sh
	@./scripts/unit_test.sh

integration_test:
	@chmod 755 scripts/integration_test.sh
	@./scripts/integration_test.sh

#generate mocks
gen_auth_mocks:
	@chmod 755 scripts/gen_mock.sh
	@./scripts/gen_mock.sh auth
gen_order_mocks:
	@chmod 755 scripts/gen_mock.sh
	@./scripts/gen_mock.sh order
gen_user_mocks:
	@chmod 755 scripts/gen_mock.sh
	@./scripts/gen_mock.sh user
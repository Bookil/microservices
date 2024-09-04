
install_dependencies:
	@chmod 755 scripts/install_dependencies.sh
	@./scripts/install_dependencies.sh

unit_test:
	@chmod 755 scripts/unit_test.sh
	@./scripts/unit_test.sh

integration_test:
	@chmod 755 scripts/integration_test.sh
	@./scripts/integration_test.sh
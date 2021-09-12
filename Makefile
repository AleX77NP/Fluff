PROJECT_PATH_STAG=/tmp/staging
PROJECT_PATH_PROD=/tmp/production

fluff-test-stag:
	npm install --prefix ${PROJECT_PATH_STAG} && npm test --prefix ${PROJECT_PATH_STAG}

fluff-test-prod:
	npm install --prefix ${PROJECT_PATH_PROD} && npm test --prefix ${PROJECT_PATH_PROD}

fluff-run-stag:
	cd ${PROJECT_PATH_STAG} && docker-compose pull && docker-compose build && HOST_PORT=8081 DB_PORT=27018 docker-compose up -d --remove-orphans --force-recreate

fluff-run-prod:
	cd ${PROJECT_PATH_PROD} && docker-compose pull && docker-compose build && HOST_PORT=8080 DB_PORT=27017 docker-compose up -d --remove-orphans --force-recreate

fluff-cleanup:
	docker image prune -a -f

# test 
fluff-commander-test:
	ls -l

PROJECT_PATH=/tmp/project

fluff-test:
	npm install --prefix ${PROJECT_PATH} && npm test --prefix ${PROJECT_PATH}

fluff-run:
	cd ${PROJECT_PATH} && docker-compose pull && docker-compose build &&  docker-compose up -d --remove-orphans --force-recreate

fluff-cleanup:
	docker image prune -a -f

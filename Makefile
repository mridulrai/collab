run:
	docker-compose run user build

inside:
	docker-compose run user bash

stop:
	docker-compose down

clean:
	docker exec -it localmongo1 mongo my_sales_chmapion --eval "printjson(db.dropDatabase())"
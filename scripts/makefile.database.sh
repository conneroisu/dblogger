

# append all the sql files found in ./data/schemas/*.sql into ./data/combined/schema.sql
gum spin --spinner dot --title "Combining SQL Schemas" --show-output -- \
	cat ./data/schemas/*.sql > ./data/combined/schema.sql

# append all the sql files found in ./data/queries/*.sql into ./data/combined/queries.sql
gum spin --spinner dot --title "Combining SQL Queries" --show-output -- \
	cat ./data/queries/*.sql > ./data/combined/queries.sql

# append all the sql files found in ./data/seeds/*.sql into ./data/combined/seeds.sql
gum spin --spinner dot --title "Combining SQL seeds" --show-output -- \
	cat ./data/seeds/*.sql > ./data/combined/seeds.sql

sqlc generate

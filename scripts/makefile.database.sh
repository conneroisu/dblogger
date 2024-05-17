

# append all the sql files found in ./data/schemas/*.sql into ./data/combined/schema.sql
gum spin --spinner dot --title "Combining SQL Schemas" --show-output -- \
	cat ./data/schemas/*.sql > ./data/combined/schema.sql

# append all the sql files found in ./data/queries/*.sql into ./data/combined/queries.sql
gum spin --spinner dot --title "Combining SQL Queries" --show-output -- \
	cat ./data/queries/*.sql > ./data/combined/queries.sql

# append all the sql files found in ./data/seeds/*.sql into ./data/combined/seeds.sql
gum spin --spinner dot --title "Combining SQL seeds" --show-output -- \
	cat ./data/seeds/*.sql > ./data/combined/seeds.sql

# append all the sql files found in ./data/seeds/*.sql into ./data/combined/schema.sql
gum spin --spinner dot --title "Combining SQL seeds to Schema" --show-output -- \
	cat ./data/seeds/*.sql >> ./data/combined/schema.sql

# Format the sql files
gum spin --spinner dot --title "Formatting SQL Files" --show-output -- \
	sleek ./data/combined/schema.sql ./data/combined/queries.sql ./data/combined/seeds.sql

# generate the sqlc models
gum spin --spinner dot --title "Generating SQLC Models" --show-output -- \
	sqlc generate

# replace lines 28 to the end of the file of ./querier.go with the lines 11 to the end of the file of ./data/querier.go
gum spin --spinner dot --title "Replacing Querier" --show-output -- \
	sed -i '28,/^}/d' ./querier.go && sed -i '1,10d' ./data/querier.go && cat ./data/querier.go >> ./querier.go

gum spin --spinner dot --title "Cleaning" --show-output -- \
	make clean

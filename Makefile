# Makefile
# https://github.com/conneroisu/dblogger/blob/main/Makefile
# -include .env
SHELLFLAGS = -e

tidy:
	@sh ./scripts/makefile.tidy.sh

tailwind-watch:
	@sh ./scripts/makefile.tailwind-watch.sh

database:
	@sh ./scripts/makefile.database.sh

lint:
	@sh ./scripts/makefile.lint.sh

test:
	@sh ./scripts/makefile.test.sh

dev.requirements:
	@sh ./scripts/makefile.dev.requirements.sh

clean:
	@sh ./scripts/makefile.clean.sh

format:
	@sh ./scripts/makefile.format.sh

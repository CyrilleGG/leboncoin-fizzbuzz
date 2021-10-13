setup:
	heroku create lbc-fizzbuzz --region eu --remote lbc-fizzbuzz && git push lbc-fizzbuzz master && heroku addons:create heroku-postgresql:hobby-dev --app lbc-fizzbuzz --as lbcpsql

deploy:
	git push lbc-fizzbuzz master
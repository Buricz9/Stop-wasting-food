﻿# Stop-wasting-food

Remember to write in main.go your apiKey and login, password to database

How to build app:
go build main.go postDataController.go searchRecipes.go getDataController.go

How to run app:
./main --ingredients=apples,sugar,flour --numberOfRecipes=4

if your query exist in the database you'll get the information:
Exist in DB

if your query does not exist in the database you'll get the information:
Does not exist in DB

The way of printing information:
1. Title
2. Missed ingredients
3. Used ingredients
4. Nutrients

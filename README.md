# Stop-wasting-food

Remember to write in main.go your apiKey and login, password to database

How to build app:
go build main.go postDataController.go searchRecipes.go getDataController.go

How to run app:
./main --ingredients=apples,sugar,flour --numberOfRecipes=4

if your query exist in the database:
Exist in DB

if your query does not exist in the database
Does not exist in DB

Way of printing information:
1. Title
2. Missed ingredients
3. Used ingredients
4. Nutrients
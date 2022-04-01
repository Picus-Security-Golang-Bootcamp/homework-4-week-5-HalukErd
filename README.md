# Explanation

* I have implemented workerPool but has violates foreign key constraint errors.
probably we need priority for insert author. They are commented out.
```
// TODO I have implemented workerPool but has violates foreign key constraint errors.
// probably we need priority for insert author
//authorRepo.Migrations()
//bookRepo.Migrations()
//concurrency.InitializeInsertion(bookRepo, authorRepo) //has Errors
```

* I have read BookCsvLine from csv file. It also includes author name.
```go
type BookCsvLine struct {
	Title     string
	Author    string
	Genre     string
	Height    string
	Publisher string
}

type BookCsvLines []BookCsvLine
```



# Homework 3 Week 4

- `Book` ve `Author` bilgileri bir dosyadan okunacak ve DB'ye kayıt edilecek.
- `list`, `search`, `delete(soft-delete)`, `buy` gibi `os.Args` komutları yerine DB sorguları yazılacak. 
- Bu 2 modelle alakalı GORM sorguları yazılacak.
  - GetByID
  - FindByName
  - GetBooksWithAuthor
  - GetAuthorWithBooks etc. (GORM dökümantasyondaki sorgu çeşitlerine bakılacak bu 2 modelde uygulanacak)
  - (Sadece 4 sorgu değil olabildiğince sorgu yazıp kendinizi geliştirin. Bu size artı olarak dönüş olacaktır.)

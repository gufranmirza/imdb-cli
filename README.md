# imdb-cli
A command line utility to scrap top movies from IMDB website.


## Running IMDB CLI
Clone reposiroty
```
git clone https://github.com/gufranmirza/imdb-cli
```

Go to `src` folder and run 
```
./imdb-cli top-movies --url=https://www.imdb.com/india/top-rated-indian-movies/ --limit=5
```
Where url is the IMDB page of the top listed movie 

CLI output 
```
[
        {
                "title": "Pather Panchali (1955)",
                "movie_release_year": "1955",
                "imdb_rating": "8.6",
                "summary": "Impoverished priest Harihar Ray, dreaming of a better life for himself and his family, leaves his rural Bengal village in search of work.",
                "duration": "2h 5min",
                "genre": "Drama"
        },
        {
                "title": "Ratsasan (2018)",
                "movie_release_year": "2018",
                "imdb_rating": "8.7",
                "summary": "A Sub-Inspector sets out in pursuit of a mysterious serial killer who targets teen school girls and murders them brutally",
                "duration": "2h 50min",
                "genre": "Action"
        },
        {
                "title": "Gol Maal (1979)",
                "movie_release_year": "1979",
                "imdb_rating": "8.6",
                "summary": "A man's simple lie to secure his job escalates into more complex lies when his orthodox boss gets suspicious.",
                "duration": "2h 24min",
                "genre": "Comedy"
        },
        {
                "title": "Nayakan (1987)",
                "movie_release_year": "1987",
                "imdb_rating": "8.6",
                "summary": "A common man's struggles against a corrupt police force put him on the wrong side of the law. He becomes a don, who is loved and respected by many, but his growing power and influence exact a heavy toll.",
                "duration": "2h 36min",
                "genre": "Crime"
        },
        {
                "title": "Anbe Sivam (2003)",
                "movie_release_year": "2003",
                "imdb_rating": "8.7",
                "summary": "Two men, one young and arrogant, the other damaged - physically but not spiritually - by life, are thrown together by circumstances, and find that they are in some ways bound together by fate.",
                "duration": "2h 40min",
                "genre": "Adventure"
        }
]
```

## Development
Run the below command to rebuild the project
```
make all
```

It will rebuild the project and create the `imdb-cli` binary with latest changes

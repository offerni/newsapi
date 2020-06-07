# News Api for Go

Unnoficial [News Api](https://newsapi.org/) SDK,
please check their [V2 Documentation](https://newsapi.org/docs) 

To import the package on your project: `github.com/offerni/newsapi`

## Client
To use it, first instantiate the client with your APIKey like:
```
newsapiClient, err := newsapi.NewClient(newsapi.ClientOpts{ApiKey: "yourApiKey"}) 
``` 

## Endpoint Methods
This sdk supports all the endpoints and request parameters available at [News API Endpoints](https://newsapi.org/docs/endpoints).<br>

After having an instance of `newsapiClient`, you can call the endpoints passing the request parameters in the `...Opts` structs as below:

### GetSources
```
sources, err := newsapiClient.GetSources(newsapi.SourcesOpts{Language: "pt", Country: "br"})
```
<sub>**Note:** Passing an empty `SourcesOpts` will return all sources as the request parameters are not required.</sub>
### GetTopHeadlines
```
topHeadlines, err := newsapiClient.GetTopHeadlines(newsapi.TopHeadlinesOpts{Q: "test", Country: "ca"})
```
### GetEverything
```
everything, err := newsapiClient.GetEverything(newsapi.EverythingOpts{Q: "programming", Language: "en"})
```

##### Beware this is not yet released and major changes can happen.

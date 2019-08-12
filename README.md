I took about 20 minutes to come up with an architecture proposal of this implementation.
This architecture hinted me that it would take nearly a day to complete a working application even if I use Java.
So, I thought in 2 hours the best would be I could come up with a skeleton of services & functions.
This thought process ignited me to try golang at least with pseudo code.
Architecture diagram of this implementation named TopRepos.pdf and available in this git repo.

Personally I don't like comments in the code unless I feel it's essential to successor developers.
The function names and flow of calling the functions itself should talk i.e
building small functions and compose them for big functionality.
This practice will eventually help us in multiple advantages.
1. Identify the concern of each function i.e functions with clear boundaries.
2. The above point helps us to write good variety of unit test cases (in fact it favors TDD)
3. Eventually code becomes declarative rather than imperative.

I tried to implement the functionality (with pseudocode) in the same passion.
All most all the comments are TODO comments & pseudo code.
All this comments should be removed after actual implementation and code should be self explanatory.

There are 2 main directories, each one in it's own docker image.
They are stateless and hence can be scaled horizontally on demand.
1. getRepos -> service to fetch the repositories created in last 7 days.
	This service saves the repositories to mySql DB.
	TODO: This service can be configured as a cron job or expose a REST API to fetch repos from gitHub.
2. topRepos -> service to fetch the repositories created in last 7 days.
	This service exposes graphQl endpoint to solve over fetching problem.
	This service fetches the top 10 repository details from DB and cherry pick attributes to return.

Dockertizing the DB:
It is good only for QA environment. Dockers help to ease the deployment.
Scaling is a side effect. DB holds the state and it eventually become huge.
Having such a huge foot print, I don't see real advantage for DBs in containers.
However, for QA, it makes sense to spin up the DB with needed seeds just like that.

Event driven approach: // TODO
I would choose rabbitMq as the message size seems to be considerably big.
For each second, crazy number of new repositories are getting created.
There can be a overlap in the message content among messages.
That means it's not strictly chronological,hence I think rabbitMq suits my need.

In this pseudo implementation, I could not try event driven model due to 2 hours time limit.
However, in my opinion the ask of event driven model is because to explain my understanding.
For this requirement it can be over-engineering, however, when the through-put is more and ad hoc
and data to be processed is more, then event driven approach can be considered.

my observation is repositories on gitHub are getting created for every second like crazy.
Unless there is a need to return repositories created in a date range of past(this range should be small)
I don't see a need to cache.

As the gitHub API returns tons of attributes in the response, I see the over-fetching issue.
Hence I could understand the need of API query language like GraphQl.

I've added the pseudo test cases as well.

Well,
In my opinion, I've created a decent skeleton having the functions and files with clear separation of concerns.
It has clearly defined integration points helps any changes/enhancements easily.
I believe, using this scaffolding even a junior golang developer can quickly implement the functionality.


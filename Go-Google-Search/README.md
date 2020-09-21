<h2>This assignment is a implementation of "Google Search" problem presented by Rob Pike in Google developers Go conference.</h2>

<strong>Referances :</strong>
1. <a href="https://talks.golang.org/2012/concurrency.slide#42">Concurrency Slides</a>
2. <a href="https://www.youtube.com/watch?v=f6kdp27TYZs">Go Concurrency Patterns - Talk by Rob Pike</a>

<strong>Google search v1</strong>
1. Here we simply query the server (or call fakeSearch) for web, image and video type of results one after another in sequntial manner.
2. No concurrency in used
3. Total Time Taken = T(web results) + T(image results) + T(Video results)

<strong>Google Search v2</strong>
1. Here instead of sequential mannner, we call separate concurrent goroutines for web,image and video results
2. Total Time Taken = min(T(web results) , T(image results) , T(Video results))

<strong>Google Search v2.1</strong>
1. Same as v2. Addition to that, we consider some threshold time.
2. If any goroutine takes more time to fetch result than threshold time, we discard that result.
3. Total Time Taken = min( min(T(web results) , T(image results) , T(Video results)) , threshold_time)

<strong>Google Search v3</strong>
1. In addition to v2, we add support for replication
2. We send replicated queries for fetching result on separate concurrent goroutines, consider result of one whichever returns data first.
3. Considering n number of replicas, <br>
 T(web results) = min( T(web replica 1) , T(web replica 2) , ....... , T(web replica n)) <br>
 T(image results) = min( T(image replica 1) , T(image replica 2) , ....... , T(image replica n)) <br>
 T(video results) = min( T(video replica 1) , T(video replica 2) , ....... , T(video replica n)) <br>
 Total Time Taken = min( min(T(web results) , T(image results) , T(Video results)) , threshold_time) <br>

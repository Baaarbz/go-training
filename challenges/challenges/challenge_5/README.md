<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-3.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center" style="text-align: center;">
  Challenge #5. Go routines
</h1>

In this 5th challenge we are going to use go routines to see if the ad was posted also on other
platforms

## Instructions


* Create a new method that receives the ad, reads the list of sites that should be searched from the database and saves
  in a new table the connections between the ad id and id's of sites that was also posted
* The sites table can be filled with how many sites you want (manual or script, however you want).
* You can use random functions to return if a ad was posted to a site or no
* The method should be called async while saving the ad and not interrupt the save flow.


## Resources
1. https://go.dev/tour/concurrency/1


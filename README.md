# Haversine Formula

 Haversine formula calculates the distance between two points by using latitude and logitude on a sphere. There are lot of applications using navigation system and haversine formula are playing big role to calcutate the shortest distance between locations.
 # Formula
 
 Suppose we have two ponits p1 and p2. let radius of p1 is (φ1, λ1) and p2 is (φ2, λ2).
 So difference between p1 and p2, by using haversine formula, is 
```
a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
c = 2 ⋅ atan2( √a, √(1−a) )
d = R ⋅ c
```
where Δφ = φ2 - φ1, Δλ = λ2 - λ1 and R is the Earth radius (6,371km)


License
----

MIT

**Free Software, Hell Yeah!**